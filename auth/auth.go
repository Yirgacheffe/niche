package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"fortio.org/fortio/log"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

const (
	privKeyPath = "keys/niche.rsa"
	pubKeyPath  = "keys/niche.rsa.pub"
)

var signKey, verifyKey []byte

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

	//
	if user.UserName != "xyz1234" && user.Password != "1qwe4edc" {
		w.WriteHeader(http.StatusForbidden)
		fmt.Println(w, "Wrong credentials.")
		return
	}

}

func authHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, AllowCredentials: true, AllowedMethods: []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
	})

	router := mux.NewRouter()
	router.HandleFunc("/login", loginHandler).Methods("POST")

	handler := c.Handler(router)
	log.Fatal(http.ListenAndServe(":80801", handler))

}
