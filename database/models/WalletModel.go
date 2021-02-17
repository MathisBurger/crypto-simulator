package models

import "database/sql"

type WalletModel struct {
	ID            int     `json:"id"`
	UUID          string  `json:"UUID"`
	BalanceUSD    float64 `json:"balanceUSD"`
	CurrencyArray string  `json:"currencyArray"`
}

func (c WalletModel) Parse(resp *sql.Rows) WalletModel {
	var cache WalletModel
	_ = resp.Scan(&cache.ID, &cache.UUID, &cache.BalanceUSD, &cache.CurrencyArray)
	return cache
}
