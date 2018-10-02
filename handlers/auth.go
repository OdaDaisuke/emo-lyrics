package handlers

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type AuthHandler struct {
	dbCtx *gorm.DB
}

func NewAuthHandler(dbCtx *gorm.DB) *AuthHandler {
	return &AuthHandler{dbCtx}
}

func (c *AuthHandler) ApiAuth() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Println("api auth")
		setHeader(w, r)
		// JWT Authentication
	}
}
