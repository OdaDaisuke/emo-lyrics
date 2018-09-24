package migrations

import (
	"github.com/jinzhu/gorm"
	"github.com/OdaDaisuke/emo-lyrics-api/migrations/20180924"
)

func Migration(db *gorm.DB) {
	if err := v20180924.Migration(db); err != nil {
		panic(err)
	}
}
