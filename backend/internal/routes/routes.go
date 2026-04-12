package routes

import (
	_ "github.com/dev-xero/chess.io/docs"
	"github.com/dev-xero/chess.io/internal/config"
	"github.com/dev-xero/chess.io/internal/handlers"
	"github.com/dev-xero/chess.io/internal/middleware"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/adaptor"
	promhttp "github.com/prometheus/client_golang/prometheus/promhttp"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"gorm.io/gorm"
)

func Bootstrap(app *fiber.App, db *gorm.DB, cfg *config.Config) {
	baseHandler := &handlers.BaseHandler{
		DB: db,
	}
	rootHandler := &handlers.RootHandler{}

	middleware.PromInit()
	app.Use(middleware.TrackMetrics())

	app.Get("/", baseHandler.Get)
	app.Get("/health", baseHandler.Health)
	app.Get("/metrics", adaptor.HTTPHandler(promhttp.Handler()))
	app.Get("/schema/swagger/*", adaptor.HTTPHandler(httpSwagger.Handler()))

	api := app.Group("/api/v1")
	api.Get("/", rootHandler.Get)

	authHandler := &handlers.AuthenticationHandler{}
	apiAuthEndpoints := api.Group("/auth")
	apiAuthEndpoints.Post("/register", authHandler.Register)
}
