package services

import "github.com/jinzhu/gorm"

type LyricService struct {
	dbCtx *gorm.DB
}

type Lyric interface {
	Create(title string) error
	Get(id uint) error
	Update(id uint) error
	Delete(id uint) error
}

func NewLyricService(dbCtx *gorm.DB) *LyricService {
	return &LyricService{
		dbCtx: dbCtx,
	}
}