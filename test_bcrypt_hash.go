package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func hashIt(s string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	return string(hashed), err
}

func checkId(password, hashed string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	return err == nil
}

func main() {

	password := "very_secret"
	hashed, _ := hashIt(password)

	fmt.Println("Password:", password)
	fmt.Println("Hashed:  ", hashed)

	isMatch := checkId("secret", hashed)
	fmt.Println("Is match:", isMatch)

	matchAgain := checkId(password, hashed)
	fmt.Println("Is match:", matchAgain)

}
