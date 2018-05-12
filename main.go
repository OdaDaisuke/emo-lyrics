package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"github.com/emo-lyrics-api/models"
	"github.com/emo-lyrics-api/controllers"
)

func AuthHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "JWT Authentication.")
}

func main() {
	fmt.Println("server running on port 8888")
	db := models.GormConnect()
	defer db.Close()

	db.Set("gorm:table_options", "ENGINE=InnoDB")
	db.AutoMigrate(&models.Lyric{})
	lyricCtrl := controllers.NewLyricCtrl(db)

	router := httprouter.New()

	router.GET("/api/v1/lyric", lyricCtrl.GetLyrics())
	router.POST("/api/v1/lyric", lyricCtrl.CreateLyric())
	router.DELETE("/api/v1/lyric", lyricCtrl.DeleteLyrics())
	router.GET("/api/v1/auth", AuthHandler)

	log.Fatal(http.ListenAndServe(":8888", router))
}
