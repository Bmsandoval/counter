package main

import (
	"fmt"
	"log"
)

// ModifyCounter modifies the counter (increment or decrement)
func (s EventStruct) ModifyCounter(delta int) error {
	svc := s.Services.CounterService

	// TODO : allow overriding defaults based on config file
	filePath, err := svc.GetCounterPath("", "")
	if err != nil {
		fmt.Println("Error determining file path:", err)
		return err
	}

	var counter int
	counter, err = svc.GetCounter(filePath)
	if err != nil {
		fmt.Printf("error getting counter, assuming zero - error: %s", err.Error())
	}

	// Modify the counter
	counter += delta

	err = svc.SetCounter(filePath, counter)
	if err != nil {
		log.Printf("error setting counter - error: %s", err.Error())
	} else {
		fmt.Printf("Counter updated to: %d\n", counter)
	}

	return nil
}
