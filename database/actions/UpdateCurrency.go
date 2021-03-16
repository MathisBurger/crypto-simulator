package actions

func UpdateCurrency(
	rank int, supply float64, maxSupply float64, marketCapUSD float64, VolumeUSD24Hr float64, priceUSD float64,
	changePercent24Hr float64, vwap24Hr float64, symbol string) {

	conn, _ := Connect()
	defer conn.Close()

	stmt, _ := conn.Prepare("UPDATE `currencys` SET `rank` = ?, `supply` = ?, `maxSupply` = ?, `marketCapUSD` = ?, `volumeUSD24Hr` = ?, `priceUSD` = ?, `changePercent24Hr` = ?, `vwap24Hr` = ? WHERE `symbol`=?;")
	defer stmt.Close()

	stmt.Exec(rank, supply, maxSupply, marketCapUSD, VolumeUSD24Hr, priceUSD, changePercent24Hr, vwap24Hr, symbol)
}
