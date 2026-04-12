package handlers

import (
	"net/http"

	"github.com/dev-xero/chess.io/internal/utils"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type AuthenticationHandler struct {
	DB *gorm.DB
}

// Authentication Handler - Register
//
//	@summary	Handles player registration requests
//	@tags		authentication
//	@accept		json
//	@produce	json
//	@success	201	{object}	models.Player
//	@router		/api/v1/auth/register [post]
func (h *AuthenticationHandler) Register(c fiber.Ctx) error {
	return utils.Success(
		c,
		http.StatusOK,
		"Successfully called the auth/register endpoint",
		nil,
	)
}
