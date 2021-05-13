package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func lineByLine(s string) error {
	var err error

	f, err := os.Open(s)
	if err != nil {
		return err
	}

	defer f.Close()
	r := bufio.NewReader(f)

	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("error reading file %s\n", err)
			break
		}
		fmt.Print(line)
	}

	return nil
}

func main() {

	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Printf("usage: test_read_byline <file1> [<file2> ...]\n")
		return
	}

	for _, file := range flag.Args() {
		err := lineByLine(file)
		if err != nil {
			fmt.Println(err)
		}
	}

	// flag.Args compare the os.Args... ???

}
