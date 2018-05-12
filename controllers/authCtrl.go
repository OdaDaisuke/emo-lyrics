package controllers

import (
  "github.com/jinzhu/gorm"
  "github.com/julienschmidt/httprouter"
  "net/http"
  "fmt"
)

type AuthCtrl struct {
  dbCtx *gorm.DB
}

func NewAuthCtrl(dbCtx *gorm.DB) *AuthCtrl {
  return &AuthCtrl{dbCtx}
}

func (c *AuthCtrl) ApiAuth() httprouter.Handle {
  return func (w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    fmt.Println("api auth")
    // JWT Authentication
  }
}
