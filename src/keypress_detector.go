package main

import (
	"golang.design/x/hotkey"
	"log"
)

// GlobalHotkeyListener listens for any hotkey
func GlobalHotkeyListener(yield func(string) bool) {
	var hotkeys []*hotkey.Hotkey
	// TODO : allow creating custom hotkeys
	hotkeys = append(hotkeys, hotkey.New([]hotkey.Modifier{hotkey.ModCtrl, hotkey.ModShift}, hotkey.KeyS))
	hotkeys = append(hotkeys, hotkey.New(nil, hotkey.KeyEscape))

	for _, hk := range hotkeys {
		err := hk.Register()
		if err != nil {
			log.Fatalf("hotkey: failed to register hotkey: %v", err)
			return
		}
	}
	keypressEventChannel := make(chan string)
	defer func(hks []*hotkey.Hotkey, kpEvent chan string) {
		for _, hk := range hks {
			err := hk.Unregister()
			if err != nil {
				log.Fatalf("hotkey: failed to unregister hotkey: %v", err)
			}
		}
		close(kpEvent)
		log.Println("hotkey: all channels closed")
	}(hotkeys, keypressEventChannel)

	for _, hk := range hotkeys {
		go func(hk *hotkey.Hotkey, outChan chan string) {
			for range hk.Keydown() {
				<-hk.Keyup()
				outChan <- hk.String()
			}
		}(hk, keypressEventChannel)
	}

	for key := range keypressEventChannel {
		if !yield(key) {
			return
		}
	}
}
