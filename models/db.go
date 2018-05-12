package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func NewDBContext(db *gorm.DB) *gorm.DB {
	return db
}

func GormConnect() *gorm.DB {
	DBMS := "mysql"
	USER := "root"
	PASS := "daisukeoda"
	PROTOCOL := "tcp(127.0.0.1:3306)"
	DBNAME := "emo_lyric"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	return db
}
