package controller

import (
	"github.com/MathisBurger/crypto-simulator/database/actions"
	"github.com/gofiber/fiber/v2"
)

type checkBalanceRequest struct {
	Username  string `json:"username"`
	AuthToken string `json:"auth_token"`
}

type checkBalanceResponse struct {
	Status  bool    `json:"status"`
	Message string  `json:"message"`
	Balance float64 `json:"balance"`
}

func CheckBalanceController(c *fiber.Ctx) error {

	// parse get params to object
	obj := checkBalanceRequest{
		c.Query("username", ""),
		c.Query("auth_token", ""),
	}

	// check default values
	if !checkCheckBalanceRequest(obj) {
		return c.JSON(checkBalanceResponse{
			false,
			"Invalid JSON body",
			0,
		})
	}

	// check login
	if actions.LoginWithToken(obj.Username, obj.AuthToken) {

		UUID := actions.GetUserByUsername(obj.Username).WalletUUID

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

// checks request
func checkCheckBalanceRequest(obj checkBalanceRequest) bool {
	return obj.Username != "" && obj.AuthToken != ""
}
