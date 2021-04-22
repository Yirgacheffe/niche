package student

import "gorm.io/gorm"

type Service struct {
	DB *gorm.DB
}

type Student struct {
	gorm.Model
	FirstName string
	LastName  string
	Age       int
	School    string
}

type StudentService interface {
	GetAllStudents() ([]Student, error)
}

func NewService(db *gorm.DB) *Service {
	return &Service{DB: db}
}
