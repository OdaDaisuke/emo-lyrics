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
	fmt.Println("server running on port", configs.API_SERVER_PORT)

	db := models.NewDBContext()
	defer db.Close()

	migrations.Migration(db)

	// Init handlers
	lyricHandler := handlers.NewLyricHandler(db)
	masterDataHandler := handlers.NewMasterDataHandler(db)

	router := httprouter.New()

	// lyric
	router.GET("/api/v1/lyric", lyricHandler.GetLyrics())
	router.GET("/api/v1/404_lyric", lyricHandler.Get404Lyric())
	router.POST("/api/v1/lyric", lyricHandler.CreateLyric())
	router.DELETE("/api/v1/lyric", lyricHandler.DeleteLyrics())

	// Master data
	router.POST("/api/v1/master_data", masterDataHandler.SetMasterData())

	servePort := ":" + configs.API_SERVER_PORT
	log.Fatal(http.ListenAndServe(servePort, router))
	appengine.Main()
}
