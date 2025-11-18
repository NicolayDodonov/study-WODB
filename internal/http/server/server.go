package server

import (
	"net/http"
	"study-WODB/internal/config"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

// HttpServer - структура описывающая работу сервера
type HttpServer struct {
}

func New() *HttpServer {
	return &HttpServer{}
}

func (s *HttpServer) Start(cnf *config.Config) {

	// создаём и настраиваем роутер
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	}))

	// ===- определение путей и end-point`ов -=== \\

	// настройка файлового сервера
	r.Route("/", func(r chi.Router) {
		http.FileServer(http.Dir(cnf.FileServer))
	})

	// настройка ресурсов аутентификации
	r.Route("/auth", func(r chi.Router) {
		r.Post("/local", todo)  // стандартный аутентификатор
		r.Post("/google", todo) // Oauth2 аутентификатор от google
	})

	// настройка graphQL точки входа
	r.Route("/menu", func(r chi.Router) {

	})
	r.Route("/order", func(r chi.Router) {

	})

	// настройка методов api
	r.Route("/api", func(r chi.Router) {

	})

	// настраиваем сервер
	srv := &http.Server{
		Addr:    cnf.HttpServer.Addr + string(cnf.HttpServer.Port),
		Handler: r,
	}

	// запускаем сервер
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			//todo: log error
		}
	}()
}

// todo метод заглушка
func todo(w http.ResponseWriter, r *http.Request) {
}
