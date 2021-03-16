package actions

func CheckUserOwnsRefreshToken(user string, token string) bool {
	conn, _ := connect()
	defer conn.Close()
	stmt, _ := conn.Prepare("SELECT * FROM `refresh-token` WHERE `username`=? AND `token`=?")
	defer stmt.Close()
	resp, _ := stmt.Query(user, token)
	defer resp.Close()
	return resp.Next()
}
