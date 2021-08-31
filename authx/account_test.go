package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
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

var db *sql.DB

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
		db.Exec("INSERT INTO account(username, password, email) VALUES($1, $2, $3)", "user"+strconv.Itoa(i), "pwd", "email")
	}
}

func TestMain(m *testing.M) {
	conn, err := sql.Open("sqlite3", "./account.db")
	if err != nil {
		panic(err)
	}

	db = conn
	ensureTableExists()
	exitCode := m.Run()

	clearTable()
	os.Exit(exitCode)
}

func Test_AddItem(t *testing.T) {

	clearTable()
	addItems(1)

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
}
