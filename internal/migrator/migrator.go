package migrator

import (
	"study-WODB/internal/config"
	"study-WODB/internal/storage/mongo"
	"study-WODB/internal/storage/postgres"
)

type Migrator struct {
	p     *postgres.Storage
	m     *mongo.Storage
	pPath string //путь до файла миграции postgres.sql
	mPath string //путь до файла миграции для mongo.db
}

func New(c *config.Config, p *postgres.Storage, m *mongo.Storage) *Migrator {
	return &Migrator{}
}

func (m *Migrator) Migrate() {
	//todo: выполнить запросы
}
