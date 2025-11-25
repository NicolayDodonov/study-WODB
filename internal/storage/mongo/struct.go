package mongo

import (
	"fmt"
	"study-WODB/internal/config"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Storage struct {
	client *mongo.Client
}

func New(cnf *config.Mongo) (*Storage, error) {
	client, err := mongo.Connect(
		options.Client().ApplyURI(url(cnf)),
	)
	if err != nil {
		return nil, err
	}
	return &Storage{client: client}, err
}

func url(cnf *config.Mongo) string {
	return fmt.Sprintf("mongodb://%s:%d@%s",
		cnf.Host, cnf.Port)
}
