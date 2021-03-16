package actions

import (
	"database/sql"
	"github.com/MathisBurger/crypto-simulator/utils"
	"time"
)

func RegisterAccount(username string, password string) bool {
	conn, _ := Connect()
	defer conn.Close()

	walletStatus, walletUUID := createWallet(conn)

	if !walletStatus {
		return false
	}

	stmt, _ := conn.Prepare("SELECT * FROM `user` WHERE `username`=?;")
	defer stmt.Close()

	resp, _ := stmt.Query(username)
	defer resp.Close()

	if resp.Next() {
		return false
	}

	stmt, _ = conn.Prepare("INSERT INTO `user` (`ID`, `username`, `password`, `walletUUID`, `AuthToken`, `created-at`) VALUES (NULL, ?, ?, ?, 'None', ?);")

	_, err := stmt.Exec(username, utils.HashPassword(password), walletUUID, time.Now().Unix())
	if err != nil {
		return false
	} else {
		return true
	}
}

func createWallet(conn *sql.DB) (bool, string) {

	stmt, _ := conn.Prepare("INSERT INTO `wallets` (`id`, `UUID`, `balanceUSD`, `currencyArray`) VALUES (NULL, ?, '100', '');")
	defer stmt.Close()

	uuid := utils.GenerateUUID()

	_, err := stmt.Exec(uuid)
	if err != nil {
		return false, ""
	}

	return true, uuid
}
