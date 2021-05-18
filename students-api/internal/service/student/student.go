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
	PostStudent(student Student) (Student, error)
	GetStudentByID(ID uint) (Student, error)
	UpdateStudent(IO uint, newStudent Student) (Student, error)
	DeleteStudent(ID uint) error

	GetAllStudents() ([]Student, error)
	GetStudentBySchool(school string) (Student, error)
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

func (s *Service) GetStudentByID(ID uint) (Student, error) {
	var student Student
	result := s.DB.First(&student, ID)
	if result.Error != nil {
		return student, result.Error
	}
	return student, nil
}

func (s *Service) GetStudentBySchool(school string) ([]Student, error) {
	var students []Student
	result := s.DB.Find(&students).Where("school=?", school)
	if result.Error != nil {
		return nil, result.Error
	}
	return students, nil
}

func (s *Service) PostStudent(student Student) (Student, error) {
	if result := s.DB.Save(&student); result.Error != nil {
		return Student{}, result.Error
	}
	return student, nil
}

func (s *Service) UpdateStudent(ID uint, newStudent Student) (Student, error) {
	student, err := s.GetStudentByID(ID)
	if err != nil {
		return Student{}, err
	}

	if result := s.DB.Model(&student).Updates(newStudent); result.Error != nil {
		return Student{}, result.Error
	}

	return student, nil
}

func (s *Service) DeleteStudent(ID uint) error {
	if result := s.DB.Delete(&Student{}, ID); result.Error != nil {
		return result.Error
	}
	return nil
}
