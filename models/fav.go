package models

import (
	"github.com/jinzhu/gorm"
)

type Fav struct {
	gorm.Model
	UserID  string `json:user_id`
	LyricID string `json:lyric_id`
	Lyric   Lyric  `json:"lyric" gorm:"AssociationForeignKey:ID;ForeignKey:LyricID"`
}

func (e Fav) TableName() string {
	return "favs"
}

func PreLoadFavRelations(db *gorm.DB) *gorm.DB {
	db = db.Preload("Lyric")

	return db
}
