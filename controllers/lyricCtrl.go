package controllers

import (
  "github.com/jinzhu/gorm"
  "github.com/julienschmidt/httprouter"
  "net/http"
  "fmt"
  "encoding/json"
  "github.com/emo-lyrics-api/models"
  "github.com/emo-lyrics-api/configs"
)

type LyricCtrl struct {
  dbCtx *gorm.DB
}

func NewLyricCtrl(dbCtx *gorm.DB) *LyricCtrl {
  return &LyricCtrl{dbCtx}
}

func (c *LyricCtrl) GetLyrics() httprouter.Handle {
  return func (w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    fmt.Println("get lyrics")
    lyrics := []models.Lyric{}
    c.dbCtx.Limit(configs.LYRICS_FETCH_LIMITS).Find(&lyrics)
    encoder := json.NewEncoder(w)
    encoder.Encode(lyrics)
  }
}

func (c *LyricCtrl) DeleteLyrics() httprouter.Handle {
  return func (w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    fmt.Println("delete lyrics")
    lyricEx := models.Lyric{}
    c.dbCtx.Delete(lyricEx)
    encoder := json.NewEncoder(w)
    encoder.Encode(lyricEx)
  }
}

func (c *LyricCtrl) CreateLyric() httprouter.Handle {
  return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    fmt.Println("create lyric")
    newLyric := models.Lyric{}
    newLyric.Content = r.FormValue("content")
    newLyric.Title = r.FormValue("title")
    newLyric.Singer = r.FormValue("singer")
    newLyric.Url = r.FormValue("url")
  	c.dbCtx.Create(&newLyric)
    encoder := json.NewEncoder(w)
    encoder.Encode(newLyric)
  }
}
