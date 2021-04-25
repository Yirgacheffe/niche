package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase() (*gorm.DB, error) {

	fmt.Println("Connection to DB")

	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	table := os.Getenv("DB_TABLE")

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host,
		port,
		username,
		table, password)

	db, err := gorm.Open(postgres.Open(connStr))
	if err != nil {
		return db, err
	}

	return db, nil

}

/*
func InitSqlite3DB() (*gorm.DB, error) {

	fmt.Println("Connect to sqlite3")
	table := os.Getenv("DB_TABLE")

	db, err := gorm.Open(sqlite.Open(table), &gorm.Config{})
	if err != nil {
		return db, err
	}

	return db, nil

}
*/
