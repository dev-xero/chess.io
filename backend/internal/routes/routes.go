package routes

import (
	_ "github.com/dev-xero/chess.io/docs"
	"github.com/dev-xero/chess.io/internal/config"
	"github.com/dev-xero/chess.io/internal/handlers"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/adaptor"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"gorm.io/gorm"
)

func Bootstrap(app *fiber.App, db *gorm.DB, cfg *config.Config) {
	rootHandler := &handlers.RootHandler{}

	api := app.Group("/api/v1")

	api.Get("/", rootHandler.Get)
	api.Get("/swagger/*", adaptor.HTTPHandler(httpSwagger.Handler()))
}
