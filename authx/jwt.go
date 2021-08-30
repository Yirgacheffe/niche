package main

import (
	"io/ioutil"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
)

const (
	appId  = "1873974"
	jwtIss = "nichesoft.io"
	jwtAud = appId
	jwtKid = "DHFbpoIUqrY8t2zpA2qXfCmr5VO5ZEr4RzHU_-envvQ"
)

var hmacSecret []byte

func init() {
	if keyData, e := ioutil.ReadFile("keys/hmac"); e != nil {
		log.Fatal(e)
	} else {
		hmacSecret = keyData
	}
}

// GenerateJWT - Generate jwt token
func GenerateJWT(id int, name, email string) (string, error) {
	claims := jwt.MapClaims{
		"iss":   jwtIss,
		"aud":   jwtAud,
		"kid":   jwtKid,
		"exp":   time.Now().Add(time.Minute * 30).Unix(),
		"sub":   id,
		"email": email,
		"name":  name,
	}

	// create a signer for rsa 256
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := t.SignedString(hmacSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil // jwt string witout error ......
}
