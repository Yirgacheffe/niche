package main

import (
	"bytes"
	"io"
	"log"
	"os"
)

type User struct {
	Name string
	Age  int
}

func main() {
	u := User{Name: "KYC", Age: 18}

	w1 := &bytes.Buffer{}
	w2 := os.Stdout
	w3, err := os.OpenFile("log.txt", os.O_WRONLY|os.O_CREATE, 0644)

	if err != nil {
		log.Fatalf("create file log.txt failed: %v", err)
	}

	logger := log.New(io.MultiWriter(w1, w2, w3), "", log.Lshortfile|log.LstdFlags)
	logger.Printf("%s login, age:%d", u.Name, u.Age)
}
