package main

import (
	"bufio"
	"fmt"
	"io"
)

func CountLines(r io.Reader) (int, error) {
	var (
		br    = bufio.NewReader(r)
		lines int
		err   error
	)

	for {
		_, err = br.ReadString('\n')
		lines++
		if err != nil {
			break
		}
	}

	if err != io.EOF {
		return 0, err
	}
	return lines, nil
}

func CountLinesBetter(r io.Reader) (int, error) {
	sc := bufio.NewScanner(r)
	lines := 0

	for sc.Scan() {
		lines++
	}
	return lines, sc.Err()
}

// ----------------# ------------------------ #----
type header struct {
	key, value string
}

type status struct {
	code   int
	reason string
}

// not good enough error handling
func writeResponse(w io.Writer, st status, headers []header, body io.Reader) error {
	_, err := fmt.Fprintf(w, "HTTP/1.1 %d %s\n", st.code, st.reason)
	if err != nil {
		return err
	}

	for _, h := range headers {
		_, err := fmt.Fprintf(w, "%s: %s\n", h.key, h.value)
		if err != nil {
			return err
		}
	}

	if _, err := fmt.Fprintf(w, "\n"); err != nil {
		return err
	}

	_, err = io.Copy(w, body)
	return err
}

// better way for error handling
type errWriter struct {
	io.Writer
	err error
}

func (e *errWriter) Write(buf []byte) (int, error) {
	if e.err != nil {
		return 0, e.err
	}

	var n int
	n, e.err = e.Writer.Write(buf)
	return n, nil
}

func anotherResponse(w io.Writer, st status, headers []header, body io.Reader) error {
	ew := &errWriter{Writer: w}
	fmt.Fprintf(ew, "HTTP/1.1 %d %s\n", st.code, st.reason)

	for _, h := range headers {
		fmt.Fprintf(ew, "%s: %s\n", h.key, h.value)
	}

	fmt.Fprintf(ew, "\n")

	io.Copy(ew, body)
	return ew.err
}

// ----------------# ------------------------ #----
