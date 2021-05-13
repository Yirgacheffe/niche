package main

import (
	"fmt"
    "math/rand"
    "time"
)

var (
    letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
)

func GenRandomString(n int) (string, error) {

    if n >= 10 {
        return "", fmt.Errorf("Over size")
    }

    rand.Seed(time.Now().UnixNano())

    b := make([]rune, n)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }

    return string(b), nil
}

func main() {
    s, err := GenRandomString(8)
    if err != nil {
        panic(err)
    }

    fmt.Println(s)
}
