package models

import (
  "github.com/jinzhu/gorm"
)

type Lyric struct {
  gorm.Model
  Content string `json:content`
  Title string `json:title`
  Singer string `json:singer`
  Url string `json:url`
}

func (e Lyric) TableName() string {
  return "lyrics"
}
