package repositories

import "github.com/jinzhu/gorm"

type FavRepo struct {
	dbCtx *gorm.DB
}

func NewFavRepo(dbCtx *gorm.DB) *FavRepo {
	return &FavRepo{
		dbCtx: dbCtx,
	}
}