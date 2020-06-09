package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	log "github.com/sirupsen/logrus"
)

// User - Represent user object in json
type User struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

// AppError - Application Error for logging
type AppError struct {
	HTTPStatus int    `json:"status"`
	Error      string `json:"error"`
	Message    string `json:"message"`
}

// ErrorResource - Wrap up the AppError
type ErrorResource struct {
	Data AppError `json:"data"`
}

// LoginResponse - Also the authentication response
type LoginResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

// LoginHandler - Do user login
func LoginHandler(w http.ResponseWriter, r *http.Request) {

	var user User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		DisplayAppError(
			w,
			err,
			"Invalid login data.",
			500,
		)
		return
	}

	if user.UserName != "xyz1234" && user.Password != "1qaz2wsx" {
		DisplayAppError(
			w,
			err,
			"Invalid login credentials.",
			401,
		)
		return
	}

	jwt, err := GenerateJWT(user.UserName, "member")
	if err != nil {
		DisplayAppError(
			w,
			err,
			"Error while generating access token.",
			500,
		)
	}

	loginResp := LoginResponse{Message: "Authenticated!", Token: jwt}
	j, err := json.Marshal(loginResp)
	if err != nil {
		DisplayAppError(
			w,
			err,
			"An unexpected error has happened, try again later.",
			500,
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)

}

// JwksHandler - Provide JWK url for token verification
func JwksHandler(w http.ResponseWriter, r *http.Request) {

}

// HealthCheckHandler - monitor purpose
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	_, err := w.Write([]byte("{\"alive\": true}"))
	if err != nil {
		log.Error(err)
	}

}

// DisplayAppError - Send back the AppError
func DisplayAppError(w http.ResponseWriter, handlerError error, message string, code int) {

	errObj := AppError{
		Error:      handlerError.Error(),
		HTTPStatus: code,
		Message:    message,
	}

	log.Fatalf("[AppError]: %s\n", handlerError)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	if j, err := json.Marshal(ErrorResource{Data: errObj}); err != nil {
		w.Write(j)
	}

}

func main() {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, AllowCredentials: true, AllowedMethods: []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
	})

	router := mux.NewRouter()
	handler := c.Handler(router)

	router.HandleFunc("/login", LoginHandler).Methods("POST")
	router.HandleFunc("/jwks", JwksHandler).Methods("GET")
	router.HandleFunc("/health", HealthCheckHandler).Methods("GET", "OPTIONS")

	log.Info(http.ListenAndServe(":8081", handler))
}
