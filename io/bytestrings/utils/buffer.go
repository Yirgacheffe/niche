package utils

import (
	"bytes"
	"io"
	"io/ioutil"
)

// Buffer demonstrates some tricks for initializing bytes
// Buffers
// These buffers implement an io.Reader interface
func Buffer(rawString string) *bytes.Buffer {

	rawBytes := []byte(rawString)

	var b = new(bytes.Buffer)
	b.Write(rawBytes)

	b = bytes.NewBuffer(rawBytes)
	b = bytes.NewBufferString(rawString)

	return b // Show multi ways to create bytes.Buffer

}

// ToString is an example of taking an io.Reader and consuming
// it all, then returning a string
func toString(r io.Reader) (string, error) {

	rawBytes, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}

	return string(rawBytes), nil

}
