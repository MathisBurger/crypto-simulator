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
	raw := string(c.Body())
	obj := sellCryptoRequest{}
	err := json.Unmarshal([]byte(raw), &obj)
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
	if actions.LoginWithToken(obj.Username, obj.Token) {
		user := actions.GetUserByUsername(obj.Username)
		currency := actions.GetCurrency(obj.CurrencyID)
		if actions.GetAmountOfCryptoByUUID(user.WalletUUID, currency.Symbol) >= obj.Amount {
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

func checkSellCryptoRequest(obj sellCryptoRequest) bool {
	return obj.Username != "" && obj.Token != "" && obj.CurrencyID != "" && obj.Amount > 0
}
