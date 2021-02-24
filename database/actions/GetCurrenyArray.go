package actions

import (
	"github.com/MathisBurger/crypto-simulator/database/models"
	"strconv"
	"strings"
)

func GetCurrencyArray(UUID string) []models.CurrencyArrayModel {
	conn, _ := connect()
	defer conn.Close()
	stmt, _ := conn.Prepare("SELECT * FROM `wallets` WHERE `UUID`=?")
	defer stmt.Close()
	resp, _ := stmt.Query(UUID)
	defer resp.Close()
	resp.Next()
	arr := strings.Split(models.WalletModel{}.Parse(resp).CurrencyArray, ";")
	var answers []models.CurrencyArrayModel
	for _, el := range arr {
		spl := strings.Split(el, "|")
		if len(spl) == 2 {
			fl, _ := strconv.ParseFloat(spl[1], 64)
			if fl != 0 {
				answers = append(answers, models.CurrencyArrayModel{spl[0], fl})
			}
		}
	}
	return answers
}
