package models

import "database/sql"

type CurrencyModel struct {
	ID          int
	Name        string
	Price       float64
	LastUpdated int
}

func (c CurrencyModel) Parse(resp *sql.Rows) CurrencyModel {
	var mdl CurrencyModel
	err := resp.Scan(&mdl.ID, &mdl.Name, &mdl.Price, &mdl.LastUpdated)
	if err != nil {
		panic(err.Error())
	}
	return mdl
}

func (c CurrencyModel) ParseAll(resp *sql.Rows) []CurrencyModel {
	var mdls []CurrencyModel
	for resp.Next() {
		var mdl CurrencyModel
		err := resp.Scan(&mdl.ID, &mdl.Name, &mdl.Price, &mdl.LastUpdated)
		if err != nil {
			panic(err.Error())
		}
		mdls = append(mdls, mdl)
	}
	return mdls
}
