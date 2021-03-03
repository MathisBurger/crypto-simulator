package controller

import (
	"github.com/MathisBurger/crypto-simulator/database/actions"
	"github.com/MathisBurger/crypto-simulator/database/models"
	"github.com/gofiber/fiber/v2"
)

type getCryptoWalletsForUserResponse struct {
	Status  bool                        `json:"status"`
	Message string                      `json:"message"`
	Data    []models.CurrencyArrayModel `json:"data"`
}

func GetCryptoWalletsForUser(c *fiber.Ctx) error {

	username := c.Query("username")
	token := c.Query("token")

	// check default values
	if username == "" || token == "" {
		return c.JSON(getCryptoWalletsForUserResponse{
			false,
			"Invalid JSON body",
			nil,
		})
	}

	// check login
	if actions.LoginWithToken(username, token) {
		return c.JSON(getCryptoWalletsForUserResponse{
			true,
			"successfully queried all trades from the last 7 days",
			actions.GetCurrencyArray(actions.GetUserByUsername(username).WalletUUID),
		})

	} else {
		return c.JSON(getCryptoWalletsForUserResponse{
			false,
			"Wrong login credentials",
			nil,
		})
	}
}
