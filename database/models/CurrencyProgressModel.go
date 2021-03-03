package models

import "database/sql"

type CurrencyProgressModel struct {
	ID        int
	Currency  string
	PriceUSD  float64
	Timestamp int
}

// parse single value
func (c CurrencyProgressModel) Parse(resp *sql.Rows) CurrencyProgressModel {
	var cache CurrencyProgressModel
	err := resp.Scan(&cache.ID, &cache.Currency, &cache.PriceUSD, &cache.Timestamp)
	if err != nil {
		panic(err.Error())
	}
	return cache
}

// parse all
func (c CurrencyProgressModel) ParseAll(resp *sql.Rows) []CurrencyProgressModel {
	var answers []CurrencyProgressModel
	for resp.Next() {
		var cache CurrencyProgressModel
		err := resp.Scan(&cache.ID, &cache.Currency, &cache.PriceUSD, &cache.Timestamp)
		if err != nil {
			panic(err.Error())
		}
		answers = append(answers, cache)
	}
	return answers
}
