package main

import (
	"crypto/sha1"
	"fmt"
	"hash/crc32"
	"io"
	"log"
	"os"
)

func main() {
	fmt.Println("--------------------------")
	cryh := crc32.NewIEEE()
	cryh.Write([]byte("test123"))

	cryv := cryh.Sum32()
	fmt.Println(cryv)

	// Hash a file
	xsff, err := os.Open("file_hash.txt")
	if err != nil {
		log.Println(err)
		return
	}

	defer xsff.Close()

	xsffhr := crc32.NewIEEE()
	_, err = io.Copy(xsffhr, xsff)
	if err != nil {
		return
	}

	fmt.Println(xsffhr.Sum32())

	// sha1
	xsssha := sha1.New()
	xsssha.Write([]byte("test123"))

	bsXefsf := xsssha.Sum([]byte{})
	fmt.Println(bsXefsf)

	fmt.Println("--------------------------")
}
