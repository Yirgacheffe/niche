package main

import (
	"io"
	"os"
)

func PipeExample() error {

	r, w := io.Pipe()

	go func() {
		// Could be encode json
		// base64 encode
		w.Write([]byte("testn"))
		w.Close()
	}()

	if _, err := io.Copy(os.Stdout, r); err != nil {
		return err
	}

	return nil // no errors, nouce comment :-(

}
