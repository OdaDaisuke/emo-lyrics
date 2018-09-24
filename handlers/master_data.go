package handlers

import (
  "github.com/jinzhu/gorm"
  "github.com/julienschmidt/httprouter"
  "net/http"
  "encoding/json"
  "io/ioutil"
  "github.com/OdaDaisuke/emo-lyrics-api/models"
)

type MasterDataSet struct {
  Lyrics []*models.Lyric `json:lyrics`
}

type MasterDataHandler struct {
  dbCtx *gorm.DB
}

func NewMasterDataHandler(dbCtx *gorm.DB) *MasterDataHandler {
  return &MasterDataHandler{dbCtx}
}

func (c *MasterDataHandler) SetMasterData() httprouter.Handle {
  return func (w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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

