package actions

func InsertCurrency(name string, price float64, lastUpdated int) {
	conn, _ := connect()
	defer conn.Close()
	stmt, _ := conn.Prepare("INSERT INTO `currencys` (`ID`, `name`, `price`, `last-updated`) VALUES (NULL, ?, ?, ?);")
	defer stmt.Close()
	stmt.Exec(name, price, lastUpdated)
}
