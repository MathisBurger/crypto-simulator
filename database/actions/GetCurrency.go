package actions

import "github.com/MathisBurger/crypto-simulator/database/models"

func GetCurrency(name string) models.CurrencyModel {

	conn, _ := Connect()
	defer conn.Close()

	stmt, _ := conn.Prepare("SELECT * FROM `currencys` WHERE `CoinID`=?")
	defer stmt.Close()

	resp, _ := stmt.Query(name)
	defer resp.Close()

	// check if exists
	if resp.Next() {
		return models.CurrencyModel{}.Parse(resp)
	} else {
		return models.CurrencyModel{}
	}
}
