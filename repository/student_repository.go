package repository

import (
	"enigmacamp.com/uploadDownload/model"
	"errors"
	"fmt"
)

type StudentRepository interface {
	Create(newStudent *model.Student) error
	FindById(id string) (model.Student, error)
}

type studentRepository struct {
	db []model.Student
}

func (s *studentRepository) Create(newStudent *model.Student) error {
	s.db = append(s.db, *newStudent)
	return nil
}

func (s *studentRepository) FindById(id string) (model.Student, error) {
	fmt.Println(s.db)
	for _, s := range s.db {
		if s.Id == id {
			return s, nil
		}
	}
	return model.Student{}, errors.New("Not Found")
}

func NewStudentRepository() StudentRepository {
	return new(studentRepository)
}
