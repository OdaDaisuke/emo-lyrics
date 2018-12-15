package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
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
		"KfasdhkLe",
		"sample.c2crqlialh8o.ap-northeast-1.rds.amazonaws.com",
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

	db.LogMode(true)

	return db
}
