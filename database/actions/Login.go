package actions

import (
	"github.com/MathisBurger/crypto-simulator/database/models"
	"github.com/MathisBurger/crypto-simulator/utils"
	"strings"
)

func Login(username string, password string) bool {
	conn, _ := connect()
	defer conn.Close()
	stmt, _ := conn.Prepare("SELECT * FROM `user` WHERE `username`=?")
	defer stmt.Close()
	resp, _ := stmt.Query(username)
	defer resp.Close()
	var model models.UserModel
	if !resp.Next() {
		return false
	}
	model = model.ParseModel(resp)
	spl := strings.Split(model.Password, "§")
	if utils.DoPasswordsMatch(model.Password, password, spl[0]) {
		return true
	}
	return false
}
