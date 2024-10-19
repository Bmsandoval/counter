# counter

Counter (badly named) will be a golang-based successor to "DCSB", but focused more on the counter (death or otherwise)


## Project Plan

### Problem:

* DCSB doesn't work on MacOS
* OBS text display accepts a constant string, or pulling from a file

### Goal:

create an app that clones the 'counter' feature of DCSB to allow a global hotkey for incrementing a counter

### Requirements:

* Read from/Write to file
* Global Hotkeys
* Possible to manage file and hotkeys
  * TUI/GUI?
  * For quick POC, could be passed in as flags
