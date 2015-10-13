package database

import (
	"github.com/jinzhu/gorm"
	"fmt"
	_ "github.com/lib/pq"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	_ "./../../model"
)

var db *gorm.DB = nil

func GetDB() gorm.DB {
	if db != nil {
		return *db
	}
	dbInstance, err := gorm.Open("mysql", "root:123@localhost:3306/dbname")
	dbInstance.LogMode(false)
	if(err != nil) {
		fmt.Println("[ERROR] Database Connection Failed")
		panic(err)
	}
	db = &dbInstance
	return *db;
}


func CloseDB() bool {
	if db == nil {
		fmt.Println("[ERROR] No Database Initialized")
		return false
	}
	err := (*db).Close()
	if(err != nil) {
		fmt.Println("[ERROR] Database Close Failed")
		return false
	}
	db = nil
	return true
}


func RegisterModel(model interface{}) {
	tableDb := GetDB()
	tableDb.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(model)
	tableDb.AutoMigrate(&model)
}

