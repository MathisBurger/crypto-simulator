package actions

func LoginWithToken(username string, token string) bool {
	conn, _ := connect()
	defer conn.Close()
	stmt, err := conn.Prepare("SELECT * FROM `user` WHERE `username`=? AND `AuthToken`=?;")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	resp, err := stmt.Query(username, token)
	if err != nil {
		panic(err.Error())
	}
	defer resp.Close()
	return resp.Next()
}
