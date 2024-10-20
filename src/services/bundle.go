package services

type Bundle struct {
	CounterService CounterService
}

func NewBundle(counterService CounterService) Bundle {
	return Bundle{CounterService: counterService}
}
