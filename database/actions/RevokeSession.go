package actions

func RevokeSession(user string, token string) {
	conn, _ := connect()
	defer conn.Close()
	stmt, _ := conn.Prepare("DELETE FROM `refresh-token` WHERE `username`=? AND `token`=?")
	defer stmt.Close()
	stmt.Exec(user, token)
}
