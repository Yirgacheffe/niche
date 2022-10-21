package handler

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"niche-auth/internal/model"
)

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS account
(
	id         INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	username   TEXT NOT NULL,
	password   TEXT NOT NULL,
	email      TEXT NOT NULL
)
`

const insertStmt = `INSERT INTO account(username, password, email) VALUES($1, $2, $3)`

var sqliteConn *sql.DB

func ensureTableExists() {
	if _, err := sqliteConn.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func setup() {
	conn, err := sql.Open("sqlite3", "./account.db")
	if err != nil {
		panic(err)
	}

	sqliteConn = conn
	ensureTableExists()
	clearTable()
}

func clearTable() {
	log.Println("----------------------- Clear table -------------------------")
	sqliteConn.Exec("DELETE FROM account")
	sqliteConn.Exec("DELETE FROM sqlite_sequence where name=account")
}

func addItems(count int) {
	if count < 1 {
		count = 1
	}

	for i := 0; i < count; i++ {
		idx := strconv.Itoa(i)
		sqliteConn.Exec(insertStmt, "user"+idx, "pwd"+idx, "email"+idx)
	}
}

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

	h := &AuthHandler{model.NewAccountRepo(sqliteConn)}
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

	h := &AuthHandler{model.NewAccountRepo(sqliteConn)}
	h.Login(rr, req)

	res := rr.Result()
	// respBody, _ := ioutil.ReadAll(res.Body)
	assert.Equal(t, http.StatusOK, res.StatusCode)

}
