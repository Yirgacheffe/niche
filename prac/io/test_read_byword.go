package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
)

func wordByWord(file string) error {
	var err error

	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("error reading file %s", err)
			return err
		}

		// By word ...
		r := regexp.MustCompile("[^\\s]+")
		words := r.FindAllString(line, -1)

		for _, word := range words {
			fmt.Println(word)
		}
		//-------------------------------------
	}

	return nil

}

func main() {

	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Printf("usage: test_read_byword <file1> [<file2> ...]\n")
		return
	}

	for _, file := range flag.Args() {
		err := wordByWord(file)
		if err != nil {
			fmt.Println(err)
		}
	}

	// flag.Args compare the os.Args... ???

}
