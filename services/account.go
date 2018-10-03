package services

import (
	"github.com/OdaDaisuke/emo-lyrics-api/models"
	"github.com/OdaDaisuke/emo-lyrics-api/repositories"
	"github.com/jinzhu/gorm"
)

type AccountService struct {
	dbCtx     *gorm.DB
	lyricRepo *repositories.LyricRepo
	favRepo   *repositories.FavRepo
}

func NewAccountService(dbCtx *gorm.DB, lyricRepo *repositories.LyricRepo, favRepo *repositories.FavRepo) *AccountService {
	return &AccountService{
		dbCtx:     dbCtx,
		lyricRepo: lyricRepo,
		favRepo:   favRepo,
	}
}

func (s *AccountService) Signup(token string) *models.User {
	user := &models.User{
		Token: token,
	}
	s.dbCtx.Create(user)
	return user
}

func (s *AccountService) Signin(token string) (*models.User, error) {
	user := &models.User{
		Token: token,
	}
	if err := s.dbCtx.Where(user).Last(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (s *AccountService) GetLyrics() ([]*models.Lyric, error) {
	lyrics, err := s.lyricRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return lyrics, nil
}

func (s *AccountService) PostFav(lyricId string) (*models.Fav, error) {
	fav, err := s.favRepo.PostFav(lyricId)
	if err != nil {
		return nil, err
	}
	return fav, nil
}

func (s *AccountService) UnFav(lyricId string) (*models.Fav, error) {
	fav, err := s.favRepo.UnFav(lyricId)
	if err != nil {
		return nil, err
	}
	return fav, nil
}

func (s *AccountService) GetFavList() ([]*models.Fav, error) {
	fav, err := s.favRepo.GetMyFavList()
	if err != nil {
		return nil, err
	}
	return fav, nil
}
