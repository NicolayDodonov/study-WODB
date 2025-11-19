package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"study-WODB/internal/config"
	"study-WODB/internal/http/auth"
	"study-WODB/internal/logger"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

// HttpServer - структура описывающая работу сервера
type HttpServer struct {
	*config.Config
	*logger.Logger
}

func New(cnf *config.Config, log *logger.Logger) *HttpServer {
	return &HttpServer{
		cnf,
		log,
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
	r.Route("/", func(r chi.Router) {
		http.FileServer(http.Dir(s.FileServer))
	})

	// настройка ресурсов аутентификации
	s.Debug(" - configuring authentication resources")
	gAuth := auth.New(s.Config)
	r.Route("/auth", func(r chi.Router) {
		r.Post("/local", todo)                       // стандартный аутентификатор
		r.Get("/google", gAuth.GoogleCall)           // Oauth2 аутентификатор от google
		r.Post("/google-callback", gAuth.GoogleBack) // переадресация назад
	})

	// настройка graphQL точки входа
	s.Debug(" - configuring GraphQL entry points")
	r.Route("/menu", func(r chi.Router) {

	})
	r.Route("/order", func(r chi.Router) {

	})

	// настройка методов api
	s.Debug(" - configuring application resources")
	r.Route("/api", func(r chi.Router) {

	})

	// настраиваем сервер
	srv := &http.Server{
		Addr:    s.HttpServer.Addr + strconv.Itoa(s.HttpServer.Port),
		Handler: r,
	}
	s.Info("Server initialized")
	// запускаем сервер
	go func() {
		s.Info("Server start!")
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
