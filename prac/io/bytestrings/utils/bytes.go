package utils

import (
	"bufio"
	"bytes"
	"fmt"
)

// WorkWithBuffer will make use of the buffer created by the
// Buffer function
func WorkWithBuffer() error {

	rawString := "It's easy to encode unicode into a byte array"
	b := Buffer(rawString)

	// convert back to bytes easily
	b.Bytes()
	fmt.Println(b.String())

	s, err := toString(b)
	if err != nil {
		return err
	}

	fmt.Println(s)

	// we can also take our bytes and create a bytes reader
	// these readers implement io.Reader, io.ReaderAt,
	// io.WriterTo, io.Seeker, io.ByteScanner, and io.RuneScanner interfaces
	reader := bytes.NewReader([]byte(s))

	// we can also plug it into a scanner that allows
	// buffered reading and tokenzation
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	return nil

}
