package config

import (
	"github.com/joeshaw/envdecode"
	"log"
	"time"
)

type Conf struct {
	Server serverConf
}

type serverConf struct {
	Port         int           `env:"PORT,required"`
	TimeoutRead  time.Duration `env:"READ_TIMEOUT,required"`
	TimeoutWrite time.Duration `env:"WRITE_TIMEOUT,required"`
	TimeoutIdle  time.Duration `env:"IDLE_TIMEOUT,required"`
}

func AppConfig() *Conf {
	var c Conf
	if err := envdecode.StrictDecode(&c); err != nil {
		log.Fatalf("Failed to decode: %s", err)
	}

	return &c
}
