package main

import (
	"database/sql"
	"log"
	"strconv"
	"testing"

	_ "github.com/mattn/go-sqlite3"
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

func ensureTableExists() {
	if _, err := db.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	log.Println("----------------------- Clear table -------------------------")
	db.Exec("DELETE FROM account")
	db.Exec("DELETE FROM sqlite_sequence where name=account")
}

func addItems(count int) {
	if count < 1 {
		count = 1
	}

	for i := 0; i < count; i++ {
		idx := strconv.Itoa(i)
		db.Exec(insertStmt, "user"+idx, "pwd"+idx, "email"+idx)
	}
}

var db *sql.DB

func setup() {
	conn, err := sql.Open("sqlite3", "./account.db")
	if err != nil {
		panic(err)
	}

	db = conn
	ensureTableExists()
	clearTable()
}

func Test_GetAccount(t *testing.T) {
	setup()
	addItems(1)

	repo := NewAccountRepo(db)
	account, err := repo.GetAccount("user0", "pwd0")
	if err != nil {
		t.Error("Get account failed", err)
	}

	actual := account.Email
	expect := "email0"

	if actual != expect {
		t.Errorf("GetAccount failed: got %v, want %v", actual, expect)
	}
}

/*
	rows, err := db.Query("select * from account")
	if err != nil {
		panic(err)
	}

	for rows.Next() {

		var (
			id int
			u  string
			p  string
			e  string
		)

		rows.Scan(&id, &u, &p, &e)
		fmt.Printf("%d - %s - %s - %s", id, u, p, e)
	}
*/
