package database

import (
	"github.com/MathisBurger/crypto-simulator/database/actions"
	_ "github.com/go-sql-driver/mysql"
)

func CreateRequiredTables() {
	actions.InitTables()
}

func CreateAccount(username string, password string) bool {
	return actions.RegisterAccount(username, password)
}
