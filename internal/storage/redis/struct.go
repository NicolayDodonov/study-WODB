package redis

import (
	"strconv"
	"study-WODB/internal/config"

	"github.com/redis/go-redis/v9"
)

type Storage struct {
	client *redis.Client
}

func New(cnf *config.Redis) *Storage {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cnf.Host + ":" + strconv.Itoa(cnf.Port),
		Username: cnf.User,
		Password: cnf.Password,
		DB:       cnf.Database,
	})
	return &Storage{client: rdb}
}
