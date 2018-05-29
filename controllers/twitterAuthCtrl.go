package controllers

import (
  "fmt"
  "github.com/jinzhu/gorm"
  "github.com/julienschmidt/httprouter"
  "net/http"
  "encoding/json"
  "github.com/OdaDaisuke/emo-lyrics-api/gateways"
)

type TwitterAuthCtrl struct {
  dbCtx *gorm.DB
}

func NewTwitterAuthCtrl(dbCtx *gorm.DB) *TwitterAuthCtrl {
  return &TwitterAuthCtrl{dbCtx}
}

func (c *TwitterAuthCtrl) GetAuthUrl() httprouter.Handle {
  return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    fmt.Println("get twitter auth url")
    setHeader(w, r)
    twitterGW := gateways.NewTwitterGW(r)

    encoder := json.NewEncoder(w)
    encoder.Encode(twitterGW.GetAuthUrl())
  }
}

func (c *TwitterAuthCtrl) SetVerificationCode() httprouter.Handle {
  return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    fmt.Println("set twitter verification code")
    setHeader(w, r)

    twitterGW := gateways.NewTwitterGW(r)
    verifyRs := twitterGW.SetVerificationCode()
    res := false
    if verifyRs {
      res = true
    }

    encoder := json.NewEncoder(w)
    encoder.Encode(res)
  }
}
