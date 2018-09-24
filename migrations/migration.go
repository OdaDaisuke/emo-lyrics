package migrations

import (
	"github.com/jinzhu/gorm"
	"github.com/OdaDaisuke/emo-lyrics-api/models"
)

func Migration(db *gorm.DB) {
	db.AutoMigrate(&models.Lyric{})
}
