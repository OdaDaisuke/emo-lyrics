package handlers

import (
	"encoding/json"
	"github.com/OdaDaisuke/emo-lyrics-api/configs"
	"github.com/OdaDaisuke/emo-lyrics-api/repositories"
	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type LyricHandler struct {
	dbCtx       *gorm.DB
	repoFactory *repositories.Factory
	appConfig   *configs.AppConfig
}

func NewLyricHandler(dbCtx *gorm.DB, repoFactory *repositories.Factory, appConfig *configs.AppConfig) *LyricHandler {
	return &LyricHandler{
		dbCtx:       dbCtx,
		repoFactory: repoFactory,
		appConfig:   appConfig,
	}
}

func (c *LyricHandler) Get404Lyric() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		setHeader(w, r)
		encoder := json.NewEncoder(w)

		lyric, err := c.repoFactory.LyricRepo.Get404()
		if err != nil {
			w.WriteHeader(500)
			return
		}

		encoder.Encode(lyric)
	}
}

func (c *LyricHandler) GetLyrics() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		setHeader(w, r)

		lyrics, err := c.repoFactory.LyricRepo.GetAll()
		if err != nil {
			w.WriteHeader(500)
			return
		}

		encoder := json.NewEncoder(w)
		encoder.Encode(lyrics)
	}
}

func (c *LyricHandler) DeleteLyrics() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		setHeader(w, r)
		c.repoFactory.LyricRepo.DeleteAll()
		w.WriteHeader(http.StatusOK)
	}
}

func (c *LyricHandler) CreateLyric() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		setHeader(w, r)

		lyric := r.FormValue("lyric")
		title := r.FormValue("title")
		singer := r.FormValue("singer")
		url := r.FormValue("url")

		newLyric, err := c.repoFactory.LyricRepo.Create(lyric, title, singer, url)
		if err != nil {
			w.WriteHeader(500)
			return
		}

		encoder := json.NewEncoder(w)
		encoder.Encode(newLyric)
	}
}
