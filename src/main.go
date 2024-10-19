package main

import (
	"golang.design/x/hotkey/mainthread"
	"log"
)

func main() {
	mainthread.Init(EventLoop)
}

func EventLoop() {
	for kp := range GlobalHotkeyListener {
		log.Printf("hotkey: %s is down\n", kp)
	}
}
