package actions

func InsertCurrencyChange(currency string, priceUSD float64, timestamp int) {
	conn, _ := connect()
	defer conn.Close()
	stmt, _ := conn.Prepare("INSERT INTO `currency-progress` (`ID`, `currency`, `priceUSD`, `timestamp`) VALUES (NULL, ?, ?, ?);")
	defer stmt.Close()
	stmt.Exec(currency, priceUSD, timestamp)
}
