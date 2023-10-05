package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

func Connect() {
	var err error

	dsn := fmt.Sprintf("host=localhost user=postgres password=696462 dbname=postgres port=5432 sslmode=disable TimeZone=Africa/Lagos")
	Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully connected to the database")
	}
}
