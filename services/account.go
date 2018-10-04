package services

import (
	"github.com/OdaDaisuke/emo-lyrics-api/interfaces"
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

func (s *AccountService) Signup(params *interfaces.SignupParams) *models.User {
	user := &models.User{
		TwitterId:            params.TwitterId,
		Lang:                 params.Lang,
		Location:             params.Location,
		Name:                 params.Name,
		ProfileBannerUrl:     params.ProfileBannerUrl,
		ProfileImageUrlHttps: params.ProfileImageUrlHttps,
		Protected:            params.Protected,
		ScreenName:           params.ScreenName,
		Url:                  params.Url,
	}
	s.dbCtx.Create(user)
	return user
}

func (s *AccountService) GetMe(params *interfaces.GetMeParams) (*models.User, error) {
	user := &models.User{
		TwitterId: params.TwitterId,
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

func (s *AccountService) PostFav(params *interfaces.PostFavParams) (*models.Fav, error) {
	fav, err := s.favRepo.PostFav(params)
	if err != nil {
		return nil, err
	}
	return fav, nil
}

func (s *AccountService) UnFav(params *interfaces.UnFavParams) (*models.Fav, error) {
	fav, err := s.favRepo.UnFav(params)
	if err != nil {
		return nil, err
	}
	return fav, nil
}

func (s *AccountService) GetFavList(params *interfaces.GetFavListParams) ([]*models.Fav, error) {
	fav, err := s.favRepo.GetMyFavList(params)
	if err != nil {
		return nil, err
	}
	return fav, nil
}
