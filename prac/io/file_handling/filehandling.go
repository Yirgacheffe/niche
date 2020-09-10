package main

import (
	"flag"
	"fmt"
	"io/ioutil"
)

func main() {

	fptr := flag.String("fpath", "data.txt", "target file path")
	flag.Parse()

	fmt.Println("Value of fpath is: ", *fptr)

	data, err := ioutil.ReadFile(*fptr)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Read file success: ", string(data))
}
