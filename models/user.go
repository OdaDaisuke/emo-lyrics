package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	HandleName string `gorm:"type varchar(40)"`
	TwitterID string `gorm:"type varchar(20)"`
}

func (e User) TableName() string {
	return "users"
}
