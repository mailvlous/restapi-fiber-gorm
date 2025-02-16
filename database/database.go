package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"fmt"
)

var DB *gorm.DB

func DatabaseInit() {
	var err error

	const MYSQL = "root:BackSpace@tcp(127.0.0.1:3306)/restapi-fiber-gorm?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := MYSQL
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	} 
	fmt.Println("Database connected!")
}
