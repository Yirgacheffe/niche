package main

import (
	"encoding/base64"
	"fmt"
)

func base64Encode(src []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(src))
}

func base64Decode(src []byte) ([]byte, error) {
	return base64.StdEncoding.DecodeString(string(src))
}

func main() {

	hello := "Hello world!"
	debyte := base64Encode([]byte(hello))

	fmt.Println(debyte)

	endbyte, err := base64Decode(debyte)
	if err != nil {
		fmt.Println(err.Error())
	}

	if hello != string(endbyte) {
		fmt.Println("decode failed")
	}

	fmt.Println(string(endbyte))

}