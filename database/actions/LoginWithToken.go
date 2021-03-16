package actions

func LoginWithToken(username string, token string) bool {

	conn, _ := Connect()
	defer conn.Close()

	stmt, _ := conn.Prepare("SELECT * FROM `user` WHERE `username`=? AND `AuthToken`=?;")
	defer stmt.Close()

	resp, _ := stmt.Query(username, token)
	defer resp.Close()

	return resp.Next()
}
