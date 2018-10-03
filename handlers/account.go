package handlers

import (
	"encoding/json"
	"github.com/OdaDaisuke/emo-lyrics-api/configs"
	"github.com/OdaDaisuke/emo-lyrics-api/models"
	"github.com/OdaDaisuke/emo-lyrics-api/repositories"
	"github.com/OdaDaisuke/emo-lyrics-api/services"
	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type AccountHandler struct {
	dbCtx          *gorm.DB
	accountService *services.AccountService
	appConfig      *configs.AppConfig
}

func NewAccountHandler(dbCtx *gorm.DB, repoFactory *repositories.Factory, appConfig *configs.AppConfig) *AccountHandler {
	accountService := services.NewAccountService(dbCtx, repoFactory.LyricRepo, repoFactory.FavRepo)

	return &AccountHandler{
		dbCtx:          dbCtx,
		accountService: accountService,
		appConfig:      appConfig,
	}
}

func (c *AccountHandler) Signup() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		setHeader(w, r)

		token := r.FormValue("token")
		user := c.accountService.Signup(token)

		encoder := json.NewEncoder(w)
		encoder.Encode(user)
	}
}

func (c *AccountHandler) Signin() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		setHeader(w, r)

		token := r.FormValue("token")
		user, err := c.accountService.Signin(token)
		if err != nil {
			w.WriteHeader(500)
		}

		encoder := json.NewEncoder(w)
		encoder.Encode(user)
	}
}

func (c *AccountHandler) PostFav() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		setHeader(w, r)

		lyricId := r.FormValue("lyric_id")
		fav, err := c.accountService.PostFav(lyricId)
		if err != nil {
			w.WriteHeader(500)
		}

		encoder := json.NewEncoder(w)
		encoder.Encode(fav)
	}
}

func (c *AccountHandler) UnFav() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		setHeader(w, r)

		fav := &models.Fav{
			//LyricID: r.FormValue("lyric_id"),
		}
		c.dbCtx.Where(fav).Delete(fav)

		encoder := json.NewEncoder(w)
		encoder.Encode(nil)
	}
}

func (c *AccountHandler) GetFavList() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		setHeader(w, r)

		fav := &models.Fav{
			//UserID:: r.FormValue("user_id"),
			//LyricID: r.FormValue("lyric_id"),
		}
		c.dbCtx.Where(fav).Find(fav)

		encoder := json.NewEncoder(w)
		encoder.Encode(fav)
	}
}

func (c *AccountHandler) GetLyrics() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		setHeader(w, r)

		lyrics, err := c.accountService.GetLyrics()
		if err != nil {
			w.WriteHeader(500)
			return
		}

		encoder := json.NewEncoder(w)
		encoder.Encode(lyrics)
	}
}
