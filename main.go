package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"github.com/emo-lyrics-api/models"
	"github.com/emo-lyrics-api/controllers"
	"github.com/emo-lyrics-api/configs"
)

func main() {
	fmt.Println("server running on port", configs.API_SERVER_PORT)

	db := models.GormConnect()
	defer db.Close()

	db.Set("gorm:table_options", "ENGINE=InnoDB")
	db.AutoMigrate(&models.Lyric{})

	lyricCtrl := controllers.NewLyricCtrl(db)
	authCtrl := controllers.NewAuthCtrl(db)

	router := httprouter.New()

	router.GET("/api/v1/lyric", lyricCtrl.GetLyrics())
	router.POST("/api/v1/lyric", lyricCtrl.CreateLyric())
	router.DELETE("/api/v1/lyric", lyricCtrl.DeleteLyrics())
	router.GET("/api/v1/auth", authCtrl.ApiAuth())

	servePort := ":" + configs.API_SERVER_PORT
	log.Fatal(http.ListenAndServe(servePort, router))
}
