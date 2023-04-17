package usecase

import (
	"enigmacamp.com/uploadDownload/model"
	"enigmacamp.com/uploadDownload/repository"
	"log"
	"mime/multipart"
)

type StudentRegistrationUseCase interface {
	Register(newStudent *model.Student, file *multipart.File, fileExt string) error
	GetPhotoProfile(id string) (model.Student, error)
}

type studentRegistrationUseCase struct {
	fileRepo    repository.FileRepository
	studentRepo repository.StudentRepository
}

func (s *studentRegistrationUseCase) Register(newStudent *model.Student, file *multipart.File, fileExt string) error {
	filePath, err := s.fileRepo.Save("student-"+newStudent.Id+"."+fileExt, file)
	if err != nil {
		return err
	}
	log.Println(filePath)
	newStudent.PhotoProfile = filePath
	err = s.studentRepo.Create(newStudent)
	if err != nil {
		return err
	}
	return nil
}

func (s *studentRegistrationUseCase) GetPhotoProfile(id string) (model.Student, error) {
	student, err := s.studentRepo.FindById(id)
	if err != nil {
		return model.Student{}, err
	}
	return student, nil
}

func NewStudentRegistrationUseCase(repo repository.FileRepository, studentRepo repository.StudentRepository) StudentRegistrationUseCase {
	return &studentRegistrationUseCase{fileRepo: repo, studentRepo: studentRepo}
}
