package main

import (
	"flag"
	"github.com/bmsandoval/counter/src/services"
	"golang.design/x/hotkey"
	"golang.design/x/hotkey/mainthread"
	"log"
	"strings"
)

func main() {
	hk := flag.String("hotkey", "ctrl+shift+s", "increment hotkey")
	flag.Parse()

	var finalHk []int
	for _, key := range strings.Split(*hk, "+") {
		finalHk = append(finalHk, MapKeyToID(key))
	}

	eventHandler := EventStruct{}
	mainthread.Init(eventHandler.EventLoop(finalHk))
}

type EventStruct struct {
	FileService services.FileService
}

func (e EventStruct) EventLoop(hks []int) func() {
	return func() {
		var modifiers []hotkey.Modifier
		var keyboardKey hotkey.Key
		for _, val := range hks {
			if _, ok := ModifierKeyMap[val]; ok {
				modifiers = append(modifiers, hotkey.Modifier(val))
			} else {
				keyboardKey = hotkey.Key(val)
			}
		}

		hk := hotkey.New(modifiers, keyboardKey)
		for kp := range GlobalHotkeyListener(hk) {
			if kp == hk.String() {
				log.Printf("hotkey: %s is down\n", kp)
				err := ModifyCounter(1)
				if err != nil {
					log.Fatalf("failed to modify counter: %v\n", err)
					return
				}
			}
		}
	}
}
