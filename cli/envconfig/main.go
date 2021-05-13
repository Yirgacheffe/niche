package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Config will hold the config which captured from a json file and env vars
type Config struct {
	Version string `json:"version" required:"true"`
	IsSafe  bool   `json:"is_safe" default:"true"`
	Secret  string `json:"secret"`
}

func main() {

	var err error

	f, err := ioutil.TempFile("./", "tmp.json")
	if err != nil {
		panic(err)
	}

	defer f.Close()
	defer os.Remove(f.Name())

	// json string to hold our secret
	secrets := `{
		"secret": "very secret" 
	}`

	buffer := bytes.NewBufferString(secrets)

	_, err = f.Write(buffer.Bytes())
	if err != nil {
		panic(err)
	}

	// set environment variables if needed
	if err = os.Setenv("EXAMPLE_VERSION", "1.0"); err != nil {
		panic(err)
	}

	if err = os.Setenv("EXAMPLE_ISSAFE", "false"); err != nil {
		panic(err)
	}

	c := Config{}
	err = LoadConfig(f.Name(), "EXAMPLE", &c)
	if err != nil {
		panic(err)
	}

	fullpath, _ := filepath.Abs(f.Name())

	fmt.Println("secrets file fullpath:", fullpath)
	fmt.Println("secrets file content:", secrets)

	fmt.Println("EXAMPLE_VERSION =", os.Getenv("EXAMPLE_VERSION"))
	fmt.Println("EXAMPLE_ISSAFE  =", os.Getenv("EXAMPLE_ISSAFE"))

	fmt.Printf("Final config: %#v\n", c)

}
