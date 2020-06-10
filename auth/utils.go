package main

import (
	"encoding/json"
	"net/http"
)

// AppError - Application Error for logging
type AppError struct {
	HTTPStatus int    `json:"status"`
	Message    string `json:"message"`
}

// ErrorResponse - Wrap up the AppError
type ErrorResponse struct {
	Errors AppError `json:"error"`
}

// DisplayAppError - Send back the AppError
func DisplayAppError(w http.ResponseWriter, code int, message string) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)

	errObj := AppError{
		HTTPStatus: code, Message: message,
	}

	errResponse := ErrorResponse{
		errObj,
	}

	if j, err := json.Marshal(errResponse); err == nil {
		w.Write(j)
	}

}
