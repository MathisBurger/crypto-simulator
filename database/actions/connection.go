package actions

import (
	"database/sql"
	"os"
)

func connect() (*sql.DB, error) {
	connstr := os.Getenv("DATABASE_USER") + ":" + os.Getenv("DATABASE_PASSWORD") + "@tcp(" + os.Getenv("DATABASE_HOST") + ")/" + os.Getenv("DATABASE_NAME")
	conn, err := sql.Open("mysql", connstr)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func disconnect(conn *sql.DB) {
	defer conn.Close()
}
