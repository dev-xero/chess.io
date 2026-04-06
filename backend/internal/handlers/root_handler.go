package handlers

import (
	"net/http"

	"github.com/dev-xero/chess.io/internal/utils"
	"github.com/gofiber/fiber/v3"
)

type RootHandler struct{}

func (h *RootHandler) Get(c fiber.Ctx) error {
	return utils.Success(c, http.StatusOK, "API v1.0 is online", "visit the swagger specification with /schema/swagger")
}
