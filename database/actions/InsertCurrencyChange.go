package actions

// -----------------------------------------
//              DEPRECATED
//    This function is handled anyway.
//    It is deprecated since v0.0.1-dev
// -----------------------------------------
func InsertCurrencyChange(currency string, priceUSD float64, timestamp int) {

	conn, _ := connect()
	defer conn.Close()

	stmt, _ := conn.Prepare("INSERT INTO `currency-progress` (`ID`, `currency`, `priceUSD`, `timestamp`) VALUES (NULL, ?, ?, ?);")
	defer stmt.Close()

	stmt.Exec(currency, priceUSD, timestamp)
}
