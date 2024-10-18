package main

import (
	"fmt"
	"github.com/bmsandoval/counter/pkg"
	"os"
	"path/filepath"
)

// ModifyCounter modifies the counter (increment or decrement)
func ModifyCounter(delta int) error {
	filePath, err := GetCounterFilePath()
	if err != nil {
		fmt.Println("Error determining file path:", err)
		return err
	}

	// Open and read the current value from the file
	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("failed to read file, creating. See error: %s", err.Error())
		err = nil
	}

	// Convert the content to an integer
	var counter int
	_, err = fmt.Sscanf(string(file), "%d", &counter)
	if err != nil {
		fmt.Printf("failed to parse file, assuming zero. See error: %s", err.Error())
		counter = 0
	}

	// Modify the counter
	counter += delta

	// Write the updated counter back to the file
	err = os.WriteFile(filePath, []byte(fmt.Sprintf("%d", counter)), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return err
	}
	fmt.Printf("Counter updated to: %d\n", counter)
	return nil
}

// GetCounterFilePath gets the path of the counter file
func GetCounterFilePath() (string, error) {
	// TODO : remove hardcoded file path
	home, err := pkg.HomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, "counter.txt"), nil
}
