package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	TwitterId            string `gorm:"type varchar(20)"`
	Lang                 string `gorm:"type varchar(4)"`
	Location             string `gorm:"type varchar(15)"`
	Name                 string `gorm:"type varchar(20)"`
	ProfileBannerUrl     string `gorm:"type varchar(150)"`
	ProfileImageUrlHttps string `gorm:"type varchar(150)"`
	Protected            string `gorm:"type varchar(10)"`
	ScreenName           string `gorm:"type varchar(20)"`
	Url                  string `gorm:"type varchar(150)"`
}

func (e User) TableName() string {
	return "users"
}
