package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

const (
	connStr = "root:@/dds_db?charset=utf8&parseTime=True&loc=Local"
)

type (
	User struct {
		gorm.Model
		Name     string
		Username string `gorm:"not null"`
		Password string `gorm:"not null"`
		Messages []Message
	}

	Message struct {
		Body   string `gorm:"not null`
		User   User
		UserID uint
		gorm.Model
	}
)

func main() {

	db, err := gorm.Open("mysql", connStr)
	db.DropTableIfExists(&Message{}, &User{})

	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Message{})

	var user = &User{Name: "Peter Jones", Username: "pj@email.com", Password: "password"}

	if err := db.Model(&User{}).Create(user).Error; err != nil {
		panic(err.Error())
	}

	message := &Message{Body: "Hi"}

	db.Model(user).Find(user)
	db.Model(user).Association("Messages").Append(message)
}

//import (
//	_ "encoding/json"
//	_ "github.com/go-sql-driver/mysql"
//)

//type Genre struct {
//	Title string
//	gorm.Model
//}
//
//type Game struct {
//	Name  string
//	Genre Genre
//	gorm.Model
//}
//
//func main() {
//	db, _ := getDB()
//	db.DropTableIfExists(&Game{}, &Genre{})
//
//	db.AutoMigrate(&Game{})
//	db.AutoMigrate(&Genre{})
//
//	game := &Game{
//		Name: "Drop Puzzle",
//	}
//
//	if err := db.Model(&Game{}).Create(game).Error; err != nil{
//
//		panic(err.Error())
//
//	}
//
//	genre := &Genre{
//		Title: "Arcade",
//	}
//
//	db.Model(&Game{}).Find(game)
//	db.Model(game).Association("Genre").Append(genre)
//
//
//
//	router := mux.NewRouter()
//	//router.HandleFunc("/games", getAllGames)
//
//	log.Fatal(http.ListenAndServe(":8080", router))
//
//}

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

//func getDB() (*gorm.DB, error) {
//
//	return gorm.Open("mysql", "root:@/dds_db?charset=utf8&parseTime=True&loc=Local")
//
//}
