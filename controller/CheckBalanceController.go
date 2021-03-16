package controller

import (
	"github.com/MathisBurger/crypto-simulator/database/actions"
	"github.com/MathisBurger/crypto-simulator/middleware"
	"github.com/gofiber/fiber/v2"
)

type checkBalanceResponse struct {
	Status  bool    `json:"status"`
	Message string  `json:"message"`
	Balance float64 `json:"balance"`
}

func CheckBalanceController(c *fiber.Ctx) error {

	// check login
	if status, ident := middleware.ValidateAccessToken(c); status {

		UUID := actions.GetUserByUsername(ident).WalletUUID

		return c.JSON(checkBalanceResponse{
			true,
			"successfully queried balance",
			actions.GetWalletByUUID(UUID).BalanceUSD,
		})
	} else {
		return c.JSON(checkBalanceResponse{
			false,
			"wrong login credentials",
			0,
		})
	}
}
