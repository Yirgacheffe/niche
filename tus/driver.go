package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// DB ...
type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

// ConnectSQL ... MySQL connection
func ConnectSQL(host, port, uname, passwd, dbname string) (*DB, error) {

	dbSource := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
		uname,
		passwd,
		host,
		port,
		dbname,
	)

	d, err := sql.Open("mysql", dbSource)
	if err != nil {
		return nil, err
	}

	d.SetMaxOpenConns(20)
	d.SetConnMaxLifetime(5 * time.Minute)

	dbConn.SQL = d
	return dbConn, nil
}
