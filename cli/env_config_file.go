package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

// Config will hold the config which captured from a json file and env vars
type Config struct {
	Version string `json:"version" required:"true"`
	IsSafe  bool   `json:"is_safe" default:"true"`
	Secret  string `json:"secret"`
}

// LoadConfig will load files optionally from the json file
// stored at path
// then will override those values based on the envconfig struct tags.
// The envPrefix is how we prefix our environment variables.
func LoadConfig(path, envPrefix string, config interface{}) error {

	p := strings.Trim(path, " ")
	if p == "" {
		return errors.New("error: no config file path found")
	}

	err := LoadFile(p, config)
	if err != nil {
		return errors.Wrap(err, "error loading config file")
	}

	err = envconfig.Process(envPrefix, config)
	return errors.Wrap(err, "error load the config from env")

}

// LoadFile unmarshalls a json file into a config struct ......
func LoadFile(path string, config interface{}) error {

	f, err := os.Open(path)
	if err != nil {
		return errors.Wrap(err, "failed to read json config file.")
	}

	defer f.Close()

	decoder := json.NewDecoder(f)
	err = decoder.Decode(config)
	if err != nil {
		return errors.Wrap(err, "failed to decode config file...")
	}

	return nil // Re, return without error :-|

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
