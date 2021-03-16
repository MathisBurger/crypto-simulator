package actions

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func connect() (*sql.DB, error) {

	connstr := os.Getenv("DATABASE_USER") + ":" + os.Getenv("DATABASE_PASSWORD") + "@tcp(" + os.Getenv("DATABASE_HOST") + ")/" + os.Getenv("DATABASE_NAME") + "?parseTime=true"

	conn, err := sql.Open("mysql", connstr)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return conn, nil
}
