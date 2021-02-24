package actions

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func connect() *sql.DB {
	connstr := os.Getenv("DATABASE_USER") + ":" + os.Getenv("DATABASE_PASSWORD") + "@tcp(" + os.Getenv("DATABASE_HOST") + ")/" + os.Getenv("DATABASE_NAME")
	conn, err := sql.Open("mysql", connstr)
	if err != nil {
		panic(err.Error())
	}
	return conn
}
