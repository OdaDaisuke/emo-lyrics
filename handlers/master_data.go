package handlers

import (
	"encoding/json"
	"github.com/OdaDaisuke/emo-lyrics-api/configs"
	"github.com/OdaDaisuke/emo-lyrics-api/models"
	"github.com/OdaDaisuke/emo-lyrics-api/repositories"
	"github.com/OdaDaisuke/emo-lyrics-api/services"
	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
)

type MasterDataSet struct {
	Lyrics []*models.Lyric `json:lyrics`
}

type MasterDataHandler struct {
	dbCtx         *gorm.DB
	masterService *services.MasterService
	appConfig     *configs.AppConfig
}

func NewMasterDataHandler(dbCtx *gorm.DB, repoFactory *repositories.Factory, appConfig *configs.AppConfig) *MasterDataHandler {
	masterService := services.NewMasterService(dbCtx, repoFactory.LyricRepo)

	return &MasterDataHandler{
		dbCtx:         dbCtx,
		masterService: masterService,
		appConfig:     appConfig,
	}
}

func (c *MasterDataHandler) SetMasterData() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		setHeader(w, r)
		encoder := json.NewEncoder(w)
		b, _ := ioutil.ReadAll(r.Body)
		var dataSet MasterDataSet
		if err := json.Unmarshal(b, &dataSet); err != nil {
			encoder.Encode(http.StatusInternalServerError)
		}
		for _, lyric := range dataSet.Lyrics {
			c.dbCtx.Create(lyric)
		}
		encoder.Encode(http.StatusOK)
	}
}

func (c *MasterDataHandler) DeleteLyrics() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		setHeader(w, r)
		c.masterService.DeleteAllLyrics()
		w.WriteHeader(http.StatusOK)
	}
}

func (c *MasterDataHandler) CreateLyric() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		setHeader(w, r)

		lyric := r.FormValue("lyric")
		title := r.FormValue("title")
		singer := r.FormValue("singer")
		url := r.FormValue("url")

		newLyric, err := c.masterService.CreateLyric(lyric, title, singer, url)
		if err != nil {
			w.WriteHeader(500)
			return
		}

		encoder := json.NewEncoder(w)
		encoder.Encode(newLyric)
	}
}
