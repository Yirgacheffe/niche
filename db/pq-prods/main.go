package main

import "os"

func main() {
	a := &App{}
	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"),
	)

	a.Run(":8020")
}

func ensureTableExist() {

}

func clearTable() {

}

const tableCreationQuery = `CREATE TABLE IF NOT EXIST products ()`
