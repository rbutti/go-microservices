package config

import (
	"github.com/joeshaw/envdecode"
	"log"
	"time"
)

type Config struct {
	Debug  bool `env:"DEBUG,required"`
	Server serverConfig
	Db     dbConfig
}

type serverConfig struct {
	Port         int           `env:"PORT,required"`
	TimeoutRead  time.Duration `env:"READ_TIMEOUT,required"`
	TimeoutWrite time.Duration `env:"WRITE_TIMEOUT,required"`
	TimeoutIdle  time.Duration `env:"IDLE_TIMEOUT,required"`
}

func AppConfig() *Config {
	var c Config
	if err := envdecode.StrictDecode(&c); err != nil {
		log.Fatalf("Failed to decode: %s", err)
	}

	return &c
}

type dbConfig struct {
	Host     string `env:"DB_HOST,required"`
	Port     int    `env:"DB_PORT,required"`
	Username string `env:"DB_USER,required"`
	Password string `env:"DB_PASSWORD,required"`
	DbName   string `env:"DB_NAME,required"`
}
