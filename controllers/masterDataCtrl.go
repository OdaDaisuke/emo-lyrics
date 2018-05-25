package controllers

import (
  "github.com/jinzhu/gorm"
  "github.com/julienschmidt/httprouter"
  "net/http"
  "fmt"
  "encoding/json"
  "io/ioutil"
  "../models"
)

type MasterDataSet struct {
  Lyrics []*models.Lyric `json:lyrics`
}

type MasterDataCtrl struct {
  dbCtx *gorm.DB
}

func NewMasterDataCtrl(dbCtx *gorm.DB) *MasterDataCtrl {
  return &MasterDataCtrl{dbCtx}
}

func (c *MasterDataCtrl) SetMasterData() httprouter.Handle {
  return func (w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    fmt.Println("set masterdata")
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

