package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"github.com/OdaDaisuke/emo-lyrics-api/models"
	"github.com/OdaDaisuke/emo-lyrics-api/controllers"
	"github.com/OdaDaisuke/emo-lyrics-api/configs"
)

func main() {
	fmt.Println("server running on port", configs.API_SERVER_PORT)

	db := models.NewDBContext()
	defer db.Close()

	db.Set("gorm:table_options", "ENGINE=InnoDB")
	db.AutoMigrate(&models.Lyric{})

	lyricCtrl := controllers.NewLyricCtrl(db)
	masterDataCtrl := controllers.NewMasterDataCtrl(db)
	twitterAuthCtrl := controllers.NewTwitterAuthCtrl(db)

	router := httprouter.New()

	router.GET("/api/v1/lyric", lyricCtrl.GetLyrics())
	router.GET("/api/v1/404_lyric", lyricCtrl.Get404Lyric())
	router.POST("/api/v1/lyric", lyricCtrl.CreateLyric())
	router.DELETE("/api/v1/lyric", lyricCtrl.DeleteLyrics())
	router.POST("/api/v1/master_data", masterDataCtrl.SetMasterData())
	router.GET("/api/v1/auth/get_twitter_auth_url", twitterAuthCtrl.GetAuthUrl())
	router.POST("/api/v1/auth/twitter_verification_code", twitterAuthCtrl.SetVerificationCode())

	servePort := ":" + configs.API_SERVER_PORT
	log.Fatal(http.ListenAndServe(servePort, router))
}
