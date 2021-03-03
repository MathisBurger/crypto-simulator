package controller

import (
	"github.com/MathisBurger/crypto-simulator/database/actions"
	"github.com/MathisBurger/crypto-simulator/database/models"
	"github.com/gofiber/fiber/v2"
)

type getCurrencyResponse struct {
	Status  bool                 `json:"status"`
	Message string               `json:"message"`
	Data    models.CurrencyModel `json:"data"`
}

func GetCurrencyController(c *fiber.Ctx) error {

	username := c.Query("username")
	token := c.Query("token")
	name := c.Query("currency")

	// check default values
	if username == "" || token == "" || name == "" {
		return c.JSON(getCurrencyResponse{
			false,
			"Invalid JSON body",
			models.CurrencyModel{},
		})
	}

	// check login
	if actions.LoginWithToken(username, token) {
		return c.JSON(getCurrencyResponse{
			true,
			"successfully queried all currencys",
			actions.GetCurrency(name),
		})

	} else {
		return c.JSON(getCurrencyResponse{
			false,
			"Wrong login credentials",
			models.CurrencyModel{},
		})
	}
}
