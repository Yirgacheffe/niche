package main

import (
	"fmt"
	"regexp"
)

func main() {

	src := []byte(`
		call hello alise
		hello bob
		call hello evins
	`)

	pat := regexp.MustCompile(`(?m)(call)\s+(?P<cmd>\w+)\s+(?P<arg>.+)\s*$`)
	var res []byte

	for _, s := range pat.FindAllSubmatchIndex(src, -1) {
		res = pat.Expand(res, []byte("$cmd('$arg')\n"), src, s)
	}

	fmt.Println(string(res))

	var zslice []byte
	aslice := []byte{}

}
