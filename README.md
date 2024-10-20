# App

Counter (badly named) will be a golang-based successor to "DCSB", but focused more on the counter (death or otherwise)

## Plan

### Problem:

* DCSB allows displaying a running count on stream, but doesn't work on MacOS

### **Active Goal(s)**

* [ ]  Unit Tests
  * [ ]  Enable dependency injection

### V0.2:

* [ ]  user defined hotkey via keyboard input
* [ ]  user defined file location
* [ ]  tests (you win this time DevNinYa)

## Future

### v1.0:

* [ ]  Multiple hotkeys - allow increment/decrement
* [ ]  Multiple files - allow tracking multiple counters

### v1.1:

* [ ]  allow grouping counters hoping to track "Today" and "AllTime" with a single click

### v2.0:

* [ ]  Enable Soundboard
  * [ ]  hotkey triggers sound to play (research golang packages for cross-platform, or do bash call as DevNinYa suggested)

## Backlog

* [ ]  GUI

## Version History

### v0.1

* Can click a global hotkey to increment an integer stored in a file.
* Hotkey can be passed in via cli flag

## Usage:

TODO : fill this in
install and setup golang
ensure golang is in your PATH
clone this repo
run go mod vendor within the repo
go run source/main.go
