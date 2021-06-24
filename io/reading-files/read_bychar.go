package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func charByChar(file string) error {
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

		// By char ...
		for _, x := range line {
			fmt.Print(string(x))
		}
		//-------------------------------------
	}

	return nil

}

func main() {

	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Printf("usage: test_read_bychar <file1> [<file2> ...]\n")
		return
	}

	for _, file := range flag.Args() {
		err := charByChar(file)
		if err != nil {
			fmt.Println(err)
		}
	}

	// flag.Args compare the os.Args... ???

}
