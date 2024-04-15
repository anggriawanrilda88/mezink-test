package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	// postgres connection set
	dbHost := os.Getenv("DB_HOST")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbUsername := os.Getenv("DB_USERNAME")
	dbDbname := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	// Konfigurasi koneksi ke database
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, dbUsername, dbDbname, dbPassword)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		return nil, err
	}

	// show hide debug postgresql
	debugDb := os.Getenv("DB_DEBUG")
	if debugDb == "true" {
		db = db.Debug()
	}

	DB = db
	return db, nil
}
