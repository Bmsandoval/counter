package main

import (
	"errors"
	"fmt"
	"log"
)

// ModifyCounter modifies the counter (increment or decrement)
func (s EventStruct) ModifyCounter(delta int) error {
	cntSvc := s.Services.CounterService
	cfgSvc := s.Services.ConfigService

	cfg, err := cfgSvc.LoadConfig(s.ConfigFileLoc)
	if err != nil {
		return errors.Join(err, errors.New("unable to load config file"))
	}

	// TODO : allow overriding defaults based on config file
	filePath, err := cntSvc.GetCounterPath(cfg.CounterFileLocation)
	if err != nil {
		fmt.Println("Error determining file path:", err)
		return err
	}

	var counter int
	counter, err = cntSvc.GetCounter(filePath)
	if err != nil {
		fmt.Printf("error getting counter, assuming zero - error: %s", err.Error())
	}

	// Modify the counter
	counter += delta

	err = cntSvc.SetCounter(filePath, counter)
	if err != nil {
		log.Printf("error setting counter - error: %s", err.Error())
	} else {
		fmt.Printf("Counter updated to: %d\n", counter)
	}

	return nil
}
