package testutils

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

var dbConn *sql.DB

const schemaDirRelativePathFormat = "%s/../../schema/%s"
const fixturesDirRelativePathFormat = "%s/../../schema/fixtures/%s"

func SetupMySQLConn() func() {
	c := mysql.Config{
		DBName:               os.Getenv("TEST_MYSQL_DATABASE"),
		User:                 os.Getenv("TEST_MYSQL_USER"),
		Passwd:               os.Getenv("TEST_MYSQL_PASSWORD"),
		Addr:                 os.Getenv("TEST_MYSQL_ADDRESS"),
		Net:                  "tcp",
		Loc:                  time.UTC,
		ParseTime:            true,
		AllowNativePasswords: true,
	}
	db, err := sql.Open("mysql", c.FormatDSN())
	if err != nil {
		log.Fatalf("Could not connect to mysql: %s", err)
	}
	db.SetMaxIdleConns(0)
	dbConn = db

	return func() { dbConn.Close() }
}

// GetTestMySQLConn ... プールしてあるテスト用のDBコネクションを返す
func GetTestMySQLConn() (*sql.DB, func()) {
	if dbConn == nil {
		panic("mysql connection is not initialized yet")
	}
	return dbConn, func() { truncateTables() }
}

// setupDefaultFixtures ... 全テストに共通するFixtureのInsert
func setupDefaultFixtures() {
	_, pwd, _, _ := runtime.Caller(0)

	defaultFixtureDir := fmt.Sprintf(fixturesDirRelativePathFormat, path.Dir(pwd), "default")
	defaultFixturePathes := walkSchema(defaultFixtureDir)
	for _, dpath := range defaultFixturePathes {
		execSchema(dpath)
	}
}

// SetupOptionalFixtures ... テストケースごとに任意に設定するFixtureのInsert
func SetupOptionalFixtures(names []string) {
	_, pwd, _, _ := runtime.Caller(0)

	optionalFixtureDir := fmt.Sprintf(fixturesDirRelativePathFormat, path.Dir(pwd), "optional")
	for _, n := range names {
		opath := filepath.Join(optionalFixtureDir, fmt.Sprintf("%s.sql", n))
		execSchema(opath)
	}
}

func walkSchema(dir string) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	var paths []string
	for _, file := range files {
		paths = append(paths, filepath.Join(dir, file.Name()))
	}

	return paths
}

func execSchema(fpath string) {
	b, err := ioutil.ReadFile(fpath)
	if err != nil {
		log.Fatalf("schema reading error: %v", err)
	}

	queries := strings.Split(string(b), ";")

	for _, query := range queries[:len(queries)-1] {
		query := query
		_, err = dbConn.Exec(query)
		if err != nil {
			log.Printf("exec schema error: %v, query: %s", err, query)
			continue
		}
	}
}

func createTablesIfNotExist() {
	_, pwd, _, _ := runtime.Caller(0)
	schemaPath := fmt.Sprintf(schemaDirRelativePathFormat, path.Dir(pwd), "schema.sql")
	execSchema(schemaPath)
}

func truncateTables() {
	rows, err := dbConn.Query("SHOW TABLES")
	if err != nil {
		log.Fatalf("show tables error: %#v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var tableName string
		err = rows.Scan(&tableName)
		if err != nil {
			log.Fatalf("show table error: %#v", err)
			continue
		}

		cmds := []string{
			"SET FOREIGN_KEY_CHECKS = 0",
			fmt.Sprintf("TRUNCATE %s", tableName),
			"SET FOREIGN_KEY_CHECKS = a",
		}
		for _, cmd := range cmds {
			if _, err := dbConn.Exec(cmd); err != nil {
				log.Fatalf("truncate error: %#v", err)
				continue
			}
		}
	}
}