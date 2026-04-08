package config

import (
	"os"

	"github.com/gofiber/fiber/v3/log"
	"github.com/joho/godotenv"
)

type Config struct {
	Port      string
	DBHost    string
	DBPort    string
	DBUser    string
	DBPass    string
	DBName    string
	DBSSLMode string
}

func Load() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	log.Info("Environment variables loaded")

	return &Config{
		Port:      os.Getenv("PORT"),
		DBHost:    os.Getenv("POSTGRES_HOST"),
		DBPort:    os.Getenv("POSTGRES_PORT"),
		DBUser:    os.Getenv("POSTGRES_USER"),
		DBPass:    os.Getenv("POSTGRES_PASS"),
		DBName:    os.Getenv("POSTGRES_NAME"),
		DBSSLMode: os.Getenv("POSTGRES_SSLMODE"),
	}
}
