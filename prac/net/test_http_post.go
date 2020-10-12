package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {

	body := strings.NewReader(`
		{
			"name": "Mike Pence",
			"age": 63,
			"Salary": 10,
		}
	`)

	res, err := http.Post(
		"http://dummy.restapiexample.com/api/v1/create",
		"applicaiton/json; charset=utf-8",
		body,
	)

	if err != nil {
		log.Fatal(err)
	}

	data, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()

	fmt.Println(res.Request.Header.Get("Content-Type"))
	fmt.Printf("%s\n", data)

}
