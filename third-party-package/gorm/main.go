package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func gormConnect() *gorm.DB {
	DBMS     := "mysql"
	USER     := "root"
	PASS     := "root"
	PROTOCOL := "tcp(0.0.0.0:3306)"
	DBNAME   := "test"

	CONNECT := USER+":"+PASS+"@"+PROTOCOL+"/"+DBNAME
	db,err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err)
	}
	return db
}

type event struct {
	Id          int    `json:id sql:AUTO_INCREMENT`
	User_Id     int    `json:user_id`
	Summary     string `json:summary`
	Dtstart     string `json:dtstart`
	Dtend       string `json:dtend`
	Description string `json:description`
	Year        int    `json:year`
	Month       int    `json:month`
	Day         int    `json:day`
}


// see http://doc.gorm.io/crud.html
func main(){
	db := gormConnect()

	db.CreateTable(&event{})
	db.Begin()
}