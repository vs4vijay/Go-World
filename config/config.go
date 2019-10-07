package config

import (
	"log"

	"github.com/joeshaw/envdecode"
)

type Config struct {
	Server serverConfig
}

type serverConfig struct {
	Port         int           `env:"SERVER_PORT,required"`
}

func AppConfig() *Config {
	var config Config
	if err := envdecode.StrictDecode(&config); err != nil {
		log.Fatalf("Failed to decode: %s", err)
	}

	return &config
}