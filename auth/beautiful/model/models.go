package main

import "time"

type User struct {
	ID           string `json:"id,omitempty"`
	FirstName    string `json:"first_name,omitempty"`
	LastName     string `json:"last_name,omitempty"`
	Email        string `json:"email"`
	Password     string `json:"password,omitempty"`
	HashPassword []byte `json:"hash_password,omitempty`
}
