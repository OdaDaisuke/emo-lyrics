package repositories

import "github.com/jinzhu/gorm"

type LyricRepo struct {
	dbCtx *gorm.DB
}

func NewLyricRepo(dbCtx *gorm.DB) *LyricRepo {
	return &LyricRepo{
		dbCtx: dbCtx,
	}
}