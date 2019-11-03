package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

const (
	privKeyPath = "keys/app.rsa"
	pubKeyPath  = "keys/app.rsa.pub"
)

var (
	verifyKey, signKey []byte
)

// User structure for parsing login credentials
type User struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

func init() {
	var err error

	signKey, err = ioutil.ReadFile(privKeyPath)
	if err != nil {
		log.Fatal("Error reading private key")
		return
	}

	verifyKey, err = ioutil.ReadFile(pubKeyPath)
	if err != nil {
		log.Fatal("Error reading public key")
		return
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {

	var user User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(w, "Request body error, can not parse parameters.")
		return
	}

	if user.UserName != "xyz1234" && user.Password != "1qaz2wsx" {
		w.WriteHeader(http.StatusForbidden)
		fmt.Println(w, "Wrong credentials.")
		return
	}

	t := jwt.New(jwt.GetSigningMethod("RS256"))

}
