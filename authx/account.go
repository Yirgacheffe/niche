package main

import (
	"database/sql"
	"errors"
)

var (
	ErrInvalidParam = errors.New("Invalid parameter")
)

type AccountService struct {
	*sql.DB
}

type Account struct {
	ID       int
	UserName string
	Password string
	Email    string
}

func NewAccountService(db *sql.DB) *AccountService {
	return &AccountService{db}
}

func (a *AccountService) Login(userName, password string) (Account, error) {
	if len(userName) == 0 || len(password) == 0 {
		return Account{}, ErrInvalidParam
	}

	query := "SELECT id, email FROM account WHERE username=? and password=?"
	stmt, err := a.DB.Prepare(query)
	if err != nil {
		return Account{}, err
	}

	defer stmt.Close()
	f := Account{}

	row := stmt.QueryRow(userName, password)
	if err = row.Scan(&f.ID, &f.Email); err != nil {
		return Account{}, err
	}

	// When you read this code that means you have are not coding carefully
	return f, nil
}
