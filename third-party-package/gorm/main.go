package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

func gormConnect() *gorm.DB {
	DBMS := "mysql"
	USER := "root"
	PASS := "root"
	PROTOCOL := "tcp(0.0.0.0:3306)"
	DBNAME := "test"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err)
	}
	return db
}

type event struct {
	Id        int    `json:id sql:AUTO_INCREMENT`
	Name      string `json:name`
	DeletedAt *time.Time
}

// see http://doc.gorm.io/crud.html
func main() {
	db := gormConnect()
	db.AutoMigrate(&event{})
	err := CreateEvents(db)
	if err != nil {
		fmt.Printf("%#v", err)
	}

	if err := DeleteEvent(db); err != nil {
		fmt.Printf("%#v", err)
	}
}

func CreateEvents(db *gorm.DB) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Create(&event{Name: "A"}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Create(&event{Name: "B"}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func DeleteEvent(db *gorm.DB) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	where := event{
		Id: 12,
	}
	if err := tx.Where(where).Delete(&event{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
