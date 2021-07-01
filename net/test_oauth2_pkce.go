package main

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func pkceVerifier(length int) string {

	if length > 128 {
		length = 128
	}

	if length < 43 {
		length = 43
	}

	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-._~"
	l := len(charset)

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(l)]
	}

	return string(b)

}

func pkceChallenge(verifier string) string {
	sum := sha256.Sum256([]byte(verifier))
	challenge := base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(sum[:])
	return challenge
}

func main() {

	s := pkceVerifier(64)
	c := pkceChallenge(s)

	fmt.Println(c)

}
