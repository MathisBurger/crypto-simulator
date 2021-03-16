package actions

import "github.com/MathisBurger/crypto-simulator/database/models"

func GetRefreshToken(token string) (bool, models.RefreshTokenModel) {
	conn, _ := Connect()
	defer conn.Close()
	stmt, _ := conn.Prepare("SELECT * FROM `refresh-token` WHERE `token`=?")
	defer stmt.Close()
	resp, _ := stmt.Query(token)
	defer resp.Close()
	return models.RefreshTokenModel{}.Parse(resp)
}
