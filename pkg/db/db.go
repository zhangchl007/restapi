package db

import (
	"fmt"
	"github.com/zhangchl007/restapi/pkg/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

var db *gorm.DB
var err error

func DBURL() string {
	dbConfig := GetDBConfig()
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.User,
		dbConfig.Password,
		dbConfig.DBName,
		dbConfig.SSLMode,
	)
}
func Init() {
	dbinfo := DBURL()
	fmt.Println(dbinfo)
	db, err = gorm.Open("postgres", dbinfo)
	if err != nil {
		log.Println("Failed to connect to database")
		panic(err)
	}
	log.Println("Database connected")

	if !db.HasTable(&models.TodoModel{}) {
		err := db.CreateTable(&models.TodoModel{})
		if err != nil {
			log.Println("Table already exists")
		}
	}
    //#	defer CloseDB()
	db.AutoMigrate(&models.TodoModel{})
}
//GetDB ...
func GetDB() *gorm.DB {
	return db
}

func CloseDB() {
	db.Close()
}

