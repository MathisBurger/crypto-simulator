package actions

import "github.com/MathisBurger/crypto-simulator/database/models"

func GetWalletByUUID(UUID string) models.WalletModel {

	conn, _ := connect()
	defer conn.Close()

	stmt, _ := conn.Prepare("SELECT * FROM `wallets` WHERE `UUID`=?")
	defer stmt.Close()

	resp, _ := stmt.Query(UUID)
	defer resp.Close()

	// must exist
	resp.Next()

	return models.WalletModel{}.Parse(resp)
}
