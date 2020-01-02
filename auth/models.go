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

type Task struct {
	ID          string    `json:"id,omitempty"`
	CreatedBy   string    `json:"created_by"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedOn   time.Time `json:"created_on"`
	Due         time.Time `json:"due,omitempty"`
	Status      string    `json:"status,omitempty"`
	Tags        []string  `json:"tags,omitempty"`
}

type TaskNote struct {
	ID          string `json:"id,omitempty"`
	TaskID      string `json:"task_id,omitempty"`
	Description string `json:"description"`
	CreatedOn   string `json:"create_on,omitempty"`
}
