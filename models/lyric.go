package models

import (
  "github.com/jinzhu/gorm"
)

type Lyric struct {
  gorm.Model
  Lyric string `json:lyric`
  Title string `json:title`
  Singer string `json:singer`
  Url string `json:url`
  Genre string `json:genre`
}

func (e Lyric) TableName() string {
  return "lyrics"
}
