package pkg

import (
	"gobot.io/x/gobot/platforms/keyboard"
	"log"
	"os"
	"os/signal"
)

// ListenForHotkey will listen for a sequence of key presses
// up until the escape key pressed
// returns the sequence of keys pressed as a string, up to but excluding escape
func ListenForHotkey() {
	// Initialize the keyboard driver
	keys := keyboard.NewDriver()

	// Handle key events
	err := keys.On(keyboard.Key, func(data interface{}) {
		key := data.(keyboard.KeyEvent)
		log.Println(key.Char)
		//resp <- key.Char
	})
	if err != nil {
		return
	}

	// Listen for Ctrl+C to safely exit
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}
