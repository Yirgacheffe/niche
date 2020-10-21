package utils

import (
	"errors"
	"io"
	"os"
)

var demoDir = "example_dir"

// Operate manipulates files and directories
func Operate() error {

	// will create a directory
	// `./example_dir`
	if err := os.Mkdir(demoDir, os.FileMode(0755)); err != nil {
		return err
	}

	// cd to `.`
	if err := os.Chdir(demoDir); err != nil {
		return err
	}

	// f is a generic file object,
	// it also implements multiple interfaces
	// and can be used as a reader or writer
	// if the correct bits are set when opening
	f, err := os.Create("test.txt")
	if err != nil {
		return err
	}

	line := []byte("hello\n")

	cnt, err := f.Write(line)
	if err != nil {
		return err
	}

	if cnt != len(line) {
		return errors.New("incorrect length has been write into file")
	}

	if err := f.Close(); err != nil {
		return err
	}

	// read file
	f, err = os.Open("test.txt")
	if err != nil {
		return err
	}

	io.Copy(os.Stdout, f)
	if err := f.Close(); err != nil {
		return err
	}

	// back to `/tmp`
	if err := os.Chdir(".."); err != nil {
		return err
	}

	// cleanup, os.RemoveAll can be dangerous if you
	// point at the wrong directory, use user input, and especially if you run as root
	if err := os.RemoveAll(demoDir); err != nil {
		return err
	}

	return nil // Re. Re. Return without error

}
