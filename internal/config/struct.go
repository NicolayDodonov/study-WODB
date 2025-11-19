package config

import "time"

type Config struct {
	*HttpServer `yaml:"http-server"`
	*Postgres   `yaml:"postgres"`
	*Mongo      `yaml:"mongo"`
	*Logger     `yaml:"logger"`
}

type HttpServer struct {
	Addr            string        `yaml:"address"`
	Port            int           `yaml:"port"`
	Timeout         time.Duration `yaml:"timeout"`
	FileServer      string        `yaml:"path-file-server"`
	AuthRedirectURL string        `yaml:"google-redirect-url"`
	State           string        `yaml:"state"`
}

type Postgres struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

type Mongo struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Database string `yaml:"database"`
}

type Logger struct {
	Path string `yaml:"path"`
	Rang string `yaml:"rang"`
}
