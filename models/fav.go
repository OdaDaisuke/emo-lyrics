package models

import (
	"github.com/jinzhu/gorm"
)

type Fav struct {
	gorm.Model
	UserID uint `json:user_id`
	LyricID uint `json:lyric_id`
}

func (e Fav) TableName() string {
	return "favs"
}
