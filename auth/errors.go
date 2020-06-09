package main

import (
	"encoding/json"
	"fmt"
)

// HTTPError - represent what happened while we process the http request
type HTTPError struct {
	Code    int    `json:"-"`
	Message string `json:"message"`
	Cause   error  `json:"-"`
}

type ErrorResource struct {
	Errors HTTPError `json:"errors"`
}

func NewHTTPError(code int, message string, err error) *HTTPError {
	return &HTTPError{
		code,
		message,
		err,
	}
}

func (e *HTTPError) Error() string {
	if e.Cause == nil {
		return e.Message
	}
	return e.Message + ":" + e.Cause.Error()
}

func (e *HTTPError) SetCause(err error) *HTTPError {
	e.Cause = err
	return e
}

// ResponseBody returns JSON response body.
func (e *HTTPError) ResponseBody() ([]byte, error) {
	body, err := json.Marshal(e)
	if err != nil {
		return nil, fmt.Errorf("Error while parsing response body: %v", err)
	}
	return body, nil
}

// ResponseHeaders returns http status code and headers.
func (e *HTTPError) ResponseHeaders() (int, map[string]string) {
	return e.Code, map[string]string{
		"Content-Type": "application/json; charset=utf-8",
	}
}

// ------------------------------------------------------------~~~~~~
