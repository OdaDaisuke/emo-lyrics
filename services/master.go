package services

import (
	"github.com/OdaDaisuke/emo-lyrics-api/models"
	"github.com/OdaDaisuke/emo-lyrics-api/repositories"
	"github.com/jinzhu/gorm"
)

type MasterService struct {
	dbCtx     *gorm.DB
	lyricRepo *repositories.LyricRepo
}

func NewMasterService(dbCtx *gorm.DB, lyricRepo *repositories.LyricRepo) *MasterService {
	return &MasterService{
		dbCtx:     dbCtx,
		lyricRepo: lyricRepo,
	}
}

func (s *MasterService) CreateLyric(content, title, singer, url string) (*models.Lyric, error) {
	lyric, err := s.lyricRepo.Create(content, title, singer, url)
	if err != nil {
		return nil, err
	}

	return lyric, nil
}

func (s *MasterService) DeleteAllLyrics() error {
	err := s.lyricRepo.DeleteAll()
	if err != nil {
		return err
	}

	return nil
}
