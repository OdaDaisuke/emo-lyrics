package repositories

import (
	"github.com/OdaDaisuke/emo-lyrics-api/models"
	"github.com/jinzhu/gorm"
)

type FavRepo struct {
	dbCtx *gorm.DB
}

func NewFavRepo(dbCtx *gorm.DB) *FavRepo {
	return &FavRepo{
		dbCtx: dbCtx,
	}
}

func (r *FavRepo) PostFav(lyricId string) (*models.Fav, error) {
	fav := &models.Fav{
		LyricID: lyricId,
	}
	if err := r.dbCtx.Where(fav).Last(fav).Error; err != nil {
		return nil, err
	}
	return fav, nil
}

func (r *FavRepo) UnFav(lyricId string) (*models.Fav, error) {
	fav := &models.Fav{
		LyricID: lyricId,
	}
	if err := r.dbCtx.Where(fav).Update(fav).Error; err != nil {
		return nil, err
	}
	return fav, nil
}

func (r *FavRepo) GetMyFavList() ([]*models.Fav, error) {
	fav := []*models.Fav{}
	if err := r.dbCtx.Model(fav).Find(fav).Error; err != nil {
		return nil, err
	}
	return fav, nil
}
