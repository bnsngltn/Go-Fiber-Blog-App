package database

import (
	"fmt"
	"log"

	"gihub.com/bnsngltn/go-fiber-blog-app/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//DB pointer
var DB *gorm.DB

//Connect funct
func Connect() error {
	log.Println("Attempting to connect to the database...")

	var err error

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", config.DBHost, config.DBUser, config.DBPassword, config.DBName, config.DBPort)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	}

	CreateTables()

	return nil
}
