package routes

import (
	"github.com/dev-xero/chess.io/internal/config"
	"github.com/dev-xero/chess.io/internal/handlers"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func Bootstrap(app *fiber.App, db *gorm.DB, cfg *config.Config) {
	rootHandler := &handlers.RootHandler{}

	api := app.Group("/api/v1")
	api.Get("/", rootHandler.Get)
}
