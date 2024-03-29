package jwt

import (
	"fmt"
	"io/ioutil"
	"log"
	"testing"

	"github.com/golang-jwt/jwt"
)

var publicSecret []byte

func init() {
	if keyData, err := ioutil.ReadFile("keys/id_rsa.pub"); err != nil {
		log.Fatal(err)
	} else {
		publicSecret = keyData
	}
}

func keyFunc(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	}

	if key, err := jwt.ParseRSAPublicKeyFromPEM(publicSecret); err == nil {
		return key, nil
	} else {
		return nil, fmt.Errorf("KeyFunc: parse public key: %w", err)
	}
}

func Test_GenerateJWT(t *testing.T) {

	id := 1
	userName := "wang xiao"
	email := "wei@abcc.com"
	tokenString, err := GenerateJWT(id, userName, email)

	if err != nil {
		t.Errorf("expected nil error, received %s", err.Error())
	}

	token, err := jwt.Parse(tokenString, keyFunc)
	if err != nil {
		t.Errorf("expected nil error, received %s", err.Error())
	}

	if !token.Valid {
		t.Errorf("token invalid: %v", token)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		t.Errorf("claims invalid, not map claims: %v", claims)
	}

	actual := claims["email"]
	if actual != email {
		t.Errorf("wrong claim 'sub' value: got %v want %v", actual, email)
	}

}
