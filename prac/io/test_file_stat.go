package main

import (
	"fmt"
	"log"
	"os"
)

func printInfo(fi os.FileInfo) {
	fmt.Println("File name:", fi.Name())
	fmt.Println("File size(bytes):", fi.Size())
	fmt.Println("File mod. time:", fi.ModTime().Format("Jan 2 2016"))
	fmt.Println("Is file a dir:", fi.IsDir())
}

func main() {

	fi, err := os.Stat("info.txt")
	if err != nil {
		log.Fatal("Error:", err)
	}

	printInfo(fi)

	fmt.Println("----- new record-----")

	file, err := os.Open("info.txt")
	if err != nil {
		log.Fatal("Error:", err)
	}

	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatal("Error:", err)
	} else {
		printInfo(fileInfo)
	}

	fmt.Println("----- new record-----")

	fmode := fileInfo.Mode()

	fmt.Printf("File Permission %o (octal), %b (binary)\n", fmode.Perm(), fmode.Perm())
	fmt.Println("File mode:", fmode.String())

}
