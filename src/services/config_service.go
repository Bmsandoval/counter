package services

import (
	"encoding/json"
	"errors"
	"github.com/bmsandoval/counter/src/types"
	"os"
	"os/user"
	"path/filepath"
)

type ConfigService interface {
	SaveConfig(config types.Config) error
	LoadConfig(filePath string) (types.Config, error)
}

type configServiceImpl struct {
}

func NewConfigService() ConfigService {
	return &configServiceImpl{}
}

// SaveConfig updates the config file with key combinations
func (s configServiceImpl) SaveConfig(config types.Config) error {
	usr, err := user.Current()
	if err != nil {
		errors.Join(err, errors.New("error retrieving home directory: unable to get current user"))
	}
	home := usr.HomeDir

	configPath := filepath.Join(home, types.DefaultConfigFileName)
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(configPath, data, 0644)
}

// LoadConfig retrieves the key combination(s) from config file
func (s configServiceImpl) LoadConfig(configFileLoc string) (types.Config, error) {
	if len(configFileLoc) < 1 {
		usr, err := user.Current()
		if err != nil {
			// TODO : handle gracefully
			panic(err)
		}
		home := usr.HomeDir

		// TODO : fix hardcoded config file
		configFileLoc = filepath.Join(home, types.DefaultConfigFileName)
	}

	// Check if the config file exists
	var config types.Config
	file, err := os.ReadFile(configFileLoc)
	if os.IsNotExist(err) {
		// Set default key combinations
		config = types.Config{IncrementHotkey: "ctrl+shift+p"}
		return config, s.SaveConfig(config)
	}

	err = json.Unmarshal(file, &config)
	return config, err
}
