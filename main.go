package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"google.golang.org/appengine"
	"github.com/OdaDaisuke/emo-lyrics-api/models"
	"github.com/OdaDaisuke/emo-lyrics-api/configs"
	"github.com/OdaDaisuke/emo-lyrics-api/handlers"
	"github.com/OdaDaisuke/emo-lyrics-api/migrations"
)

func main() {
	appConfigs := configs.LoadAppConfig()
	fmt.Println("server running on port", appConfigs.ApiServerPort)

	db := models.NewDBContext()
	defer func() {
		// todo: add child transaction rollback
		db.Close()
	}()

	migrations.Migration(db)

	// Init handlers
	lyricHandler := handlers.NewLyricHandler(db, appConfigs)
	masterDataHandler := handlers.NewMasterDataHandler(db, appConfigs)
	accountHandler := handlers.NewAccountHandler(db, appConfigs)

	router := httprouter.New()

	// lyric
	router.GET("/api/v1/lyric", lyricHandler.GetLyrics())
	router.GET("/api/v1/404_lyric", lyricHandler.Get404Lyric())
	router.POST("/api/v1/lyric", lyricHandler.CreateLyric())
	router.DELETE("/api/v1/lyric", lyricHandler.DeleteLyrics())

	// Master data
	router.POST("/api/v1/master_data", masterDataHandler.SetMasterData())

	// account
	router.POST("/api/v1/account", accountHandler.Signup())
	router.GET("/api/v1/signin", accountHandler.Signin())

	// me

	servePort := ":" + appConfigs.ApiServerPort
	log.Fatal(http.ListenAndServe(servePort, router))
	appengine.Main()
}
