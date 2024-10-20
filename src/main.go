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
	cfgFileLoc := flag.String("config", "", "path to config file. Defaults to ~/.counterconfig.json")
	flag.Parse()

	svcBundle := services.NewBundle()

	cfg, err := svcBundle.ConfigService.LoadConfig(*cfgFileLoc)
	if err != nil {
		panic(err)
	}

	var finalHk []int
	for _, key := range strings.Split(cfg.IncrementHotkey, "+") {
		finalHk = append(finalHk, utils.MapKeyToID(key))
	}

	eventHandler := EventStruct{Services: svcBundle, ConfigFileLoc: *cfgFileLoc}
	mainthread.Init(eventHandler.EventLoop(finalHk))
}

type EventStruct struct {
	Services      services.Bundle
	ConfigFileLoc string
}

func (s EventStruct) EventLoop(hks []int) func() {
	fileName, err := s.Services.InputService.PromptUserForFileLoc()
	if err != nil {
		log.Fatalf("Error reading input file: %v", err)
	}
	_ = fileName

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
				err := s.ModifyCounter(1)
				if err != nil {
					log.Fatalf("failed to modify counter: %v\n", err)
					return
				}
			}
		}
	}
}
