package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	Http
	Storage
}

type Http struct {
	Port string
}

type Storage struct {
	Local
}

type Local struct {
	Path string
}

func NewConfig() *Config {
	return &Config{
		Http: Http{
			Port: os.Getenv("PORT"),
		},
		Storage: Storage{
			Local: Local{
				Path: os.Getenv("STORAGE_PATH"),
			},
		},
	}
}
