package repository

import (
	"gopkg.in/gorp.v1"
)

type UsersRepo interface {
	SelectOneById(dbMap gorp.DbMap, id int) (*User, error)
	Insert(dbMap gorp.DbMap, user *User) error
	Update(dbMap gorp.DbMap, user User) error
}

type User struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

type usersRepo struct {
}

func NewUsersRepo() *usersRepo {
	return new(usersRepo)
}

func (*usersRepo) SelectOneById(dbMap gorp.DbMap, id int) (*User, error) {
	var user User
	if err := dbMap.SelectOne(&user, "SELECT id,name FROM users WHERE id=?", id); err != nil {
		return nil, err
	}
	return &user, nil
}

func (*usersRepo) Insert(dbMap gorp.DbMap, user User) error {
	_, err := dbMap.Exec("INSERT INTO users (id,name) VALUES (?,?)", user.ID, user.Name)
	if err != nil {
		return err
	}
	return nil
}

func (*usersRepo) Update(dbMap gorp.DbMap, user User) error {
	_, err := dbMap.Exec("UPDATE users SET name=? WHERE id = ?", user.Name, user.ID)
	if err != nil {
		return err
	}
	return nil
}
