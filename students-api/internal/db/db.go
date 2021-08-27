package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"log"
	"os"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)


func InitPostgresDB() (*gorm.DB, error) {

	fmt.Println("Connection to DB")

	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	dbTable := os.Getenv("DB_TABLE")

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host,
		port,
		username,
		dbTable,
		password,
	)

	db, err := gorm.Open(postgres.Open(connStr))
	if err != nil {
		return db, err
	}

	return db, nil

}

func InitDatabase() (*gorm.DB, error) {

	log.Println("Connect to sqlite3")
	table := os.Getenv("DB_TABLE")

	if len(table) == 0 {
		table = "students.db"
	}

	db, err := gorm.Open(sqlite.Open(table), &gorm.Config{})
	if err != nil {
		return db, err
	}

	sql, err := db.DB()
	if err != nil {
		return db, err
	}

	sql.SetMaxIdleConns(5)
	sql.SetMaxOpenConns(5)
	sql.SetConnMaxIdleTime(time.Minute * 30)

	return db, nil
}
