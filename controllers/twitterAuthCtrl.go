package controllers

/*
import (
  "fmt"
  "github.com/jinzhu/gorm"
  "github.com/julienschmidt/httprouter"
  "net/http"
  "encoding/json"
  // "github.com/emo-lyrics-api/models"
  "../gateways"
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
    twitterGW := gateways.NewTwitterGW()
    authConnect := twitterGW.GetConnect()
    at, err := twitterGW.GetAccessToken(
        &oauth.Credentials{
            Token:  c.CruSession.Get("request_token").(string),
            Secret: c.CruSession.Get("request_token_secret").(string),
        },
        request.Verifier,
    )
    if err != nil {
        panic(err)
    }

    encoder := json.NewEncoder(w)
    encoder.Encode(at)
  }
}
*/
