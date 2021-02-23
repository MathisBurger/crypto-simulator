package actions

import "github.com/MathisBurger/crypto-simulator/database/models"

func AddMoneyToWallet(UUID string, amount float64) {
	conn, _ := connect()
	defer conn.Close()
	stmt, _ := conn.Prepare("SELECT * FROM `wallets` WHERE `UUID`=?")
	defer stmt.Close()
	resp, _ := stmt.Query(UUID)
	defer resp.Close()
	resp.Next()
	mdl := models.WalletModel{}.Parse(resp)
	newValue := mdl.BalanceUSD + amount
	stmt, _ = conn.Prepare("UPDATE `wallets` SET `balanceUSD`=? WHERE `UUID`=?")
	defer stmt.Close()
	stmt.Exec(newValue, UUID)
}
