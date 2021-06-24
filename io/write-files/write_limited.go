package main

import (
	"fmt"
	"io"
)

// SampleStore - sample store type
type SampleStore struct {
	data []byte
}

func (ss *SampleStore) Write(p []byte) (n int, err error) {

	if len(ss.data) == 10 {
		return 0, io.EOF // end with limit error if `10` bytes
	}

	remainCap := 10 - len(ss.data)
	writeLength := len(p)

	if remainCap <= writeLength {
		writeLength = remainCap
		err = io.EOF
	}

	ss.data = append(ss.data, p[:writeLength]...)
	return writeLength, err

}

func main() {
	ss := SampleStore{}

	bytes1, err1 := ss.Write([]byte("Hello!"))
	fmt.Printf("Bytes written %d, error: %v\n", bytes1, err1)
	fmt.Printf("Value of ss.data: %s\n\n", ss.data)

	bytes2, err2 := ss.Write([]byte(" Amazing"))
	fmt.Printf("Bytes written %d, error: %v\n", bytes2, err2)
	fmt.Printf("Value of ss.data: %s\n\n", ss.data)

	bytes3, err3 := ss.Write([]byte(" Amazing"))
	fmt.Printf("Bytes written %d, error: %v\n", bytes3, err3)
	fmt.Printf("Value of ss.data: %s\n\n", ss.data)

}
