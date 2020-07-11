package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm"
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

	_, err := gorm.Open("mysql", "root:@/dds_db?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err.Error())
	}

	router := mux.NewRouter()
	router.HandleFunc("/", helloWorld)

	log.Fatal(http.ListenAndServe(":8080", router))

}

func helloWorld(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Hello World")

}
