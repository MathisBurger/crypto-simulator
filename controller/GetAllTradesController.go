package controller

import (
	"github.com/MathisBurger/crypto-simulator/database/actions"
	"github.com/MathisBurger/crypto-simulator/database/models"
	"github.com/MathisBurger/crypto-simulator/middleware"
	"github.com/gofiber/fiber/v2"
)

type getAllTradesResponse struct {
	Status  bool                 `json:"status"`
	Message string               `json:"message"`
	Data    []models.TradesModel `json:"data"`
}

func GetAllTradesController(c *fiber.Ctx) error {

	// check login
	if status, ident := middleware.ValidateAccessToken(c); status {
		return c.JSON(getAllTradesResponse{
			true,
			"successfully queried all trades from the last 7 days",
			actions.GetAllTradesForUser(actions.GetUserByUsername(ident).WalletUUID),
		})

	} else {
		return c.JSON(getAllTradesResponse{
			false,
			"Wrong login credentials",
			nil,
		})
	}
}
