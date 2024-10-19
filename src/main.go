package main

import (
	"flag"
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

	mainthread.Init(EventLoop(finalHk))
}

func EventLoop(hks []int) func() {
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

		for kp := range GlobalHotkeyListener {
			if kp == hotkey.New(modifiers, keyboardKey).String() {
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
