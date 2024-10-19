package pkg

import (
	"fmt"
	"sync"
)

var (
	mu         sync.Mutex              // Mutex to protect the keys slice
	activeKeys = []string{}            // Slice to store currently pressed keys
	keyStates  = make(map[string]bool) // Map to track the state (pressed/released) of each key
)

// AddKey adds keys to activeKeys slice marking it as pressed
func AddKey(key string) {
	mu.Lock()
	defer mu.Unlock()

	// Check if the key is already in the activeKeys slice
	for _, k := range activeKeys {
		if k == key {
			return // Key is already in the slice
		}
	}
	// Add the new key to the slice
	activeKeys = append(activeKeys, key)
	fmt.Println("Keys pressed:", activeKeys)
}

// RemoveKey marks activeKeys slice items as released
func RemoveKey(key string) {
	mu.Lock()
	defer mu.Unlock()

	// Remove the key from the activeKeys slice
	for i, k := range activeKeys {
		if k == key {
			activeKeys = append(activeKeys[:i], activeKeys[i+1:]...)
			break
		}
	}
	fmt.Println("Keys pressed:", activeKeys)

	// If no keys are pressed, stop or perform some action
	if len(activeKeys) == 0 {
		fmt.Println("All keys released, stopping.")
	}
}

// HandleKeyEvent takes a key event and registers it as a press or release
func HandleKeyEvent(key string, action string) {
	mu.Lock()
	defer mu.Unlock()

	if action == "press" {
		// Handle key press
		if !keyStates[key] { // If the key is not already pressed
			AddKey(key)
			keyStates[key] = true
		}
	} else if action == "release" {
		// Handle key release
		if keyStates[key] { // If the key was pressed
			RemoveKey(key)
			keyStates[key] = false
		}
	}
}
