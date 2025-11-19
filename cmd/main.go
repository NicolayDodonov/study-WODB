package main

import (
	"log"
	"study-WODB/internal/config"
	"study-WODB/internal/http/server"
	"study-WODB/internal/logger"
)

func main() {
	// инициализация
	c := config.MustLoad()
	l, err := logger.New(c)
	if err != nil {
		log.Fatal(err)
	}
	srv := server.New(c, l)

	// запуск http сервера
	srv.Start()
}
