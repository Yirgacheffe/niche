package main

import (
	"io"
	"os"
)

// ----------------# ------------------------ #----
type Document struct {
	os.File
}

func Save(f os.File, doc *Document) error {
	return nil
}

func SaveBetter(w io.Writer, doc *Document) error {
	return nil
}

// anyPostive indicates if any value is greater than zero.
// the caller should pass at least one parameter
func anyPositive(first int, rest ...int) bool {
	if first > 0 {
		return true
	}

	for _, v := range rest {
		if v > 0 {
			return true
		}
	}

	return false
}

// ----------------# ------------------------ #----

func Copy(from, to string) error {
	return nil
}

// following is better then above -->
type Source struct {
	path string
}

func (s *Source) Copy(dest string) error {
	return nil
}

// ----------------# ------------------------ #----
