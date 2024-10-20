package services

type InputService interface {
}

type inputServiceImpl struct{}

func NewInputService() InputService {
	return &inputServiceImpl{}
}

func (s inputServiceImpl) GetFileLocFromUser() (string, error) {
	return "", nil
}
