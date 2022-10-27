package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type DB struct {
}

var (
	db *gorm.DB
)

func connect() {
	// conn, err := gorm.Open("postgres", "postgres:0000@/bookStore?charset=utf8&parseTime=True&loc=Local")
	conn, err := gorm.Open("postgres", "user=postgres password=0000 dbname=bookStore sslmode=disable")
	if err != nil {
		panic(err)
	}
	fmt.Println("Connect to database successful...")
	db = conn
}

func GetInstance() *gorm.DB {
	if db == nil {
		connect()
	}
	return db
}
