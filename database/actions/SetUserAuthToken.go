package actions

func SetUserAuthToken(username string, token string) {

	conn, _ := Connect()
	defer conn.Close()

	stmt, _ := conn.Prepare("UPDATE `user` SET `AuthToken`=? WHERE `username`=?")
	defer stmt.Close()

	stmt.Exec(token, username)
}
