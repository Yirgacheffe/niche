package main

import (
	"crypto/rsa"
	"io/ioutil"
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// using asymmetric crypto/RSA keys, openssl genrsa
const (
	privKeyPath = "keys/niche.rsa"
	pubKeyPath  = "keys/niche.rsa.pub"
)

var (
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
)

// User structure for parsing login credentials
type User struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

func InitKeys() {
	var err error

	signBytes, err := ioutil.ReadFile(privKeyPath)
	if err != nil {
		log.Fatalf("[initKeys]: %s\n", err)
		panic(err)
	}

	verifyBytes, err := ioutil.ReadFile(pubKeyPath)
	if err != nil {
		log.Fatalf("[initKeys]: %s\n", err)
		panic(err)
	}

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		log.Fatal("error")
	}

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		log.Fatal("error")
	}

}

func GenerateJWT(name, role string) (string, error) {

	claims := jwt.MapClaims{
		"iss":  "admin",
		"exp":  time.Now().Add(time.Minute * 20).Unix(),
		"name": name,
		"role": role,
	}

	// create a signer for rsa 256
	t := jwt.NewWithClaims(jwt.GetSigningMethod("RS256"), claims)

	tokenString, err := t.SignedString(signKey)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return tokenString, nil
}
