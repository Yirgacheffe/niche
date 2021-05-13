package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	log.Println(doGet())
}

func doGet() string {

	url := "http://localhost:8081/configs"
	client := &http.Client{Timeout: time.Second * 3}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("After request create...")

	response, err := client.Do(request)
	if err != nil {
		log.Println(err)
		return "Error"
	}

	defer response.Body.Close()

	return "No error"

}
