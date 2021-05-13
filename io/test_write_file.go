package main

import (
	"fmt"
	"os"
)

const tmpFileName = "test_lines.txt"

func writeString() {

	f, err := os.Create("test_string.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()

	// length
	length, err := f.WriteString("Hello world!")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(length, "bytes write successfully!")

}

func writeByte() {

	f, err := os.Create("test_byte.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()

	d2 := []byte{104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100}
	n2, err := f.Write(d2)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(n2, "bytes write successfully!")

}

func writeStringLineByLine() {

	f, err := os.Create(tmpFileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()

	lines := []string{"Dear President!", "I would like to share with you, my story.", "Best regards."}

	for _, v := range lines {
		fmt.Fprintln(f, v)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	fmt.Println("Line by line writes successfully!")
}

func writeStringWithAppendMode() {

	f, err := os.OpenFile(tmpFileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()

	newLine := "file append easily~~~"
	_, err = fmt.Fprintln(f, newLine)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("File appended successfully")

}

func main() {

	writeString()
	writeByte()
	writeStringLineByLine()
	writeStringWithAppendMode()

}
