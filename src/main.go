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
	hk := flag.String("hotkey", "ctrl+shift+s", "increment hotkey. Defaults to ctrl+shift+s")
	cfgFileLoc := flag.String("config", "~/.counterconfig.json", "path to config file. Defaults to ~/.counterconfig.json")
	flag.Parse()

	var finalHk []int
	for _, key := range strings.Split(*hk, "+") {
		finalHk = append(finalHk, utils.MapKeyToID(key))
	}

	svcBundle := services.NewBundle()

	eventHandler := EventStruct{Services: svcBundle, ConfigFileLoc: *cfgFileLoc}
	mainthread.Init(eventHandler.EventLoop(finalHk))
}

type EventStruct struct {
	Services      services.Bundle
	ConfigFileLoc string
}

func (s EventStruct) EventLoop(hks []int) func() {
	return func() {
		svc := s.Services.EventListenerService
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
		for kp := range svc.ListenForHotkeyEvents([]*hotkey.Hotkey{hk}) {
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
