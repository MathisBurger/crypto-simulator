package controller

import (
	"github.com/MathisBurger/crypto-simulator/database/actions"
	"github.com/MathisBurger/crypto-simulator/database/models"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type getCurrencyDataResponse struct {
	Status  bool                           `json:"status"`
	Message string                         `json:"message"`
	Data    []models.CurrencyProgressModel `json:"data"`
}

// ---------------------------------
//            DEPRECATED
//   This endpoint is deprecated
//   It can be enabled via config
// ---------------------------------
func GetCurrencyDataController(c *fiber.Ctx) error {

	username := c.Query("username")
	token := c.Query("token")
	name := c.Query("currency")
	timePeriod := c.Query("period")

	// check default values
	if username == "" || token == "" || name == "" || timePeriod == "" {
		return c.JSON(getCurrencyDataResponse{
			false,
			"Invalid JSON body",
			nil,
		})
	}

	// check login
	if actions.LoginWithToken(username, token) {

		// get int from timePeriod
		period, err := strconv.Atoi(timePeriod)

		if err != nil {
			return c.JSON(getCurrencyDataResponse{
				false,
				"Invalid period value",
				nil,
			})
		}

		return c.JSON(getCurrencyDataResponse{
			true,
			"successfully queried all currency data",
			actions.GetCurrencyData(name, period),
		})

	} else {
		return c.JSON(getCurrencyDataResponse{
			false,
			"wrong login credentials",
			nil,
		})
	}
}
