package actions

func InitTables() {
	conn, _ := connect()
	stmt, err := conn.Prepare("CREATE TABLE `user` ( `ID` INT NOT NULL AUTO_INCREMENT , `username` VARCHAR(64) NOT NULL , `password` TEXT NOT NULL , `walletUUID` VARCHAR(32)NOT NULL , `AuthToken` VARCHAR(128) NOT NULL , `created-at` BIGINT NOT NULL , PRIMARY KEY (`ID`));")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec()
	if err != nil {
	}
	stmt, err = conn.Prepare("CREATE TABLE `wallets` ( `id` INT NOT NULL AUTO_INCREMENT , `UUID` VARCHAR(32) NOT NULL , `balanceUSD` FLOAT NOT NULL , `currencyArray` TEXT NOT NULL , PRIMARY KEY (`id`));")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec()
	if err != nil {
	}
	stmt, err = conn.Prepare("CREATE TABLE `trades` ( `id` INT NOT NULL AUTO_INCREMENT , `trade` VARCHAR(8) NOT NULL , `walletUUID` INT NOT NULL , `CoinPrice` FLOAT NOT NULL , `OfferPrice` INT NOT NULL , `timestamp` BIGINT NOT NULL , PRIMARY KEY (`id`));")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec()
	if err != nil {
	}
	stmt, err = conn.Prepare("CREATE TABLE `currencys` ( `ID` INT NOT NULL AUTO_INCREMENT , `name` VARCHAR(3) NOT NULL , `price` FLOAT NOT NULL , `last-updated` BIGINT NOT NULL , PRIMARY KEY (`ID`));")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec()
	if err != nil {
	}
	stmt, err = conn.Prepare("CREATE TABLE `currency-progress` ( `ID` INT NOT NULL AUTO_INCREMENT , `currency` VARCHAR(3) NOT NULL , `priceUSD` FLOAT NOT NULL , `timestamp` BIGINT NOT NULL , PRIMARY KEY (`ID`));")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec()
	if err != nil {
	}
	disconnect(conn)
}