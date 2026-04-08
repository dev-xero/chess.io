package database

import (
	"fmt"

	"github.com/dev-xero/chess.io/internal/config"
	"github.com/dev-xero/chess.io/internal/models"
	"github.com/gofiber/fiber/v3/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(cfg *config.Config) *gorm.DB {
	datasource_name := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.DBHost,
		cfg.DBUser,
		cfg.DBPass,
		cfg.DBName,
		cfg.DBPort,
		cfg.DBSSLMode,
	)

	db, err := gorm.Open(postgres.Open(datasource_name), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect postgres:", err)
	}
	log.Info("Postgres connected successfully")

	err = db.AutoMigrate(
		&models.Country{},
		&models.Player{},
	)
	if err != nil {
		log.Fatal("Failed to perform database migrations: ", err)
	}

	return db
}
