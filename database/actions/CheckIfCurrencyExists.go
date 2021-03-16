package actions

func CheckIfCurrencyExists(name string) bool {

	conn, _ := Connect()
	defer conn.Close()

	stmt, _ := conn.Prepare("SELECT * FROM `currencys` WHERE `symbol`=?")
	defer stmt.Close()

	resp, _ := stmt.Query(name)
	defer resp.Close()

	// exists if true
	return resp.Next()
}
