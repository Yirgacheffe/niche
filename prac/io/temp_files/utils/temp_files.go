package utils

import (
	"fmt"
	"io/ioutil"
	"os"
)

// WorkWithTemp will give some basic patterns for working
// with temporary files and directories
func WorkWithTemp() error {

	// If the first arguments is ""
	// then it will use value of os.TempDir()
	// "./" will put temp folder with 'main.go' folder
	t, err := ioutil.TempDir("./", "tmp")
	if err != nil {
		return err
	}

	// This will delete everything inside the temp file when this function exits
	// if you want to do this later
	// be sure to return the directory name to the calling function
	defer os.RemoveAll(t)

	// the directory must exist to create the tempfile created.
	// t is an *os.File object.
	tf, err := ioutil.TempFile(t, "tmp-01.html")
	if err != nil {
		return err
	}

	// fmt.Println(filepath.Abs(tf.Name()))
	fmt.Println(tf.Name())

	// normally we'd delete the temporary file here, but
	// because we're placing it in a temp directory, it gets cleaned up by the earlier defer
	return nil

}
