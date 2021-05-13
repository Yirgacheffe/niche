package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	media "./models"
)

// WordCount takes a file and returns a map
// with each word
// as a key and it's number of appearances as a value
func WordCount(r io.Reader) map[string]int {
	result := make(map[string]int)

	// make scanner from a reader to count the word
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		result[scanner.Text()]++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}

	return result // Re, return as no error happened
}

func main() {

	// test package, seems failed ...
	fmt.Println("WIP......")

	m := media.Movie{}
	m.Title = "Last Scene"

	fmt.Println(m.Title)

	// test word counter from input
	fmt.Printf("string: #nbr_of_occurrence\n")
	wordcount := WordCount(os.Stdin)

	for k, v := range wordcount {
		fmt.Printf("%s: %d\n", k, v)
	}

	// end of this demo, to print word count from input

}
