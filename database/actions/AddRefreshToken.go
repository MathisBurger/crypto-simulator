package actions

import "github.com/MathisBurger/crypto-simulator/database/models"

func AddRefreshToken(tkn *models.RefreshTokenModel) {
	conn, _ := connect()
	defer conn.Close()
	stmt, _ := conn.Prepare("INSERT INTO `refresh-token` (`ID`, `username`, `token`, `Deadline`) VALUES (NULL, ?, ?, ?);")
	defer stmt.Close()
	stmt.Exec(tkn.Username, tkn.Token, tkn.Deadline)
}
