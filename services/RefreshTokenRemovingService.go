package services

import (
	"github.com/MathisBurger/crypto-simulator/database/actions"
	"time"
)

// This services deletes all expired refresh token
// It checks the database every minute for those tokens
func RefreshTokenRemovingService() {

	for _ = range time.Tick(time.Minute) {

		conn, _ := actions.Connect()
		defer conn.Close()
		stmt, _ := conn.Prepare("DELETE FROM `refresh-token` WHERE `Deadline`<?")
		defer stmt.Close()
		stmt.Exec(time.Now())

	}
}
