package controller

import (
	"encoding/json"
	"github.com/MathisBurger/crypto-simulator/database/actions"
	"github.com/gofiber/fiber/v2"
)

type buyCryptoRequest struct {
	Username   string  `json:"username"`
	Token      string  `json:"token"`
	CurrencyID string  `json:"currency_id"`
	Amount     float64 `json:"amount"`
}

type buyCryptoResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

func BuyCryptoController(c *fiber.Ctx) error {

	// parsing and checking request
	obj := buyCryptoRequest{}
	err := json.Unmarshal(c.Body(), &obj)
	if err != nil {
		return c.JSON(buyCryptoResponse{
			false,
			"Invalid JSON body",
		})
	}
	if !checkBuyCryptoRequest(obj) {
		return c.JSON(buyCryptoResponse{
			false,
			"Invalid JSON body",
		})
	}

	// amount must be greater than 0
	if obj.Amount <= 0 {
		return c.JSON(sellCryptoResponse{
			false,
			"Value must be higher than zero",
		})
	}

	// check login
	if actions.LoginWithToken(obj.Username, obj.Token) {

		currency := actions.GetCurrency(obj.CurrencyID)

		// check if currency exists in database
		if currency.CoinID == obj.CurrencyID {

			user := actions.GetUserByUsername(obj.Username)

			// check if user has enough money
			if actions.GetWalletByUUID(user.WalletUUID).BalanceUSD > currency.PriceUSD*obj.Amount {

				// buy crypto
				actions.AddCryptoToWallet(user.WalletUUID, currency.Symbol, obj.Amount)
				actions.RemoveMoneyFromWallet(user.WalletUUID, currency.PriceUSD*obj.Amount)
				actions.AddTrade("USD", currency.Symbol, user.WalletUUID, currency.PriceUSD, obj.Amount)

				return c.JSON(buyCryptoResponse{
					true,
					"successfully traded currencys",
				})
			} else {
				return c.JSON(buyCryptoResponse{
					false,
					"You do not have enough money in your Wallet",
				})
			}
		} else {
			return c.JSON(buyCryptoResponse{
				false,
				"This cryptocurrency does not exist in our database",
			})
		}

	} else {
		return c.JSON(buyCryptoResponse{
			false,
			"Wrong login credentials",
		})
	}
}

// checks request
func checkBuyCryptoRequest(obj buyCryptoRequest) bool {
	return obj.Username != "" && obj.Token != "" && obj.CurrencyID != ""
}
