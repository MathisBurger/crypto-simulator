package services

import (
	"github.com/MathisBurger/crypto-simulator/database/actions"
	"math"
	"strconv"
	"time"
)

func CurrencyUpdater() {

	for {
		for _ = range time.Tick(10 * time.Second) {
			obj := GetAllCurrencys()
			for i, el := range obj.Data {
				if i < 50 {
					price, _ := strconv.ParseFloat(el.PriceUSD, 64)
					if !actions.CheckIfCurrencyExists(el.Symbol) {
						actions.InsertCurrency(el.Symbol, price, int(math.Round(float64(obj.Timestamp/1000))))
					}
					actions.UpdateCurrency(el.Symbol, price, int(math.Round(float64(obj.Timestamp/1000))))
					actions.InsertCurrencyChange(el.Symbol, price, int(math.Round(float64(obj.Timestamp/1000))))
				} else {
					break
				}
			}
		}
	}
}
