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

func (c *AccountHandler) Signin() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		setHeader(w, r)

		user := models.User{
			HandleName: r.FormValue("handle_name"),
		}
		c.dbCtx.Create(&user)

		encoder := json.NewEncoder(w)
		encoder.Encode(user)
	}
}

func(c *AccountHandler) PostFav() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		setHeader(w, r)
	}
}

func(c *AccountHandler) UnFav() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	}
}

func(c *AccountHandler) GetFavList() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	}
}