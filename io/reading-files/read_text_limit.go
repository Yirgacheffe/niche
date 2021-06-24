package main

import (
	"fmt"
	"io"
)

// SampleStringData to hold a string data and a index
type SampleStringData struct {
	rawText string
	readIdx int // default is 0
}

func (s *SampleStringData) Read(p []byte) (n int, err error) {
	bytes := []byte(s.rawText)
	maxLen := len(bytes)

	// check current readIdx
	if s.readIdx >= maxLen {
		return 0, io.EOF
	}

	// check max length again
	nextLimit := s.readIdx + len(p)

	if nextLimit >= maxLen {
		nextLimit = maxLen
		err = io.EOF
	}

	// the actual bytes need to read into 'p'
	readBytes := bytes[s.readIdx:nextLimit]
	n = len(readBytes)

	copy(p, readBytes)
	s.readIdx = nextLimit

	return n, err // Re, return length of bytes and error
}

func main() {
	text := SampleStringData{
		rawText: "Hello, Again! Amazing world is wonderful!",
	}

	p := make([]byte, 3)

	// read until a error returned
	for {
		n, err := text.Read(p)
		fmt.Printf("%d bytes read, data: %s\n", n, p[:n])

		if err == io.EOF {
			fmt.Println("-------end-----of-----file-------")
			break
		} else if err != nil {
			fmt.Println("Huh!! Some error occured!", err)
			break
		}
	}
}
