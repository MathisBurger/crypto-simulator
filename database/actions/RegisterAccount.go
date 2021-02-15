package actions

import (
	"database/sql"
	"fmt"
	"github.com/MathisBurger/crypto-simulator/utils"
	"time"
)

func RegisterAccount(username string, password string) bool {
	conn, err := connect()
	defer conn.Close()
	if err != nil {
		return false
	}
	walletStatus, walletUUID := createWallet(conn)
	if !walletStatus {
		return false
	}
	stmt, err := conn.Prepare("SELECT * FROM `user` WHERE `username`=?;")
	if err != nil {
		return false
	}
	defer stmt.Close()
	resp, _ := stmt.Query(username)
	defer resp.Close()
	if resp.Next() {
		return false
	}
	stmt, err = conn.Prepare("INSERT INTO `user` (`ID`, `username`, `password`, `walletUUID`, `AuthToken`, `created-at`) VALUES (NULL, ?, ?, ?, 'None', ?);")
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	_, err = stmt.Exec(username, utils.HashPassword(password), walletUUID, time.Now().Unix())
	if err != nil {
		return false
	} else {
		return true
	}
}

func createWallet(conn *sql.DB) (bool, string) {
	stmt, err := conn.Prepare("INSERT INTO `wallets` (`id`, `UUID`, `balanceUSD`, `currencyArray`) VALUES (NULL, ?, '100', '[]');")
	if err != nil {
		return false, ""
	}
	uuid := utils.GenerateUUID()
	_, err = stmt.Exec(uuid)
	if err != nil {
		return false, ""
	}
	return true, uuid
}
