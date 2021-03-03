package actions

import "github.com/MathisBurger/crypto-simulator/database/models"

func RemoveMoneyFromWallet(UUID string, value float64) {

	conn, _ := connect()
	defer conn.Close()

	stmt, _ := conn.Prepare("SELECT * FROM `wallets` WHERE `UUID`=?")
	defer stmt.Close()

	resp, _ := stmt.Query(UUID)
	defer resp.Close()

	// must exist
	resp.Next()

	mdl := models.WalletModel{}.Parse(resp)
	newValue := mdl.BalanceUSD - value

	stmt, _ = conn.Prepare("UPDATE `wallets` SET `balanceUSD`=? WHERE `UUID`=?")
	defer stmt.Close()

	stmt.Exec(newValue, UUID)
}
