package controller

import (
	"github.com/MathisBurger/crypto-simulator/database/actions"
	"github.com/MathisBurger/crypto-simulator/database/models"
	"github.com/MathisBurger/crypto-simulator/middleware"
	"github.com/gofiber/fiber/v2"
)

type getCryptoWalletsForUserResponse struct {
	Status  bool                        `json:"status"`
	Message string                      `json:"message"`
	Data    []models.CurrencyArrayModel `json:"data"`
}

func GetCryptoWalletsForUser(c *fiber.Ctx) error {

	// check login
	if status, ident := middleware.ValidateAccessToken(c); status {
		return c.JSON(getCryptoWalletsForUserResponse{
			true,
			"successfully queried all trades from the last 7 days",
			actions.GetCurrencyArray(actions.GetUserByUsername(ident).WalletUUID),
		})

	} else {
		return c.JSON(getCryptoWalletsForUserResponse{
			false,
			"Wrong login credentials",
			nil,
		})
	}
}
