package types

var DefaultConfigFileName = ".counterconfig.json"

type Config struct {
	IncrementHotkey     string `json:"increment_hotkey"`
	CounterFileLocation string `json:"counter_file_location"`
}
