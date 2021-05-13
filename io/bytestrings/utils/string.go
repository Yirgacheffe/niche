package utils

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// SearchString shows a number of methods
// for searching a string
func SearchString() {
	s := "this is a test"

	fmt.Println(strings.Contains(s, "this"))
	fmt.Println(strings.ContainsAny(s, "abc"))
	fmt.Println(strings.HasPrefix(s, "this"))
	fmt.Println(strings.HasSuffix(s, "test"))
}

// ModifyString modifies a string in a number of ways
func ModifyString() {
	s := "simple string"
	fmt.Println(strings.Split(s, " "))
	fmt.Println(strings.Title(s))

	s = "   simgle string "
	fmt.Println(strings.TrimSpace(s))
}

// StringReader demonstrates how to create
// an io.Reader interface quickly with a string
func StringReader() {
	s := "simple string"
	r := strings.NewReader(s)

	io.Copy(os.Stdout, r) // Print s to stdout
}

func ReadStringOutToStdErr() {

	// test 1
	r := strings.NewReader("is test ok for yea?")
	fmt.Println("r length:", r.Len())

	b := make([]byte, 1)
	for {
		n, err := r.Read(b)
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Printf("Read %s bytes: %d\n", b, n)
	}

	// test 2
	s := strings.NewReader("This is an error!\n")
	fmt.Println("r length:", s.Len())

	n, err := s.WriteTo(os.Stderr)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Wrote %d bytes to os.Stderr\n", n)

}
