package main

import (
	"github.com/dev-xero/chess.io/internal/config"
	"github.com/dev-xero/chess.io/internal/database"
	"github.com/dev-xero/chess.io/internal/routes"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

func main() {
	cfg := config.Load()
	db := database.Connect(cfg)
	app := fiber.New()
	routes.Bootstrap(app, db, cfg)
	log.Fatal(app.Listen(":8080"))
}
