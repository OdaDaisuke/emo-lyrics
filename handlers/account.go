package handlers

import (
	"encoding/json"
	"github.com/OdaDaisuke/emo-lyrics-api/configs"
	"github.com/OdaDaisuke/emo-lyrics-api/interfaces"
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

		params := &interfaces.SignupParams{
			TwitterId:            r.Form.Get("twitter_id"),
			Lang:                 r.Form.Get("lang"),
			Location:             r.Form.Get("location"),
			Name:                 r.Form.Get("name"),
			ProfileBannerUrl:     r.Form.Get("profile_banner_url"),
			ProfileImageUrlHttps: r.Form.Get("profile_image_url_https"),
			Protected:            r.Form.Get("protected"),
			ScreenName:           r.Form.Get("screen_name"),
			Url:                  r.Form.Get("url"),
		}
		user := c.accountService.Signup(params)

		encoder := json.NewEncoder(w)
		encoder.Encode(user)
	}
}

func (c *AccountHandler) GetMe() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		setHeader(w, r)

		params := &interfaces.GetMeParams{
			TwitterId: r.FormValue("twitter_id"),
		}
		user, err := c.accountService.GetMe(params)
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

		params := &interfaces.PostFavParams{
			UserId:  r.Form.Get("user_id"),
			LyricId: r.Form.Get("lyric_id"),
		}
		fav, err := c.accountService.PostFav(params)
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

		params := &interfaces.UnFavParams{
			UserId:  r.Form.Get("user_id"),
			LyricId: r.Form.Get("lyric_id"),
		}
		_, err := c.accountService.UnFav(params)
		if err != nil {
			w.WriteHeader(500)
		}

		encoder := json.NewEncoder(w)
		encoder.Encode(nil)
	}
}

func (c *AccountHandler) GetFavList() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		setHeader(w, r)

		params := &interfaces.GetFavListParams{
			UserId: r.FormValue("user_id"),
		}
		favList, err := c.accountService.GetFavList(params)
		if err != nil {
			w.WriteHeader(500)
		}

		encoder := json.NewEncoder(w)
		encoder.Encode(favList)
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
