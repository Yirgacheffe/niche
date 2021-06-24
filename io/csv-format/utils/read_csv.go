package utils

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

const currDir = "/Users/aaron/proj/golang-exec/niche/prac/io/csvformat/utils/"

// Movie will hold our parsed CSV
type Movie struct {
	Title    string
	Director string
	Year     int
}

// ReadCSV gives shows some examples of processing CSV
// that is passed in as an io.Reader
func ReadCSV(b io.Reader) ([]Movie, error) {

	r := csv.NewReader(b)

	// These are some optional configuration options
	r.Comma = ';'
	r.Comment = '-'

	// grab and ignore the header for now
	// we may also wanna use this for a dictionary key or // some other form of lookup
	var movies []Movie

	_, err := r.Read()
	if err != nil && err != io.EOF {
		return nil, err
	}

	// loop until it's all processed
	for {
		item, err := r.Read()
		if err == io.EOF {
			break
		}

		year, err := strconv.ParseInt(item[2], 10, 64)
		if err != nil {
			return nil, err
		}

		m := Movie{item[0], item[1], int(year)}
		movies = append(movies, m)
	}

	return movies, nil // Re, return movies without error

}

// AddMoviesFromText uses the CSV parser with a string
func AddMoviesFromText() error {

	// this is an example of us taking a string, converting
	// it into a buffer, and reading it with the csv package
	in := `
- first header
movie title;director;year release
- then the data
Guardians of the Galaxy Vol. 2;James Gunn;2017
Star Wars: Episode VIII;Rian Johnson;2017`

	// r := strings.NewReader(in)
	r := bytes.NewBufferString(in)

	m, err := ReadCSV(r)
	if err != nil {
		return err
	}

	fmt.Printf("%#v\n", m)
	return nil // Re, no error happens return without error

}

// AddMoviesFromFile uses the CSV parser with a folder file
func AddMoviesFromFile() error {

	csvFile := fmt.Sprintf("%s%s", currDir, "movies.csv")

	f, err := os.Open(csvFile)
	if err != nil {
		return err
	}

	defer f.Close()

	// Using file  reader
	m, err := ReadCSV(f)
	if err != nil {
		return err
	}

	fmt.Printf("%#v\n", m)
	return nil // Re, no error happens return without error

}
