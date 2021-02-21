package actions

func UpdateCurrency(name string, price float64, lastUpdated int) {
	conn, _ := connect()
	defer conn.Close()
	stmt, _ := conn.Prepare("UPDATE `currencys` SET `price`=?, `last-updated`=? WHERE `name`=?")
	defer stmt.Close()
	stmt.Exec(price, lastUpdated, name)
}
