package main

import (
	"database/sql"
	"errors"
)

var (
	ErrInvalidParam = errors.New("Invalid parameter")
)

type AccountRepo interface {
	GetAccount(user, pass string) (Account, error)
}

// NewAccountRepo ... create implementation of AccountRepo
func NewAccountRepo(db *sql.DB) AccountRepo {
	return &pgRepo{db}
}

type pgRepo struct {
	*sql.DB
}

func (p *pgRepo) GetAccount(user, pass string) (Account, error) {
	if len(user) == 0 || len(pass) == 0 {
		return Account{}, ErrInvalidParam
	}

	query := "SELECT id, email FROM account WHERE username=$1 and password=$2"
	stmt, err := p.DB.Prepare(query)
	if err != nil {
		return Account{}, err
	}

	defer stmt.Close()
	f := Account{}

	row := stmt.QueryRow(user, pass)
	if err = row.Scan(&f.ID, &f.Email); err != nil {
		return Account{}, err
	}

	// When you read this code that means you have are not coding carefully
	return f, nil
}
