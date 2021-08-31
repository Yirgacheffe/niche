package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

type DB struct {
	*sql.DB
}

type Tx struct {
	*sql.Tx
}

var dbConn = &DB{}

// Begin starts a transaction, wrap and return db transaction with 'Tx'
func (d *DB) Begin() (*Tx, error) {
	tx, err := d.DB.Begin()
	if err != nil {
		return nil, errors.Wrap(err, "tx begin failed")
	}

	return &Tx{tx}, nil
}

func ConnectSQL(host, port, user, password, dbname string) (*DB, error) {
	psqlConn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		dbname,
	)

	db, err := sql.Open("postgres", psqlConn)
	if err != nil {
		return nil, err
	}

	// sql.Open("postgres", "postgres://user:pass@localhost/db")

	// Note: IdleConns should always be less than or equal to openConns
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxIdleTime(10 * time.Minute)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "ping db failed")
	}

	// dbConn.SQL = db
	return &DB{db}, nil
}
