package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/OdaDaisuke/emo-lyrics-api/models"
	"github.com/OdaDaisuke/emo-lyrics-api/configs"
)

type LyricRepo struct {
	dbCtx *gorm.DB
	appConfig *configs.AppConfig
}

func NewLyricRepo(dbCtx *gorm.DB, appConfig *configs.AppConfig) *LyricRepo {
	return &LyricRepo{
		dbCtx: dbCtx,
		appConfig: appConfig,
	}
}

func (l *LyricRepo) Get404() (*models.Lyric, error) {
	lyric := &models.Lyric{
		Url: "https://www.youtube.com/watch?v=EvBDa4TX3Bo",
	}
	err := l.dbCtx.Model(lyric).Where(lyric).Find(lyric).Last(lyric).Error
	if err != nil {
		return nil, err
	}

	return lyric, nil
}

func (l *LyricRepo) GetAll() ([]*models.Lyric, error) {
	lyrics := []*models.Lyric{}
	err := l.dbCtx.Limit(l.appConfig.LyricFetchLimits).Find(&lyrics).Error
	if err != nil {
		return nil, err
	}

	return lyrics, nil
}

func (l *LyricRepo) Create(lyric, title, singer, url string) (*models.Lyric, error) {
	newLyric := &models.Lyric{
		Lyric: "",
		Title: "",
		Singer: "",
		Url: "",
	}
	err := l.dbCtx.Create(newLyric).Error
	if err != nil {
		return nil, err
	}

	return newLyric, nil
}

func (l *LyricRepo) DeleteAll() error {
	l.dbCtx.Delete(models.Lyric{})
	return nil
}
