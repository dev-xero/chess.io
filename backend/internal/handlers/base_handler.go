package handlers

import (
	"net/http"

	"github.com/dev-xero/chess.io/internal/utils"
	"github.com/gofiber/fiber/v3"
)

type BaseHandler struct{}

// BaseHandler handles requests to the / base endpoint
//
//	@summary	Handle requests to the base endpoint
//	@tags		base
//	@accept		json
//	@produce	json
//	@success	200
//	@router		/ [get]
func (h *BaseHandler) Get(c fiber.Ctx) error {
	return utils.Success(
		c,
		http.StatusOK,
		"chess.io web server is running",
		fiber.Map{
			"metrics": "/metrics",
			"swagger": "/schema/swagger",
			"health":  "/health",
		},
	)
}

// BaseHandler Health handles requests to the /health endpoint
//
//	@summary	Handle requests to the health check endpoint
//	@tags		base
//	@accept		json
//	@produce	json
//	@success	200
//	@router		/health [get]
func (h *BaseHandler) Health(c fiber.Ctx) error {
	return utils.Success(
		c,
		http.StatusOK,
		"health check results",
		fiber.Map{
			"database": "healthy",
		},
	)
}
