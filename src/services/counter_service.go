package services

import (
	"errors"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

type CounterService interface {
	GetCounter(filePath string) (int, error)
	SetCounter(filePath string, counter int) error
	GetCounterPath(filePath, fileName string) (string, error)
}

type counterServiceImpl struct{}

func NewCounterService() CounterService {
	return &counterServiceImpl{}
}

func (s counterServiceImpl) GetCounter(filePath string) (int, error) {
	// Open and read the current value from the file
	file, err := os.ReadFile(filePath)
	if err != nil {
		return 0, errors.Join(err, errors.New("unable to read file"))
	}

	// Convert the content to an integer
	var counter int
	_, err = fmt.Sscanf(string(file), "%d", &counter)
	if err != nil {
		return 0, errors.Join(err, errors.New("unable to parse file"))
	}

	return counter, nil
}

func (s counterServiceImpl) SetCounter(filePath string, counter int) error {
	// Write the updated counter back to the file
	err := os.WriteFile(filePath, []byte(fmt.Sprintf("%d", counter)), 0644)
	if err != nil {
		return errors.Join(err, errors.New("unable to write counter"))
	}
	return nil
}

// GetCounterPath gets the path of the counter file
// defaults to '~/counter.txt'
func (s counterServiceImpl) GetCounterPath(filePath, fileName string) (string, error) {
	if len(filePath) < 1 {
		usr, err := user.Current()
		if err != nil {
			return "", errors.Join(err, errors.New("error retrieving home directory: unable to get current user"))
		}
		filePath = usr.HomeDir
	}

	if len(fileName) < 1 {
		fileName = "counter.txt"
	}

	return filepath.Join(filePath, fileName), nil
}
