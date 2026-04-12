package handlers

import (
	"net/http"

	"github.com/dev-xero/chess.io/internal/dto"
	"github.com/dev-xero/chess.io/internal/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
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
//	@param		payload	body		dto.RegistrationRequest	true	"Registration Payload"
//	@success	201		{object}	models.Player
//	@router		/api/v1/auth/register [post]
func (h *AuthenticationHandler) Register(c fiber.Ctx) error {
	dto := new(dto.RegistrationRequest)
	err := c.Bind().All(dto)
	if err != nil {
		if validationErrs, ok := err.(validator.ValidationErrors); ok {
			log.Warn("Failed to bind registration request struct: ", err)
			for _, e := range validationErrs {
				return utils.Error(
					c,
					http.StatusBadRequest,
					"Request validation failed",
					fiber.Map{
						"field": e.Field(),
						"error": e.Error(),
					},
				)
			}
		}
		log.Error("Failed to bind registration request struct: ", err)
		return utils.Error(
			c,
			http.StatusBadRequest,
			"Cannot handle this request",
			nil,
		)
	}
	return utils.Success(
		c,
		http.StatusOK,
		"Successfully called the auth/register endpoint",
		nil,
	)
}
