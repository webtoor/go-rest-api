package config

import (
	"fmt"

	"github.com/webtoor/go-rest-api/helper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const DB_USERNAME = "root"
const DB_PASSWORD = "Rahasia123"
const DB_NAME = "go_rest_api"
const DB_HOST = "192.168.1.6"
const DB_PORT = "3306"

var DB *gorm.DB

func InitDb() *gorm.DB {
	DB = connectDB()
	return DB
}

func connectDB() *gorm.DB {
	var err error
	dsn := DB_USERNAME + ":" + DB_PASSWORD + "@tcp" + "(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?" + "parseTime=true&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Error connecting to database : error")
		helper.PanicIfError(err)
	}

	return db
}