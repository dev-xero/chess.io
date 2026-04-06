package main

import (
	"github.com/dev-xero/chess.io/internal/config"
	"github.com/dev-xero/chess.io/internal/database"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

func main() {
	cfg := config.Load()
	database.Connect(cfg)
	app := fiber.New()

	log.Fatal(app.Listen(":8080"))
}
