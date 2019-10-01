package repository

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"gopkg.in/gorp.v1"
	"gopkg.in/testfixtures.v2"
	"log"
	"os"
	"testing"
)

var (
	fixtures        *testfixtures.Context
	usersRepository *usersRepo
	dbMap           *gorp.DbMap
)

func TestMain(m *testing.M) {
	// in memory db setup
	db, err := sql.Open("sqlite3", "file:fixture_test.db?mode=memory")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if _, err := db.Exec("CREATE TABLE users (id int PRIMARY KEY,name varchar(255))"); err != nil {
		log.Fatal(err)
	}

	// fixture setup
	fixtures, err = testfixtures.NewFiles(db, &testfixtures.SQLite{}, "fixtures/users.yml")
	if err != nil {
		log.Fatal(err)
	}
	testfixtures.SkipDatabaseNameCheck(true)

	// repository setup
	usersRepository = NewUsersRepo()
	dbMap = &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}

	os.Exit(m.Run())
}

func TestSuccessSelectOneById(t *testing.T) {
	// Fixture
	if err := fixtures.Load(); err != nil {
		assert.Fail(t, err.Error())
	}

	// Exec
	expected := User{ID: 1, Name: "Ken"}
	user, err := usersRepository.SelectOneById(*dbMap, expected.ID)
	if err != nil {
		assert.Fail(t, err.Error())
	}

	// Assert
	assert.Equal(t, expected.ID, user.ID)
	assert.Equal(t, expected.Name, user.Name)
}

func TestSuccessInsert(t *testing.T) {
	// Exec
	expected := User{ID: 2, Name: "Jon"}
	if err := usersRepository.Insert(*dbMap, expected); err != nil {
		assert.Fail(t, err.Error())
	}

	// Assert
	var user User
	if err := dbMap.SelectOne(&user, "SELECT id,name FROM users WHERE id=?", expected.ID); err != nil {
		assert.Fail(t, err.Error())
	}
	assert.Equal(t, expected.ID, user.ID)
	assert.Equal(t, expected.Name, user.Name)
}

func TestSuccessUpdate(t *testing.T) {
	// Fixture
	if err := fixtures.Load(); err != nil {
		assert.Fail(t, err.Error())
	}

	// Exec
	expected := User{ID: 3, Name: "Ren"}
	if err := usersRepository.Update(*dbMap, expected); err != nil {
		assert.Fail(t, err.Error())
	}

	// Assert
	var user User
	if err := dbMap.SelectOne(&user, "SELECT id,name FROM users WHERE id=?", expected.ID); err != nil {
		assert.Fail(t, err.Error())
	}
	assert.Equal(t, expected.ID, user.ID)
	assert.Equal(t, expected.Name, user.Name)
}
