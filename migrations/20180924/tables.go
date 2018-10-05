package v20180924

import "github.com/jinzhu/gorm"

// Lyric
type Lyric struct {
	gorm.Model
	Lyric  string `json:lyric`
	Title  string `json:title`
	Singer string `json:singer`
	Url    string `json:url`
	Genre  string `json:genre`
}

func (e Lyric) TableName() string {
	return "lyrics"
}

// Fav
type Fav struct {
	gorm.Model
	UserID  uint `json:user_id`
	LyricID uint `json:lyric_id`
}

func (e Fav) TableName() string {
	return "favs"
}

// User
type User struct {
	gorm.Model
	ProviderId           string `gorm:"type varchar(20)"`
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
