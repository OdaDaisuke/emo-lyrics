package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/OdaDaisuke/emo-lyrics-api/configs"
)

type Factory struct {
	dbCtx *gorm.DB
	appConfig *configs.AppConfig
	LyricRepo *LyricRepo
	FavRepo *FavRepo
	UserRepo *UserRepo
}

func NewFactory(dbCtx *gorm.DB, appConfig *configs.AppConfig) *Factory {
	return &Factory{
		dbCtx: dbCtx,
		LyricRepo: NewLyricRepo(dbCtx, appConfig),
		FavRepo: NewFavRepo(dbCtx),
		UserRepo: NewUserRepo(dbCtx),
	}
}