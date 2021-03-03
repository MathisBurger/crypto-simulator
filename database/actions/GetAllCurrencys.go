package actions

import "github.com/MathisBurger/crypto-simulator/database/models"

func GetAllCurrencys() []models.CurrencyModel {

	conn, _ := connect()
	defer conn.Close()

	stmt, _ := conn.Prepare("SELECT * FROM `currencys`;")
	defer stmt.Close()

	resp, _ := stmt.Query()
	defer resp.Close()

	return models.CurrencyModel{}.ParseAll(resp)
}
