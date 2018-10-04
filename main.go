package main

import (
	"fmt"
	"github.com/OdaDaisuke/emo-lyrics-api/configs"
	"github.com/OdaDaisuke/emo-lyrics-api/handlers"
	"github.com/OdaDaisuke/emo-lyrics-api/migrations"
	"github.com/OdaDaisuke/emo-lyrics-api/models"
	"github.com/OdaDaisuke/emo-lyrics-api/repositories"
	"github.com/julienschmidt/httprouter"
	"google.golang.org/appengine"
	"log"
	"net/http"
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

	// Init factories
	repoFactory := repositories.NewFactory(db, appConfigs)

	// Init handlers
	defaultHandler := handlers.NewDefaultHandler(db, repoFactory, appConfigs)
	masterDataHandler := handlers.NewMasterDataHandler(db, repoFactory, appConfigs)
	accountHandler := handlers.NewAccountHandler(db, repoFactory, appConfigs)

	// Init router
	router := httprouter.New()

	router.GET("/api/v1/lyric", accountHandler.GetLyrics())
	router.GET("/api/v1/404_lyric", defaultHandler.Get404Lyric())

	// Master data
	router.POST("/api/v1/master_data", masterDataHandler.SetMasterData())
	router.POST("/api/v1/lyric", masterDataHandler.CreateLyric())
	router.DELETE("/api/v1/lyric", masterDataHandler.DeleteLyrics())

	// account
	router.POST("/api/v1/account", accountHandler.Signup())
	router.GET("/api/v1/account/me", accountHandler.GetMe())

	// me

	servePort := ":" + appConfigs.ApiServerPort
	log.Fatal(http.ListenAndServe(servePort, router))
	appengine.Main()
}
