package actions

import "github.com/MathisBurger/crypto-simulator/database/models"

func GetUserByUsername(username string) models.UserModel {

	conn, _ := connect()
	defer conn.Close()

	stmt, _ := conn.Prepare("SELECT * FROM `user` WHERE `username`=?")
	defer stmt.Close()

	resp, _ := stmt.Query(username)
	defer resp.Close()

	// must exist
	resp.Next()

	return models.UserModel{}.ParseModel(resp)
}
