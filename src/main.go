package main

import (
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm"
	"log"
	"net/http"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/", helloWorld)

	log.Fatal(http.ListenAndServe(":8080", router))

}

func helloWorld(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Hello World")

}
