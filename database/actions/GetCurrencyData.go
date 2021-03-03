package actions

import (
	"github.com/MathisBurger/crypto-simulator/database/models"
	"time"
)

// ---------------------------------
//            DEPRECATED
//    This action is deprecated
//   It can be enabled via config
// ---------------------------------
func GetCurrencyData(name string, timePeriod int) []models.CurrencyProgressModel {

	timestamp := time.Now().Unix() - int64(timePeriod)

	conn, _ := connect()
	defer conn.Close()

	stmt, _ := conn.Prepare("SELECT * FROM `currency-progress` WHERE `currency`=? AND `timestamp`>=?")
	defer stmt.Close()

	resp, _ := stmt.Query(name, timestamp)
	defer resp.Close()

	return models.CurrencyProgressModel{}.ParseAll(resp)
}
