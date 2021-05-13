package utils

import (
	"bytes"
	"encoding/csv"
	"io"
	"os"
)

// Book has an Author and Title
type Book struct {
	Author string
	Title  string
}

// Books is list of book
type Books []Book

// ToCSV takes a set of Books and writes to an io.Writer
// it returns any errors
func (books *Books) ToCSV(w io.Writer) error {

	n := csv.NewWriter(w)

	err := n.Write([]string{"Author", "Title"})
	if err != nil {
		return err
	}

	for _, b := range *books {
		err = n.Write([]string{b.Author, b.Title})
		if err != nil {
			return err
		}
	}

	n.Flush()
	return n.Error() // Return error if occured when flush

}

// WriteCSVOutput initializes a set of books
// and writes the to os.Stdout
func WriteCSVOutput() error {

	b := Books{
		Book{Author: "F Scott Fitzgerald", Title: "The Great Gatsby"},
		Book{Author: "J D Salinger", Title: "The Catcher in the Rye"},
	}

	return b.ToCSV(os.Stdout)

}

// WriteCSVBuffer returns a buffer csv for a set of books
func WriteCSVBuffer() (*bytes.Buffer, error) {

	b := Books{
		Book{Author: "F Scott Fitzgerald", Title: "The Great Gatsby"},
		Book{Author: "J D Salinger", Title: "The Catcher in the Rye"},
	}

	w := &bytes.Buffer{}
	return w, b.ToCSV(w)

}
