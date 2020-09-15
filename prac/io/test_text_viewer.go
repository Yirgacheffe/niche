package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	file, err := os.Open("customers.csv")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		items := strings.Split(line, ",")
		fmt.Printf("Name: %s %s Email: %s\n", items[1], items[2], items[3])
		fmt.Println("------")
	}

}
