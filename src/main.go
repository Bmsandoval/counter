package main

import (
	"flag"
	"github.com/bmsandoval/counter/src/services"
	"github.com/bmsandoval/counter/src/utils"
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
		finalHk = append(finalHk, utils.MapKeyToID(key))
	}

	svcBundle := services.Bundle{
		CounterService: services.NewCounterService(),
	}

	eventHandler := EventStruct{Services: svcBundle}
	mainthread.Init(eventHandler.EventLoop(finalHk))
}

type EventStruct struct {
	Services services.Bundle
}

func (s EventStruct) EventLoop(hks []int) func() {
	return func() {
		var modifiers []hotkey.Modifier
		var keyboardKey hotkey.Key
		for _, val := range hks {
			if _, ok := utils.ModifierKeyMap[val]; ok {
				modifiers = append(modifiers, hotkey.Modifier(val))
			} else {
				keyboardKey = hotkey.Key(val)
			}
		}

		hk := hotkey.New(modifiers, keyboardKey)
		for kp := range GlobalHotkeyListener(hk) {
			if kp == hk.String() {
				log.Printf("hotkey: %s is down\n", kp)
				err := s.ModifyCounter(1)
				if err != nil {
					log.Fatalf("failed to modify counter: %v\n", err)
					return
				}
			}
		}
	}
}
