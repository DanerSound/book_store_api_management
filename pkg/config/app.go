package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	//localsetup:                 "user":pass    @tpc(localhost:port)/dbnm?charset=utf8&parseTime=True&loc=Local
	d, err := gorm.Open("mysql", "myroot:password@tcp(localhost:3306)/test?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println("errore while connection to database  on app.go: ", err)
		panic(err)
	}
	db = d
	fmt.Printf("Connection to database success...")
}

func GetDB() *gorm.DB {
	return db
}
