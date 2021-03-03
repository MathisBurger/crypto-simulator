package actions

func InsertCurrency(
	coinid string, rank int, symbol string, name string, supply float64, maxSupply float64, marketCapUSD float64,
	volumeUSD24Hr float64, priceUSD float64, changePercent24Hr float64, vwap24Hr float64) {

	conn, _ := connect()
	defer conn.Close()

	stmt, _ := conn.Prepare("INSERT INTO `currencys` (`id`, `CoinID`, `rank`, `symbol`, `name`, `supply`, `maxSupply`, `marketCapUSD`, `volumeUSD24Hr`, `priceUSD`, `changePercent24Hr`, `vwap24Hr`) VALUES (NULL, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);")
	defer stmt.Close()

	stmt.Exec(coinid, rank, symbol, name, supply, maxSupply, marketCapUSD, volumeUSD24Hr, priceUSD, changePercent24Hr, vwap24Hr)
}
