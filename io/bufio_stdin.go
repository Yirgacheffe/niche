package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
)

func main() {
	reader := bufio.NewReader(os.Stdin) // bufio.NewReaderSize(io.Reader, size)

	fmt.Println("What is your name?")
	text, _ := reader.ReadString('\n')

	fmt.Printf("Hello %v", text)

	// Scanner get input
	fmt.Println("Keep writing pls:")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println("Scanned from input: ", scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "read input:", err)
	}

	fmt.Printf("We are running on %v, %v", runtime.Version(), runtime.GOOS)
}
