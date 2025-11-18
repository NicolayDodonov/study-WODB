package config

import "time"

type Config struct {
	*HttpServer
	*Postgres
	*Mongo
	*Logger
}

type HttpServer struct {
	Addr    string        `yaml:"addr"`
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
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
