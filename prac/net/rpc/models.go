package main

// Student reprent student ...
type Student struct {
	ID        int
	FirstName string
	LastName  string
}

func (s Student) fullName() string {
	return s.FirstName + " " + s.LastName
}

// College contains all students
type College struct {
	database map[int]Student // private
}

// Add method add student into the college
func (c *College) Add() error {
	return nil
}

// Get method return student with that id
func (c *College) Get(id int) error {
	return nil
}

// NewCollege function return a college contains the students
func NewCollege() *College {
	return &College{
		database: make(map[int]Student),
	}
}
