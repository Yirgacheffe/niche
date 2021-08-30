package main

import (
	"fmt"
	"testing"

	"github.com/golang-jwt/jwt"
)

func keyFunc(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	}
	return []byte("sdkfjsdkksdfjafiemr3434jk"), nil
}

func Test_GenerateJWT(t *testing.T) {

	userName := "wang xiao"
	roles := []string{}
	tokenString, err := GenerateJWT(userName, roles)

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

	actual := claims["sub"]
	expect := "wang xiao"

	if actual != expect {
		t.Errorf("wrong claim 'sub' value: got %v want %v", actual, expect)
	}

}
