package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
)

func NewDBContext() *gorm.DB {

	dbargs := "charset=utf8mb4"
	timezone := "parseTime=true&loc=Asia%2FTokyo"
	connectionStr := fmt.Sprintf("%s:%s@tcp([%s]:%s)/%s?%s&%s&%s",
		//os.Getenv("MYSQL_USER"),
		//os.Getenv("MYSQL_PASSWORD"),
		//os.Getenv("MYSQL_HOST"),
		//os.Getenv("MYSQL_PORT"),
		//os.Getenv("MYSQL_DATABASE"),
		"root",
		"password",
		"127.0.0.1",
		"3306",
		"emo",
		"allowNativePasswords=true",
		dbargs,
		timezone,
	)

	db, err := gorm.Open("mysql", connectionStr)
	if err != nil {
		panic(err.Error())
	}

	db = db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4")


	return db
}
