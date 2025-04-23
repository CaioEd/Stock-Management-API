package database

import (
	"fmt"
	"log"
	"os"
	"stock_api/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

func Init() {
	// DB CONNECTION
	err := godotenv.Load()
	if err != nil {
		log.Println("Aviso: .env não encontrado, usando variáveis do sistema.")
	}

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, dbName)


	DB, err = gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatal("Error opening MySQL database connection ", err)
	}

	fmt.Println("Succesfully connect to the MySQL database!")

	err = DB.AutoMigrate(
		&models.User{},
		&models.UserLogin{},
		&models.Register{},
		&models.Product{},
		&models.Category{},
	)
	if err != nil {
		log.Fatal("Error during AutoMigrate: ", err)
	}
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