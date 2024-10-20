package services

type Bundle struct {
	InputService         InputService
	ConfigService        ConfigService
	CounterService       CounterService
	EventListenerService EventListenerService
}

func NewBundle() Bundle {
	return Bundle{
		InputService:         NewInputService(),
		ConfigService:        NewConfigService(),
		CounterService:       NewCounterService(),
		EventListenerService: NewEventListenerService(),
	}
}
