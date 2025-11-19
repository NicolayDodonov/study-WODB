package config

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

// MustLoad читает файл конфигурации в структуру Config.
// При любой ошибки чтения вызывает панику.
func MustLoad() *Config {
	path := os.Getenv("Config_Path")

	if path == "" {
		log.Fatal("config file path is empty")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Fatalf("config file %s does not exist", path)
	}

	var config Config

	if err := cleanenv.ReadConfig(path, &config); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &config
}
