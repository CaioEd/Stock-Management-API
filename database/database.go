package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	// DB CONNECTION
	dsn := "root@tcp(127.0.0.1:3307)/stock_db"
	var err error

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error opening MySQL database connection ", err)
	}

	fmt.Println("Succesfully connect to the MySQL database!")
}


func Close() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal("Error closing the database connection: ", err)
	}

	if err := sqlDB.Close(); err != nil {
		log.Fatal("Error closing database connection: ", err)
	}
	fmt.Println("Database connection closed")
}