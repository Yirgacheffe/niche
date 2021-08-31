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

func Test_AuthHandler_NotCorrectUser(t *testing.T) {

	data := url.Values{}
	data.Set("username", "123")
	data.Set("password", "456")

	req, err := http.NewRequest("POST", "/oauth/auth", strings.NewReader(data.Encode()))
	if err != nil {
		t.Error("Failed to make request", err)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	rr := httptest.NewRecorder()
	AuthHandler(rr, req)

	resp := rr.Result()
	respBody, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)

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
