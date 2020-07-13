package main

import (
	_ "encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
)

type Genre struct {
	id    uint
	Title string
	gorm.Model
}

type Game struct {
	id    uint
	Name  string
	Genre Genre `gorm:"foreignkey:id"`
	gorm.Model
}

func main() {
	db, _ := getDB()
	db.DropTableIfExists(&Game{}, &Genre{})
	db.CreateTable(&Game{}, Genre{})

	genre := Genre{
		Title: "Puzzle",
	}

	db.NewRecord(genre)
	db.Create(&genre)

	game := Game{
		Name: "drop Puzzle",
	}

	db.NewRecord(game)
	db.Create(&game)

	router := mux.NewRouter()
	//router.HandleFunc("/games", getAllGames)

	log.Fatal(http.ListenAndServe(":8080", router))

}

//func getAllGames(w http.ResponseWriter, r *http.Request) {
//
//	db, err := getDB()
//
//	if err != nil {
//		panic(err.Error())
//	}
//
//	game := Game{}
//
//	result, err := db.Model(&game).Where("id = ?", "1").Select("id, created_at, deleted_at").Rows()
//
//	if err != nil{
//		panic(err.Error())
//	}
//
//	for result.Next() {
//
//		err = result.Scan(&id, &Name, &Genre{})
//
//		if err != nil{
//			panic(err.Error())
//		}
//
//	}
//
//}

func getDB() (*gorm.DB, error) {

	return gorm.Open("mysql", "root:@/dds_db?charset=utf8&parseTime=True&loc=Local")

}
