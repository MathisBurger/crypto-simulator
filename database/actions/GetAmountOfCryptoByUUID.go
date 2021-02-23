package actions

import (
	"github.com/MathisBurger/crypto-simulator/database/models"
	"strconv"
	"strings"
)

func GetAmountOfCryptoByUUID(UUID string, symbol string) float64 {
	conn, _ := connect()
	defer conn.Close()
	stmt, _ := conn.Prepare("SELECT * FROM `wallets` WHERE `UUID`=?")
	defer stmt.Close()
	resp, _ := stmt.Query(UUID)
	defer resp.Close()
	resp.Next()
	currencyArray := strings.Split(models.WalletModel{}.Parse(resp).CurrencyArray, ";")
	if len(currencyArray) > 1 {
		for _, el := range currencyArray {
			spl := strings.Split(el, "|")
			if spl[0] == symbol {
				fl, _ := strconv.ParseFloat(spl[1], 64)
				return fl
			}
		}
		return 0
	} else {
		return 0.0
	}
}
