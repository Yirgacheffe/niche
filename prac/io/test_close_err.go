package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	f, err := os.OpenFile("info.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0744)
	if err != nil {
		log.Fatal("Error:", err)
	}

	n, err := io.WriteString(f, "Hello World!")
	if err != nil {
		log.Fatal("Error:", err)
	} else {
		fmt.Println("Successful! %d bytes has been written.\n", n)
	}

	f.Close()

	x, err := io.WriteString(f, "Hello Hello! Again~")
	if err != nil {
		log.Fatal("Error:", err)
	} else {
		fmt.Println("Successful! %d bytes has been written.\n", x)
	}

}
