package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"study-WODB/internal/config"
	"study-WODB/internal/handler"
	"study-WODB/internal/http/graphQl"
	"study-WODB/internal/logger"
	"study-WODB/internal/storage/mongo"
	"study-WODB/internal/storage/postgres"
	"study-WODB/internal/storage/redis"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

// HttpServer - структура описывающая работу сервера
type HttpServer struct {
	*config.Config
	*logger.Logger
	pStorage *postgres.Storage
	mStorage *mongo.Storage
	rStorage *redis.Storage
}

func New(cnf *config.Config, log *logger.Logger,
	postgres *postgres.Storage, mongo *mongo.Storage, redis *redis.Storage) *HttpServer {
	return &HttpServer{
		cnf,
		log,
		postgres,
		mongo,
		redis,
	}
}

func (s *HttpServer) Start() {
	// создаём и настраиваем роутер
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	}))

	// ===- определение путей и end-point`ов -=== \\
	s.Info("Server initialization...")
	// настройка файлового сервера
	r.Handle("/", http.FileServer(http.Dir(s.FileServer)))

	// настройка ресурсов аутентификации
	s.Debug(" - configuring authentication resources")
	// создание обработчиков
	gAuth := handler.NewGoogleAuth(s.Config, s.Logger)
	yAuth := handler.NewYandexAuth(s.Config, s.Logger)
	r.Route("/auth", func(r chi.Router) {
		r.Get("/local", todo)                       // стандартный аутентификатор
		r.Get("/google", gAuth.GoogleCall)          // Oauth2 аутентификатор от google
		r.Get("/google-callback", gAuth.GoogleBack) // переадресация назад
		r.Get("/ya", yAuth.YandexCall)              // Oauth2 аутентификатор от google
		r.Get("/ya-callback", yAuth.YandexBack)     // переадресация назад
	})

	// настройка graphQL точки входа
	s.Debug(" - configuring GraphQL entry points")
	gQL := graphQl.NewGraphQL(s.Config, s.Logger)
	r.Route("/menu", func(r chi.Router) {
		r.Get("/check", gQL.Check)
	})

	// настройка системы заказов
	s.Debug(" - configuring order entry points")
	order := handler.NewOrderHandler(s.pStorage, s.mStorage, s.Logger)
	r.Route("/order", func(r chi.Router) {
		r.Post("/make", order.Make)
		r.Post("/close", order.Close)
		r.Post("/pay", order.Pay)
		r.Post("/feedback", order.Feedback)
	})

	// настройка методов api
	s.Debug(" - configuring application resources")
	restHandler := handler.NewRestHandler(s.pStorage, s.Logger)
	dishHandler := handler.NewDishHandler(s.pStorage, s.Logger)
	r.Route("/api/restaurants", func(r chi.Router) {
		r.Post("/add", restHandler.Add)
		r.Get("/get", restHandler.Get)
		r.Delete("/del", restHandler.Del)
	})
	r.Route("/api/dish", func(r chi.Router) {
		r.Post("/add", dishHandler.Add)
		r.Get("/get", dishHandler.Get)
		r.Delete("/del", dishHandler.Del)
	})

	// настраиваем сервер
	srv := &http.Server{
		Addr:    s.HttpServer.Addr + ":" + strconv.Itoa(s.HttpServer.Port),
		Handler: r,
	}
	s.Info("Server initialized")
	// запускаем сервер
	go func() {
		s.Info("Server start! " +
			"Listening on " + s.HttpServer.Addr + ":" + strconv.Itoa(s.HttpServer.Port))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.Error(err.Error())
		}
	}()

	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		s.Error(err.Error())
	}
	s.Info("Server shutdown")
}

// todo метод заглушка
func todo(w http.ResponseWriter, r *http.Request) {
}
