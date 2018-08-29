package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
	sq "gopkg.in/Masterminds/squirrel.v1"
)

const (
	// Table
	userTable = "users"
)

type UserRepository struct {
	*sql.DB
}

func (r *UserRepository) WithTransaction(ctx context.Context, txFunc func(*sql.Tx) error) (err error) {
	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // re-throw panic after Rollback
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	err = txFunc(tx)
	return err
}

func (r *UserRepository) GetUser(ctx context.Context, userID UserID) (*User, error) {
	user := &User{
		ID: userID,
	}

	err := sq.Select("name").
		From("users").
		Where(sq.Eq{
			"id": userID,
		}).
		RunWith(r.DB).
		QueryRow().
		Scan(&user.Name)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, nil
		default:
			return nil, errors.Wrap(err, fmt.Sprintf("failed mysql GetUser query."))
		}
	}
	return user, nil
}

func (r *UserRepository) CreateUser(ctx context.Context, user *User) error {
	_, err := sq.Insert(userTable).
		Columns("id", "name").
		Values(user.ID, user.Name).
		RunWith(r.DB).Exec()
	return err
}
