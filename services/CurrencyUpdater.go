package services

import (
	"github.com/MathisBurger/crypto-simulator/database/actions"
	"strconv"
	"time"
)

func CurrencyUpdater() {

	for {
		for _ = range time.Tick(10 * time.Second) {
			obj := GetAllCurrencys()
			for i, el := range obj.Data {
				if i < 50 {
					rank, _ := strconv.Atoi(el.Rank)
					supply, _ := strconv.ParseFloat(el.Supply, 64)
					maxSupply, _ := strconv.ParseFloat(el.MaxSupply, 64)
					marketCapUSD, _ := strconv.ParseFloat(el.MarketCapUsd, 64)
					volumeUSD24Hr, _ := strconv.ParseFloat(el.VolumeUsd24Hr, 64)
					priceUSD, _ := strconv.ParseFloat(el.PriceUSD, 64)
					changePercent24Hr, _ := strconv.ParseFloat(el.ChangePercent24Hr, 64)
					vwap24Hr, _ := strconv.ParseFloat(el.Vwap24Hr, 64)
					if !actions.CheckIfCurrencyExists(el.Symbol) {
						actions.InsertCurrency(rank, el.Symbol, el.Name, supply, maxSupply, marketCapUSD, volumeUSD24Hr,
							priceUSD, changePercent24Hr, vwap24Hr)
					}
					actions.UpdateCurrency(rank, supply, maxSupply, marketCapUSD, volumeUSD24Hr, priceUSD, changePercent24Hr,
						vwap24Hr, el.Symbol)
					actions.InsertCurrencyChange(el.Symbol, priceUSD, int(time.Now().Unix()))
				} else {
					break
				}
			}
		}
	}
}
