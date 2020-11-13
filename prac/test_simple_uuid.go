package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

func uuid() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)

	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:]), nil
}

func sessionid() (string, error) {
	b := make([]byte, 32)

	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return "", err
	}

	fmt.Println(b)
	return base64.URLEncoding.EncodeToString(b), nil

}

func main() {
	id, err := uuid()
	if err != nil {
		panic(err)
	}

	fmt.Println(id)

	sessionid, err := sessionid()
	if err != nil {
		panic(err)
	}

	fmt.Println(sessionid)
}
