package main

import (
	"bufio"
	"fmt"
)

type WordCounter int
type LineCounter int

type scanFunc func(p []byte, EOF bool) (advance int, token []byte, err error)

func scanBytes(p []byte, fn scanFunc) (cnt int) {
	for true {
		advance, token, _ := fn(p, true)
		if len(token) == 0 {
			break
		}
		p = p[advance:]
		cnt++
	}
	return cnt
}

// WordCounter
func (c *WordCounter) Write(p []byte) (int, error) {
	cnt := scanBytes(p, bufio.ScanWords)
	*c += WordCounter(cnt)
	return cnt, nil
}

func (c WordCounter) String() string {
	return fmt.Sprintf("contains %d words", c)
}

// LineCounter
func (c *LineCounter) Write(p []byte) (int, error) {
	cnt := scanBytes(p, bufio.ScanLines)
	*c += LineCounter(cnt)
	return cnt, nil
}

func (c LineCounter) String() string {
	return fmt.Sprintf("contains %d lines", c)
}

func main() {
	var c WordCounter
	fmt.Println(c)

	fmt.Fprintf(&c, "This is an sentence.")
	fmt.Println(c)

	c = 0
	fmt.Fprintf(&c, "This")
	fmt.Println(c)

	var l LineCounter
	fmt.Println(l)

	fmt.Fprintf(&l, `This is another line`)
	fmt.Println(l)

	l = 0
	fmt.Fprintf(&l, "This is another\nline")
	fmt.Println(l)

	fmt.Fprintf(&l, "This is one line")
	fmt.Println(l)
}
