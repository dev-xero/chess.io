package utils

import (
	"time"

	"github.com/gofiber/fiber/v3"
)

func Success(c fiber.Ctx, status int, message string, data any) error {
	return c.Status(status).JSON(fiber.Map{
		"success":   true,
		"message":   message,
		"data":      data,
		"timestamp": time.Now(),
	})
}

func Error(c fiber.Ctx, status int, message string) error {
	return c.Status(status).JSON(fiber.Map{
		"success":   false,
		"message":   message,
		"data":      nil,
		"timestamp": time.Now(),
	})
}
