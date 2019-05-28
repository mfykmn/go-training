package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	*sql.DB
}

func newDB() (*DB, error) {
	var (
		connectionName = os.Getenv("CLOUDSQL_CONNECTION_NAME")
		user           = os.Getenv("CLOUDSQL_USER")
		password       = os.Getenv("CLOUDSQL_PASSWORD")
	)
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@unix(/cloudsql/%s)/", user, password, connectionName))
	if err != nil {
		return nil, err
	}
	return &DB{db}, nil
}

func (db *DB) Show() ([]byte, error) {
	rows, err := db.Query("SHOW DATABASES")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	buf := bytes.NewBufferString("Databases:\n")
	for rows.Next() {
		var dbName string
		if err := rows.Scan(&dbName); err != nil {
			return nil, err
		}
		fmt.Fprintf(buf, "- %s\n", dbName)
	}
	return buf.Bytes(), nil
}
