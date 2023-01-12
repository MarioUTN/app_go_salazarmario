package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB
var err error

func DatabaseConnection() (*gorm.DB, error) {
	host := "localhost"
	port := "5432"
	dbName := "app_golang"
	dbUser := "postgres"
	password := "password-postgresql"
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host,
		port,
		dbUser,
		dbName,
		password,
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the database...", err)
	} else {
		fmt.Println("Database connection successful...")
	}
	return DB, err
}
