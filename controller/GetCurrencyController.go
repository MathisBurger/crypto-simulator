package controller

import (
	"github.com/MathisBurger/crypto-simulator/database/actions"
	"github.com/MathisBurger/crypto-simulator/database/models"
	"github.com/MathisBurger/crypto-simulator/middleware"
	"github.com/gofiber/fiber/v2"
)

type getCurrencyResponse struct {
	Status  bool                 `json:"status"`
	Message string               `json:"message"`
	Data    models.CurrencyModel `json:"data"`
}

func GetCurrencyController(c *fiber.Ctx) error {

	name := c.Query("currency")

	// check default values
	if name == "" {
		return c.JSON(getCurrencyResponse{
			false,
			"Invalid JSON body",
			models.CurrencyModel{},
		})
	}

	// check login
	if status, _ := middleware.ValidateAccessToken(c); status {
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
