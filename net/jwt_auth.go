package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

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

// Token structure JWT
type Token struct {
	Token string `json:"token"`
}

// Response text
type Response struct {
	Text string `json:"text"`
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

	// Customize claims for login users
	claims := jwt.MapClaims{
		"iss": "admin",
		"nbf": time.Now().Unix(),
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Minute * 30).Unix(),
	}

	t := jwt.NewWithClaims(jwt.GetSigningMethod("RS256"), claims)

	tokenString, err := t.SignedString(signKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Sorry, error happened while Signing Token!")
		log.Printf("Token Signing error: %v\n", err)
		return
	}

	response := Token{tokenString}
	jsonResponse(response, w)

}

func jsonResponse(response interface{}, w http.ResponseWriter) {

	json, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)

}

func authHandler(w http.ResponseWriter, r *http.Request) {

	token, err := jwt.ParseFromRequest(r, func(token *jwt.Token) (interface{}, error) {
		return verifyKey, nil
	})

	if err != nil {
		switch err.(type) {
		case *jwt.ValidationError:
			vErr := err.(*jwt.ValidationError)
			switch vErr.Errors {
			case jwt.ValidationErrorExpired:
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintln(w, "Token Expired, get a new one.")
				return
			default:
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintln(w, "Error while Parse Token!")
				log.Printf("ValidationError error: %+v\n", vErr.Errors)
				return
			}
		default:
			w.WriterHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, "Error while Parsing Token!")
			log.Printf("Token parse error: %v\n", err)
			return
		}
	}

	var response Response

	if token.Valid {
		response = Response{"Authorized to the system"}
	} else {
		response = Response{"Invalid token"}
	}

	jsonResponse(response, w)

}

func authMiddleware(w http.ResponseWriter, r *http.Request, next http.HandleFunc) {

	token, err := jwt.ParseFromRequest(r, func(token *jwt.Token) (interface{}, error) {
		return verifyKey, nil
	})

	if err == nil && token.Valid {
		next(w, r)
	} else {
		w.WriterHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Authentication failed.")
	}

}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/login", loginHandler).Methods("POST")
	r.HandleFunc("/auth", authHandler).Methods("POST")

	r.Handle("/admin",
		negroni.New(negroni.HandleFunc(authMiddleware), negroni.Wrap(http.HandleFunc(adminHandler))))

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	log.Println("Listening...")
	server.ListenAndServe()

}
