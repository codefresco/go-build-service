package config

import (
	"log"
	"sync"

	"github.com/caarlos0/env/v9"
)

type Config struct {
	Environment  string `env:"ENVIRONMENT,required"`
	LogFilePath  string `env:"LOG_FILE_PATH"`
	Port         string `env:"PORT" envDefault:"3000"`
	AllowOrigins string `env:"ALLOW_ORIGINS,required"`
	AllowHeaders string `env:"ALLOW_HEADERS,required"`
	Version      string `env:"VERSION,required"`
	PostgresURL  string `env:"POSTGRES_URL,required"`
}

var (
	cfg  *Config
	once sync.Once
)

func GetConfig() *Config {
	once.Do(func() {
		cfg = &Config{}
		if err := env.Parse(cfg); err != nil {
			log.Fatalf("could not load env: %v", err)
		}
	})
	return cfg
}
