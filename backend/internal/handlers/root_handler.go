package handlers

import (
	"net/http"

	"github.com/dev-xero/chess.io/internal/utils"
	"github.com/gofiber/fiber/v3"
)

type RootHandler struct{}

// RootHandler handles requests to the /api/v1 endpoint
//
//	@summary	Handle requests to the base endpoint
//	@tags		base
//	@accept		json
//	@produce	json
//	@success	200
//	@router		/api/v1/ [get]
func (h *RootHandler) Get(c fiber.Ctx) error {
	return utils.Success(
		c,
		http.StatusOK,
		"API v1.0 is online",
		"visit the swagger specification at api/v1/swagger/",
	)
}
