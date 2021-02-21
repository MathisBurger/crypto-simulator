package actions

func CheckIfCurrencyExists(name string) bool {
	conn, _ := connect()
	defer conn.Close()
	stmt, _ := conn.Prepare("SELECT * FROM `currencys` WHERE `name`=?")
	defer stmt.Close()
	resp, err := stmt.Query(name)
	if err != nil {
		panic(err.Error())
	}
	defer resp.Close()
	return resp.Next()
}
