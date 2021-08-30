package main

import (
	"time"

	"github.com/golang-jwt/jwt"
)

const (
	appId  = "1873974"
	jwtIss = "nichesoft.io"
	jwtAud = appId
	jwtKid = "DHFbpoIUqrY8t2zpA2qXfCmr5VO5ZEr4RzHU_-envvQ"
)

var key = []byte("sdkfjsdkksdfjafiemr3434jk")

// GenerateJWT - Generate jwt token
func GenerateJWT(name string, roles []string) (string, error) {
	claims := jwt.MapClaims{
		"iss": jwtIss,
		"aud": jwtAud,
		"sub": name, // TODO: change to user id and may add email
		"exp": time.Now().Add(time.Minute * 30).Unix(),
		"kid": jwtKid,
	}

	// create a signer for rsa 256
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := t.SignedString(key)
	if err != nil {
		return "", err
	}

	return tokenString, nil // jwt string witout error ......
}
