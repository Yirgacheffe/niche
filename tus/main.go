package main

import (
	"fmt"
	"log"
	"os"
)

var (
	dbHost, dbPort, dbUser, dbPazz, dbName string
)

func init() {

}

func main() {

	dbHost := os.Getenv("DB_HOST")
	if len(dbHost) == 0 {
		dbHost = "127.0.0.1"
	}

	dbPort := os.Getenv("DB_PORT")
	if len(dbPort) == 0 {
		dbPort = "3306"
	}

	dbUser := os.Getenv("DB_USER")
	if len(dbUser) == 0 {
		dbUser = "tollgate"
	}

	dbPazz := os.Getenv("DB_PASS")
	if len(dbPazz) == 0 {
		dbPazz = "zMUjpGW67o"
	}

	dbName := os.Getenv("DB_NAME")
	if len(dbName) == 0 {
		dbName = "TOLLGATE"
	}

	db, err := ConnectSQL(dbHost, dbPort, dbUser, dbPazz, dbName)
	if err != nil {
		log.Println(err)
		os.Exit(-1)
	}

	fileRepo := NewMySQLFileRepo(db.SQL)

	/*

		f := File{
			Offset:       101,
			UploadLength: 101,
			IsComplete:   "N",
		}

		id, err := fileRepo.Create(&f)
		if err != nil {
			log.Println("Create file ", err)

		}

	*/

	f, err := fileRepo.GetByID(1)
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("ID = %d\n", f.ID)
	fmt.Println(f)

}
