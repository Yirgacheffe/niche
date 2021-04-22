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
	GetStudentByID(ID uint) (Student, error)
	GetStudentBySchool(school string) (Student, error)

	PostStudent(student Student) (Student, error)
	UpdateStudent(IO uint, newStudent Student) (Student, error)
	DeleteStudent(ID uint) error
}

func NewService(db *gorm.DB) *Service {
	return &Service{DB: db}
}

func (s *Service) GetAllStudents() ([]Student, error) {
	var students []Student

	if result := s.DB.Find(&students); result.Error != nil {
		return nil, result.Error
	}

	return students, nil
}
