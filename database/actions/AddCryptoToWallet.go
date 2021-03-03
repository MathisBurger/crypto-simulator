package actions

import (
	"fmt"
	"github.com/MathisBurger/crypto-simulator/database/models"
	"strconv"
	"strings"
)

func AddCryptoToWallet(UUID string, symbol string, amount float64) {

	conn, _ := connect()
	defer conn.Close()

	stmt, _ := conn.Prepare("SELECT * FROM `wallets` WHERE `UUID`=?")
	defer stmt.Close()

	resp, _ := stmt.Query(UUID)
	defer resp.Close()

	resp.Next()

	// get array of user currencies
	currencyArray := strings.Split(models.WalletModel{}.Parse(resp).CurrencyArray, ";")
	var builder strings.Builder

	// get currency from array and increase it
	if len(currencyArray) != 0 {
		found := false
		for i, el := range currencyArray {
			spl := strings.Split(el, "|")
			if spl[0] == symbol {
				fl, _ := strconv.ParseFloat(spl[1], 64)
				currencyArray[i] = spl[0] + "|" + fmt.Sprintf("%f", fl+amount)
				found = true
			}
		}

		// if array does not contain add it
		if !found {
			currencyArray = append(currencyArray, symbol+"|"+fmt.Sprintf("%f", amount))
		}

		// build currencyArray to string
		for i, el := range currencyArray {
			if i == 0 {
				builder.WriteString(el)
				continue
			}
			builder.WriteString(";" + el)
		}

	} else {
		builder.WriteString(symbol + "|" + fmt.Sprintf("%f", amount))
	}

	stmt, _ = conn.Prepare("UPDATE `wallets` SET `currencyArray`=? WHERE `UUID`=?")
	defer stmt.Close()

	stmt.Exec(builder.String(), UUID)
}
