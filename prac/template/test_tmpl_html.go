package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type Entry struct {
	Number int
	Double int
	Square int
}

var DATA []Entry
var tFile string

func dbFileHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf(
		"Host: %s Path: %s\n", r.Host, r.URL.Path)

	myT := template.Must(template.ParseGlob(tFile))
	myT.ExecuteTemplate(w, tFile, DATA)
}

func main() {

	arguments := os.Args
	if len(arguments) != 3 {
		fmt.Println("Need Database File + Template File!")
		return
	}

	// Munipulate the database - sqlite3 lightweight db
	database := arguments[1]
	tFile = arguments[2]

	db, err := sql.Open("sqlite3", database)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Purge database table.")
	_, err = db.Exec("DELETE FROM data")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Populating", database)
	stmt, _ := db.Prepare("INSERT INTO data(number, double, square) values (?, ?, ?)")
	for i := 20; i < 50; i++ {
		_, _ = stmt.Exec(i, 2*i, i*i)
	}

	rows, err := db.Query("SELECT * FROM data")
	if err != nil {
		fmt.Println(err)
		return
	}

	var n, d, s int
	for rows.Next() {
		err = rows.Scan(&n, &d, &s)
		temp := Entry{Number: n, Double: d, Square: s}
		DATA = append(DATA, temp)
	}

	// startup http server to serve request
	http.HandleFunc("/", dbFileHandler)
	if err = http.ListenAndServe(":9090", nil); err != nil {
		fmt.Println(err)
		return
	}

}
