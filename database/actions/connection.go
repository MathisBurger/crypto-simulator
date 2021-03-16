package actions

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func connect() (*sql.DB, error) {

	connstr := os.Getenv("DATABASE_USER") + ":" + os.Getenv("DATABASE_PASSWORD") + "@tcp(" + os.Getenv("DATABASE_HOST") + ")/" + os.Getenv("DATABASE_NAME") + "?parseTime=true"

	conn, err := sql.Open("mysql", connstr)

	if err != nil {
		return nil, err
	}

	return conn, nil
}

// DEPRECATED
func disconnect(conn *sql.DB) {
	defer conn.Close()
}
