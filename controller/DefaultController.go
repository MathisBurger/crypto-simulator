package controller

import (
	"github.com/gofiber/fiber/v2"
	"os"
)

type DefaultResponse struct {
	Message     string `json:"message"`
	Status      string `json:"status"`
	RateLimiter string `json:"rate_limiter"`
}

func DefaultController(c *fiber.Ctx) error {
	return c.JSON(DefaultResponse{
		"crypto-service is running...",
		"ok",
		os.Getenv("RATE_LIMITER"),
	})
}
