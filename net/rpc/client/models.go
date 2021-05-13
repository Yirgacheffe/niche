package main

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
