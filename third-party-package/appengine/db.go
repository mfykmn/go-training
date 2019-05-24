package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func newDB() (*sql.DB, error) {
	var (
		connectionName = os.Getenv("CLOUDSQL_CONNECTION_NAME")
		user           = os.Getenv("CLOUDSQL_USER")
		password       = os.Getenv("CLOUDSQL_PASSWORD")
	)
	return sql.Open("mysql", fmt.Sprintf("%s:%s@unix(/cloudsql/%s)/", user, password, connectionName))
}
