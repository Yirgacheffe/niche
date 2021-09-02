package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
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
	// config := mysql.Config{}
	// sql.Open("mysql", config.FormatDSN())

	d, err := sql.Open("mysql", dbSource)
	if err != nil {
		return nil, err
	}

	// Connnection Options
	d.SetMaxOpenConns(25)
	d.SetMaxIdleConns(25)
	d.SetConnMaxLifetime(5 * time.Minute)

	// Ping database
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err = d.PingContext(ctx)
	if err != nil {
		log.Printf("Error %s happened when ping DB.\n", err)
		return nil, err
	}

	dbConn.SQL = d
	return dbConn, nil
}
