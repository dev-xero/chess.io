//	@title			chess.io specification
//	@version		1.0
//	@description	endpoints for interacting with the backend

package main

import (
	"github.com/dev-xero/chess.io/internal/config"
	"github.com/dev-xero/chess.io/internal/database"
	"github.com/dev-xero/chess.io/internal/routes"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

type StructValidator struct {
	validate *validator.Validate
}

func (v *StructValidator) Validate(out any) error {
	return v.validate.Struct(out)
}

func main() {
	cfg := config.Load()
	db := database.Connect(cfg)

	app := fiber.New(fiber.Config{
		StructValidator: &StructValidator{
			validate: validator.New(),
		},
	})

	routes.Bootstrap(app, db, cfg)

	log.Fatal(app.Listen(":" + cfg.Port))
}
