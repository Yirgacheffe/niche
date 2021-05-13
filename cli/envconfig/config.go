package main

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

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
