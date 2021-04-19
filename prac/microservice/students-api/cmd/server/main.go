package main

import (
	"fmt"
	"net/http"
	internalHttp "students-api/internal/http"
)

func Run() error {
	fmt.Println("Running App")
	handler := internalHttp.NewHandler()
	handler.InitRoutes()
	if err := http.ListenAndServe(":9000", handler.Router); err != nil {
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
