package main

import (
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
)

type Game struct {
	id uint
	gorm.Model
	Name  string
	Genre []Genre
}

type Genre struct {
	id    uint
	Title string
}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/games", getAllGames)

	log.Fatal(http.ListenAndServe(":8080", router))

}

func getAllGames(w http.ResponseWriter, r *http.Request) {

	db, err := getDB()

	if err != nil {
		panic(err.Error())
	}

	game := Game{}

	w.Header().Set("Content-Type", "application/json")

	result := db.Debug().Find(&game)

	json.NewEncoder(w).Encode(result)

}

func getDB() (*gorm.DB, error) {

	return gorm.Open("mysql", "root:@/dds_db?charset=utf8&parseTime=True&loc=Local")

}
