package models

import "database/sql"

type TradesModel struct {
	ID         int     `json:"id"`
	Trade      string  `json:"trade"`
	WalletUUID string  `json:"wallet_uuid"`
	CoinPrice  float64 `json:"coin_price"`
	Amount     float64 `json:"amount"`
	Timestamp  int     `json:"timestamp"`
}

// parse single value
func (c TradesModel) Parse(resp *sql.Rows) TradesModel {
	var cache TradesModel
	_ = resp.Scan(&cache.ID, &cache.Trade, &cache.WalletUUID, &cache.CoinPrice, &cache.Amount, &cache.Timestamp)
	return cache
}

// parse all
func (c TradesModel) ParseAll(resp *sql.Rows) []TradesModel {
	var answers []TradesModel
	for resp.Next() {
		var cache TradesModel
		_ = resp.Scan(&cache.ID, &cache.Trade, &cache.WalletUUID, &cache.CoinPrice, &cache.Amount, &cache.Timestamp)
		answers = append(answers, cache)
	}
	return answers
}
