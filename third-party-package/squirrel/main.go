package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

type UserID int

type User struct {
	ID   UserID
	Name string
}

func main() {
	conf := &mysql.Config{
		User:                 "user",
		Passwd:               "pass",
		Addr:                 "localhost:8080",
		Net:                  "tcp",
		DBName:               "demo",
		ParseTime:            true,
		AllowNativePasswords: true,
	}

	db, err := sql.Open("mysql", conf.FormatDSN())
	if err != nil {
		panic(err)
	}

	repo := &UserRepository{db}

	user := &User{
		ID:   UserID(1),
		Name: "taro",
	}
	ctx := context.Background()
	repo.WithTransaction(ctx, func(tx *sql.Tx) error {
		err := repo.CreateUser(ctx, user)
		if err != nil {
			return err
		}
		return nil
	})

	fmt.Println(repo.GetUser(ctx, user.ID))
}
