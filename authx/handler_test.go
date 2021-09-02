package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AuthHandler_Login401(t *testing.T) {

	setup()
	addItems(1)

	data := url.Values{}
	data.Set("username", "123")
	data.Set("password", "456")

	req, err := http.NewRequest("POST", "/oauth/auth", strings.NewReader(data.Encode()))
	if err != nil {
		t.Error("Failed to make request", err)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	rr := httptest.NewRecorder()

	h := &AuthHandler{NewAccountRepo(db)}
	h.Login(rr, req)

	res := rr.Result()
	respBody, _ := ioutil.ReadAll(res.Body)

	assert.Equal(t, http.StatusUnauthorized, res.StatusCode)

	var v map[string]string
	if err = json.Unmarshal(respBody, &v); err != nil {
		t.Error("parse response body failed", err)
	}

	actual := v["code"]
	expect := "AUT001"

	if actual != expect {
		t.Errorf("response error code failed: got %v want %v", actual, expect)
	}

}

func Test_AuthHandler_Login200(t *testing.T) {

	setup()
	addItems(1)

	data := url.Values{}
	data.Set("username", "user0")
	data.Set("password", "pwd0")

	req, err := http.NewRequest("POST", "/oauth/auth", strings.NewReader(data.Encode()))
	if err != nil {
		t.Error("Failed to make request", err)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	rr := httptest.NewRecorder()

	h := &AuthHandler{NewAccountRepo(db)}
	h.Login(rr, req)

	res := rr.Result()
	respBody, _ := ioutil.ReadAll(res.Body)
	assert.Equal(t, http.StatusOK, res.StatusCode)

}
