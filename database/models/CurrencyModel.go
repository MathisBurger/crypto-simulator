package models

import "database/sql"

type CurrencyModel struct {
	ID                int     `json:"id"`
	CoinID            string  `json:"coin_id"`
	Rank              int     `json:"rank"`
	Symbol            string  `json:"symbol"`
	Name              string  `json:"name"`
	Supply            float64 `json:"supply"`
	MaxSupply         float64 `json:"max_supply"`
	MarketCapUSD      float64 `json:"market_cap_usd"`
	VolumeUSD24Hr     float64 `json:"volume_usd_24_hr"`
	PriceUSD          float64 `json:"price_usd"`
	ChangePercent24Hr float64 `json:"change_percent_24_hr"`
	Vwap24Hr          float64 `json:"vwap_24_hr"`
}

// parse single value
func (c CurrencyModel) Parse(resp *sql.Rows) CurrencyModel {
	var mdl CurrencyModel
	err := resp.Scan(&mdl.ID, &mdl.CoinID, &mdl.Rank, &mdl.Symbol, &mdl.Name, &mdl.Supply, &mdl.MaxSupply, &mdl.MarketCapUSD, &mdl.VolumeUSD24Hr, &mdl.PriceUSD, &mdl.ChangePercent24Hr, &mdl.Vwap24Hr)
	if err != nil {
		panic(err.Error())
	}
	return mdl
}

// parse all
func (c CurrencyModel) ParseAll(resp *sql.Rows) []CurrencyModel {
	var mdls []CurrencyModel
	for resp.Next() {
		var mdl CurrencyModel
		err := resp.Scan(&mdl.ID, &mdl.CoinID, &mdl.Rank, &mdl.Symbol, &mdl.Name, &mdl.Supply, &mdl.MaxSupply, &mdl.MarketCapUSD, &mdl.VolumeUSD24Hr, &mdl.PriceUSD, &mdl.ChangePercent24Hr, &mdl.Vwap24Hr)
		if err != nil {
			panic(err.Error())
		}
		mdls = append(mdls, mdl)
	}
	return mdls
}
