package actions

import "github.com/MathisBurger/crypto-simulator/database/models"

func GetCurrency(name string) models.CurrencyModel {
	conn, _ := connect()
	defer conn.Close()
	stmt, err := conn.Prepare("SELECT * FROM `currencys` WHERE `CoinID`=?")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	resp, err := stmt.Query(name)
	if err != nil {
		panic(err.Error())
	}
	defer resp.Close()
	if resp.Next() {
		return models.CurrencyModel{}.Parse(resp)
	} else {
		return models.CurrencyModel{}
	}
}
