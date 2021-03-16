package actions

import "time"

func AddTrade(from string, to string, walletUUID string, price float64, amount float64) {

	conn, _ := Connect()
	defer conn.Close()

	stmt, _ := conn.Prepare("INSERT INTO `trades` (`id`, `trade`, `walletUUID`, `CoinPrice`, `Amount`, `timestamp`) VALUES (NULL, ?, ?, ?, ?, ?);")
	defer stmt.Close()

	stmt.Exec(from+"->"+to, walletUUID, price, amount, time.Now().Unix())
}
