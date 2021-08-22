package main

import (
	"os"
)

// export APP_DB_USERNAME=postgres
// export APP_DB_PASSWORD=
// export APP_DB_NAME=postgres

func main() {
	a := &App{}
	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))

	a.Run(":8020")
}
