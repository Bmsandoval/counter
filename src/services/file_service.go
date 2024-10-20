package services

import (
	"errors"
	"os"
)

type FileService interface {
	GetFile(filePath string) (string, error)
}

type fileServiceImpl struct{}

func NewFileService() FileService {
	return &fileServiceImpl{}
}

func (s fileServiceImpl) GetFile(filePath string) (string, error) {
	// Open and read the current value from the file
	file, err := os.ReadFile(filePath)
	if err != nil {
		return "", errors.Join(err, errors.New("failed to read file"))
	}

	return string(file), nil
}
