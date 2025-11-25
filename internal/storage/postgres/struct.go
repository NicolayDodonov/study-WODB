package postgres

import (
	"fmt"
	"study-WODB/internal/config"

	"github.com/jmoiron/sqlx"
)

type Storage struct {
	db *sqlx.DB
}

func New(cnf *config.Postgres) (*Storage, error) {
	db, err := sqlx.Connect("postgres", dsn(cnf))
	if err != nil {
		return nil, err
	}
	return &Storage{db: db}, nil
}

func dsn(cnf *config.Postgres) string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cnf.Host, cnf.Port, cnf.User, cnf.Password, cnf.Database, "disable")
}
