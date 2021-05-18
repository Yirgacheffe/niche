package main

import (
	"fmt"
	"net/http"
	"students-api/internal/db"

	internalHttp "students-api/internal/http"
	"students-api/internal/service/student"
)

func Run() error {
	fmt.Println("Running App")

	conn, err := db.InitDatabase()
	if err != nil {
		return err
	}

	err = db.MigrateDB(conn)
	if err != nil {
		return err
	}

	service := student.NewService(conn)
	handler := internalHttp.NewHandler(service)
	handler.InitRoutes()

	if err := http.ListenAndServe(":9010", handler.Router); err != nil {
		return err
	}

	return nil
}

func main() {
	err := Run()
	if err != nil {
		fmt.Println("Error running app")
		fmt.Println(err)
	}
}
