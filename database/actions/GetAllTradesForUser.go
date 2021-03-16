package actions

import (
	"github.com/MathisBurger/crypto-simulator/database/models"
	"time"
)

func GetAllTradesForUser(UUID string) []models.TradesModel {

	conn, _ := Connect()
	defer conn.Close()

	stmt, _ := conn.Prepare("SELECT * FROM `trades` WHERE `walletUUID`=? AND `timestamp`>?")
	defer stmt.Close()

	resp, _ := stmt.Query(UUID, time.Now().Unix()-604800)
	defer resp.Close()

	return models.TradesModel{}.ParseAll(resp)
}
