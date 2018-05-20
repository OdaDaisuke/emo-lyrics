package controllers

import (
  "github.com/jinzhu/gorm"
  "github.com/julienschmidt/httprouter"
  "net/http"
  "fmt"
  "encoding/json"
  "../models"
  "../configs"
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
    setHeader(w, r)
    lyrics := []models.Lyric{}
    c.dbCtx.Limit(configs.LYRICS_FETCH_LIMITS).Find(&lyrics)
    encoder := json.NewEncoder(w)
    encoder.Encode(lyrics)
  }
}

func (c *LyricCtrl) DeleteLyrics() httprouter.Handle {
  return func (w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    fmt.Println("delete lyrics")
    setHeader(w, r)
    c.dbCtx.Delete(models.Lyric{})
    encoder := json.NewEncoder(w)
    encoder.Encode(http.StatusOK)
  }
}

func (c *LyricCtrl) CreateLyric() httprouter.Handle {
  return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    fmt.Println("create lyric")
    setHeader(w, r)
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
