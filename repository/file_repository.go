package repository

import (
	"enigmacamp.com/uploadDownload/util"
	"mime/multipart"
	"path/filepath"
)

type FileRepository interface {
	Save(fileName string, file *multipart.File) (string, error)
}

type fileRepository struct {
	fileBasePath string
}

func (f *fileRepository) Save(fileName string, file *multipart.File) (string, error) {
	fileLocation := filepath.Join(f.fileBasePath, fileName)
	err := util.SaveToLocalFile(fileLocation, file)
	if err != nil {
		return "", err
	}
	return fileLocation, nil
}

func NewFileRepository(basePath string) FileRepository {
	fileRepo := fileRepository{fileBasePath: basePath}
	return &fileRepo
}
