package controller

import (
	"github.com/MathisBurger/crypto-simulator/database/actions"
	"github.com/gofiber/fiber/v2"
)

type getTokenStatusResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Valid   bool   `json:"valid"`
}

func GetTokenStatusController(c *fiber.Ctx) error {
	username := c.Query("username")
	token := c.Query("token")
	if username == "" || token == "" {
		return c.JSON(getTokenStatusResponse{
			false,
			"Invalid JSON body",
			false,
		})
	}
	if actions.LoginWithToken(username, token) {
		return c.JSON(getTokenStatusResponse{
			true,
			"login credentials are valid",
			true,
		})
	} else {
		return c.JSON(getTokenStatusResponse{
			true,
			"wrong login credentials",
			false,
		})
	}
}
