package main

//import (
//	"encoding/json"
//	"github.com/bmsandoval/counter/pkg"
//	"os"
//	"path/filepath"
//)
//
//// SaveConfig updates the config file with key combinations
//func SaveConfig(config Config) error {
//	home, err := pkg.HomeDir()
//	if err != nil {
//		return err
//	}
//	configPath := filepath.Join(home, ".filecounterconfig.json")
//	data, err := json.MarshalIndent(config, "", "  ")
//	if err != nil {
//		return err
//	}
//	return os.WriteFile(configPath, data, 0644)
//}
//
//// LoadConfig retrieves the key combination(s) from config file
//func LoadConfig() (Config, error) {
//	home, err := pkg.HomeDir()
//	if err != nil {
//		// TODO : handle gracefully
//		panic(err)
//	}
//
//	// TODO : fix hardcoded config file
//	configPath := filepath.Join(home, ".filecounterconfig.json")
//	var config Config
//
//	// Check if the config file exists
//	if _, err := os.Stat(configPath); os.IsNotExist(err) {
//		// Set default key combinations
//		config = Config{IncrementKey: "ctrl+shift+p"}
//		return config, SaveConfig(config)
//	}
//
//	file, err := os.ReadFile(configPath)
//	if err != nil {
//		return config, err
//	}
//	err = json.Unmarshal(file, &config)
//	return config, err
//}
//
//// Config structure to store key combinations
//type Config struct {
//	IncrementKey string `json:"increment_key"`
//}
