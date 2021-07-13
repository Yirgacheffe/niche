package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

type Order struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	PrepTime int    `json:"prepTime"` // preparation time in seconds
}

type Courier struct {
	Id   string
	Name string
}

func orderGenerator(done <-chan bool, orders []Order) <-chan Order {
	input := make(chan Order)

	go func() {
		defer close(input)

		workGen := time.Tick(500 * time.Millisecond)
		for _, o := range orders {
			select {
			case <-done:
				return
			case <-workGen:
				input <- o
			}
		}
	}()

	return input
}

func main() {

	data, err := ioutil.ReadFile("./data/input.json")
	if err != nil {
		log.Fatal("Open  input file error: ", err)
	}

	var orders []Order
	err = json.Unmarshal(data, &orders)
	if err != nil {
		log.Fatal("Parse input file error: ", err)
	}

	done := make(chan bool)
	defer close(done)

	generator := orderGenerator(done, orders)

	for v := range generator {
		fmt.Println(v)
	}

}
