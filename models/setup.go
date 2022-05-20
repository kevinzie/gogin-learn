package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "root:root@tcp(127.0.0.1:3306)/gogin"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//database, err := gorm.Open("sqlite3", "test.db")

	if err != nil {
		panic("Failed to connect to database!")
	}
	//panic("Failed to connect to daasdasdtabase!")
	database.AutoMigrate(&Book{})

	DB = database

}
