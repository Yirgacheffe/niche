package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {

	var buf bytes.Buffer
	buf.Write([]byte("test"))

	// file testing
	file, err := os.Open("build.sh")
	if err != nil {
		return
	}

	defer file.Close()

	// the file size
	stat, err := file.Stat()
	if err != nil {
		return
	}

	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return
	}

	fmt.Println(string(bs))

	// Another way to read file
	bss, err := ioutil.ReadFile("build.sh")
	if err != nil {
		return
	}

	fmt.Println(string(bss))

	// another way is define a small size of []byte read file when encounter io.EOF

	// Create a file
	fileS, err := os.Create("test.txt")
	if err != nil {
		return
	}

	defer fileS.Close()
	fileS.WriteString("test it with create")

	// Directory operation
	dir, err := os.Open(".")
	if err != nil {
		return
	}

	defer dir.Close()

	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		return
	}

	for _, fi := range fileInfos {
		fmt.Println(fi.Name())
	}

	// file path walk
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})

}
