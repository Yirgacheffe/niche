package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type appError struct {
	HTTPStatus int    `json:"status"`
	Error      string `json:"error"`
	Message    string `json:"message"`
}

type errorResource struct {
	Data appError `json:"data"`
}

func DisplayAppError(w http.ResponseWriter, handlerError error, message string, code int) {

	errObj := appError{
		Error:      handlerError.Error(),
		Message:    message,
		HTTPStatus: code,
	}

	log.Fatalf("[AppError]: %s\n", handlerError)

	w.Header().Set("Content-Type", "application/json; charset=utf8")
	w.WriteHeader(code)
	if j, err := json.Marshal(errorResource{Data: errObj}); err == nil {
		w.Write(j)
	}

}

type AppConfig struct {
	Server, DBHostURL, DBUser, DBPassword, Database string
}

func initConfig() {

}

func loadAPPConfig() {

	file, err := os.Open("common/config.json")
	defer file.Close()
	if err != nil {
		log.Fatalf("[LoadAppConfig]: %s\n", err)
	}

	decoder := json.NewDecoder(file)
	appConfig := AppConfig{}
	err = decoder.Decode(&appConfig)
	if err != nil {
		log.Fatalf("[LoadAppConfig]: %s\n", err)
	}
	
}
