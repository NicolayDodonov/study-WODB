package main

import (
	"log"
	"study-WODB/internal/config"
	"study-WODB/internal/http/server"
	"study-WODB/internal/logger"
	"study-WODB/internal/storage/mongo"
	"study-WODB/internal/storage/postgres"
	"study-WODB/internal/storage/redis"
)

func main() {
	// инициализация
	c := config.MustLoad()
	l, err := logger.New(c)
	if err != nil {
		log.Fatal(err)
	}
	// подключение к бд и применение миграций
	ps, err := postgres.New(c.Postgres)
	if err != nil {
		l.Fatal("Can`t init Postgres:" + err.Error())
	}
	md, err := mongo.New(c.Mongo)
	if err != nil {
		l.Fatal("Can`t init Mongo:" + err.Error())
	}
	rd := redis.New(c.Redis)

	// создание http сервера
	srv := server.New(c, l, ps, md, rd)

	// запуск http сервера
	srv.Start()
}
