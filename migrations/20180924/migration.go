package v20180924

import (
	"github.com/jinzhu/gorm"
)

func Migration(db *gorm.DB) error {
	if err := db.AutoMigrate(&Lyric{}).Error; err != nil {
		return err
	}

	if err := db.AutoMigrate(&Fav{}).Error; err != nil {
		return err
	}

	if err := db.AutoMigrate(&User{}).Error; err != nil {
		return err
	}

	return nil
}