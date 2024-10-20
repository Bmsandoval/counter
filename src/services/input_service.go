package services

import "errors"

type InputService interface {
}

type inputServiceImpl struct{}

func NewInputService() InputService {
	return &inputServiceImpl{}
}

// GetFileLocFromUser is used to prompt the user for a location to store a counter
func (s inputServiceImpl) GetFileLocFromUser() (string, error) {
	return "", errors.New("not implemented")
}
