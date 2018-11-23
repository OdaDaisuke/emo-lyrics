package handlers

import (
	"encoding/json"
	"github.com/OdaDaisuke/emo-lyrics-api/configs"
	"github.com/OdaDaisuke/emo-lyrics-api/repositories"
	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type DefaultHandler struct {
	dbCtx       *gorm.DB
	repoFactory *repositories.Factory
	appConfig   *configs.AppConfig
}

func NewDefaultHandler(dbCtx *gorm.DB, repoFactory *repositories.Factory, appConfig *configs.AppConfig) *DefaultHandler {
	return &DefaultHandler{
		dbCtx:       dbCtx,
		repoFactory: repoFactory,
		appConfig:   appConfig,
	}
}

func (c *DefaultHandler) Get404Lyric() httprouter.Handle {
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

func (c *DefaultHandler) HealthCheck() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.WriteHeader(200)
	}
}
