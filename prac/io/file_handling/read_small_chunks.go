package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {

	fptr := flag.String("fpath", "data.txt", "target file path")
	flag.Parse()

	f, err := os.Open(*fptr)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()

	r := bufio.NewReader(f)
	buf := make([]byte, 3)

	for {
		n, err := r.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(string(buf[0:n]))
	}

}
