package handlers

import (
  "github.com/jinzhu/gorm"
  "github.com/julienschmidt/httprouter"
  "net/http"
  "encoding/json"
  "github.com/OdaDaisuke/emo-lyrics-api/models"
  "github.com/OdaDaisuke/emo-lyrics-api/configs"
)

type LyricHandler struct {
  dbCtx *gorm.DB
  appConfig *configs.AppConfig
}

func NewLyricHandler(dbCtx *gorm.DB, appConfig *configs.AppConfig) *LyricHandler {
  return &LyricHandler{
    dbCtx: dbCtx,
    appConfig: appConfig,
  }
}

func (c *LyricHandler) Get404Lyric() httprouter.Handle {
  return func (w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    setHeader(w, r)
    encoder := json.NewEncoder(w)
    lyric := &models.Lyric{
      Url: "https://www.youtube.com/watch?v=EvBDa4TX3Bo",
    }
    c.dbCtx.Model(lyric).Where(lyric).Find(lyric).Last(lyric)

    encoder.Encode(lyric)
  }
}

func (c *LyricHandler) GetLyrics() httprouter.Handle {
  return func (w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    setHeader(w, r)
    lyrics := []models.Lyric{}
    c.dbCtx.Limit(c.appConfig.LyricFetchLimits).Find(&lyrics)
    encoder := json.NewEncoder(w)
    encoder.Encode(lyrics)
  }
}

func (c *LyricHandler) DeleteLyrics() httprouter.Handle {
  return func (w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    setHeader(w, r)
    c.dbCtx.Delete(models.Lyric{})
    encoder := json.NewEncoder(w)
    encoder.Encode(http.StatusOK)
  }
}

func (c *LyricHandler) CreateLyric() httprouter.Handle {
  return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    setHeader(w, r)

    newLyric := models.Lyric{}
    newLyric.Lyric = r.FormValue("lyric")
    newLyric.Title = r.FormValue("title")
    newLyric.Singer = r.FormValue("singer")
    newLyric.Url = r.FormValue("url")
    c.dbCtx.Create(&newLyric)

    encoder := json.NewEncoder(w)
    encoder.Encode(newLyric)
  }
}
