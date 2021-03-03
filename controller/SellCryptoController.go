package controller

import (
	"encoding/json"
	"github.com/MathisBurger/crypto-simulator/database/actions"
	"github.com/gofiber/fiber/v2"
)

type sellCryptoRequest struct {
	Username   string  `json:"username"`
	Token      string  `json:"token"`
	CurrencyID string  `json:"currency_id"`
	Amount     float64 `json:"amount"`
}

type sellCryptoResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

func SellCryptoController(c *fiber.Ctx) error {

	// parsing and checking request
	obj := sellCryptoRequest{}
	err := json.Unmarshal(c.Body(), &obj)
	if err != nil {
		return c.JSON(sellCryptoResponse{
			false,
			"Invalid JSON body",
		})
	}
	if !checkSellCryptoRequest(obj) {
		return c.JSON(sellCryptoResponse{
			false,
			"Invalid JSON body",
		})
	}

	// must be minimum 0
	if obj.Amount <= 0 {
		return c.JSON(sellCryptoResponse{
			false,
			"Value must be higher than zero",
		})
	}

	// check login
	if actions.LoginWithToken(obj.Username, obj.Token) {

		user := actions.GetUserByUsername(obj.Username)
		currency := actions.GetCurrency(obj.CurrencyID)

		// check if user own given number of crypto
		if actions.GetAmountOfCryptoByUUID(user.WalletUUID, currency.Symbol) >= obj.Amount {

			// sell crypto
			actions.RemoveCryptoFromWallet(user.WalletUUID, currency.Symbol, obj.Amount)
			actions.AddMoneyToWallet(user.WalletUUID, currency.PriceUSD*obj.Amount)
			actions.AddTrade(currency.Symbol, "USD", user.WalletUUID, currency.PriceUSD, obj.Amount)

			return c.JSON(sellCryptoResponse{
				true,
				"Successfully sold crypto currency",
			})

		} else {
			return c.JSON(sellCryptoResponse{
				false,
				"You do not have enough crypto in your Wallet",
			})
		}

	} else {
		return c.JSON(sellCryptoResponse{
			false,
			"Wrong login credentials",
		})
	}
}

// checks request
func checkSellCryptoRequest(obj sellCryptoRequest) bool {
	return obj.Username != "" && obj.Token != "" && obj.CurrencyID != ""
}
