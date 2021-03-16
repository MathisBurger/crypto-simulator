package controller

import (
	"github.com/MathisBurger/crypto-simulator/database/actions"
	"github.com/MathisBurger/crypto-simulator/database/models"
	"github.com/MathisBurger/crypto-simulator/middleware"
	"github.com/gofiber/fiber/v2"
)

type getAllCurrencysRespose struct {
	Status  bool                   `json:"status"`
	Message string                 `json:"message"`
	Data    []models.CurrencyModel `json:"data"`
}

func GetAllCurrencysController(c *fiber.Ctx) error {

	// check login
	if status, _ := middleware.ValidateAccessToken(c); status {
		return c.JSON(getAllCurrencysRespose{
			true,
			"successfully queried all currencys",
			actions.GetAllCurrencys(),
		})

	} else {
		return c.JSON(getAllCurrencysRespose{
			false,
			"Wrong login credentials",
			nil,
		})
	}
}
