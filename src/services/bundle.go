package services

type Bundle struct {
	InputService         InputService
	CounterService       CounterService
	EventListenerService EventListenerService
}

func NewBundle() Bundle {
	return Bundle{
		InputService:         NewInputService(),
		CounterService:       NewCounterService(),
		EventListenerService: NewEventListenerService(),
	}
}
