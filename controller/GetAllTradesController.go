package controller

import (
	"github.com/MathisBurger/crypto-simulator/database/actions"
	"github.com/MathisBurger/crypto-simulator/database/models"
	"github.com/gofiber/fiber/v2"
)

type getAllTradesResponse struct {
	Status  bool                 `json:"status"`
	Message string               `json:"message"`
	Data    []models.TradesModel `json:"data"`
}

func GetAllTradesController(c *fiber.Ctx) error {

	username := c.Query("username")
	token := c.Query("token")

	// check default values
	if username == "" || token == "" {
		return c.JSON(getAllTradesResponse{
			false,
			"Invalid JSON body",
			nil,
		})
	}

	// check login
	if actions.LoginWithToken(username, token) {
		return c.JSON(getAllTradesResponse{
			true,
			"successfully queried all trades from the last 7 days",
			actions.GetAllTradesForUser(actions.GetUserByUsername(username).WalletUUID),
		})

	} else {
		return c.JSON(getAllTradesResponse{
			false,
			"Wrong login credentials",
			nil,
		})
	}
}
