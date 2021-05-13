package main

import (
	"fmt"
)

// Student reprent student ...
type Student struct {
	ID        int
	FirstName string
	LastName  string
}

// FullName is fullname
func (s Student) FullName() string {
	return s.FirstName + " " + s.LastName
}

// College contains all students
type College struct {
	database map[int]Student // private
}

// Add method add student into the college
func (c *College) Add(payload Student, reply *Student) error {

	id := payload.ID
	if _, ok := c.database[id]; ok {
		return fmt.Errorf("Student with id %d already exists", id)
	}

	c.database[id] = payload
	*reply = payload

	return nil // Return nil without error

}

// Get method return student with that id
func (c *College) Get(payload int, reply *Student) error {

	s, ok := c.database[payload]
	if !ok {
		return fmt.Errorf("No student found with id %d", payload)
	}

	*reply = s
	return nil // Return nil without error

}

// NewCollege function return a college contains the students
func NewCollege() *College {
	return &College{
		database: make(map[int]Student),
	}
}
