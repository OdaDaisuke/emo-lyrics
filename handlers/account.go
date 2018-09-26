package handlers

import (
	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"github.com/OdaDaisuke/emo-lyrics-api/models"
	"encoding/json"
)

type AccountHandler struct {
	dbCtx *gorm.DB
}

func NewAccountHandler(dbCtx *gorm.DB) *AccountHandler {
	return &AccountHandler{dbCtx}
}

func (c *AccountHandler) Signup() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		setHeader(w, r)

		user := models.User{
			Token: r.FormValue("token"),
		}
		c.dbCtx.Create(&user)

		encoder := json.NewEncoder(w)
		encoder.Encode(user)
	}
}

func (c *AccountHandler) Signin() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		setHeader(w, r)

		user := &models.User{
			Token: r.FormValue("token"),
		}
		c.dbCtx.Where(user).Last(user)

		encoder := json.NewEncoder(w)
		encoder.Encode(user)
	}
}

func(c *AccountHandler) PostFav() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		setHeader(w, r)

		fav := &models.Fav{
			//LyricID: r.FormValue("lyric_id"),
		}
		c.dbCtx.Where(fav).Last(fav)

		encoder := json.NewEncoder(w)
		encoder.Encode(fav)
	}
}

func(c *AccountHandler) UnFav() httprouter.Handle {
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

func(c *AccountHandler) GetFavList() httprouter.Handle {
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