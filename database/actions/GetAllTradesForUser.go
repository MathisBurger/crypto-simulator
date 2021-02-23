package actions

import (
	"github.com/MathisBurger/crypto-simulator/database/models"
	"time"
)

func GetAllTradesForUser(UUID string) []models.TradesModel {
	conn, _ := connect()
	defer conn.Close()
	stmt, err := conn.Prepare("SELECT * FROM `trades` WHERE `walletUUID`=? AND `timestamp`>?")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	resp, err := stmt.Query(UUID, time.Now().Unix()-604800)
	if err != nil {
		panic(err.Error())
	}
	defer resp.Close()
	return models.TradesModel{}.ParseAll(resp)
}
