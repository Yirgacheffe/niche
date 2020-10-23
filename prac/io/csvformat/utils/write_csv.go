package utils

import (
	"encoding/csv"
	"io"
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
func WriteCSVOutput() error {}

// WriteCSVBuffer returns a buffer csv for a set of books
func WriteCSVBuffer() error {}
