package utils

import "strings"

var ModifierKeyMap = map[int]struct{}{
	0x1000: {},
	0x200:  {},
	0x800:  {},
	0x100:  {},
}

func MapKeyToID(keyboardKey string) int {
	var KeyToIDMapper = map[string]int{
		"CTRL":   0x1000, // ModCtrl
		"SHIFT":  0x200,  // ModShift
		"OPTION": 0x800,  // ModOption
		"CMD":    0x100,  // ModCmd
		"SPACE":  49,     // KeySpace
		"1":      18,     // Key1
		"2":      19,     // Key2
		"3":      20,     // Key3
		"4":      21,     // Key4
		"5":      23,     // Key5
		"6":      22,     // Key6
		"7":      26,     // Key7
		"8":      28,     // Key8
		"9":      25,     // Key9
		"0":      29,     // Key0
		"A":      0,      // KeyA
		"B":      11,     // KeyB
		"C":      8,      // KeyC
		"D":      2,      // KeyD
		"E":      14,     // KeyE
		"F":      3,      // KeyF
		"G":      5,      // KeyG
		"H":      4,      // KeyH
		"I":      34,     // KeyI
		"J":      38,     // KeyJ
		"K":      40,     // KeyK
		"L":      37,     // KeyL
		"M":      46,     // KeyM
		"N":      45,     // KeyN
		"O":      31,     // KeyO
		"P":      35,     // KeyP
		"Q":      12,     // KeyQ
		"R":      15,     // KeyR
		"S":      1,      // KeyS
		"T":      17,     // KeyT
		"U":      32,     // KeyU
		"V":      9,      // KeyV
		"W":      13,     // KeyW
		"X":      7,      // KeyX
		"Y":      16,     // KeyY
		"Z":      6,      // KeyZ
		"RETURN": 0x24,   // KeyReturn
		"ESCAPE": 0x35,   // KeyEscape
		"DELETE": 0x33,   // KeyDelete
		"TAB":    0x30,   // KeyTab
		"LEFT":   0x7B,   // KeyLeft
		"RIGHT":  0x7C,   // KeyRight
		"UP":     0x7E,   // KeyUp
		"DOWN":   0x7D,   // KeyDown
		"F1":     0x7A,   // KeyF1
		"F2":     0x78,   // KeyF2
		"F3":     0x63,   // KeyF3
		"F4":     0x76,   // KeyF4
		"F5":     0x60,   // KeyF5
		"F6":     0x61,   // KeyF6
		"F7":     0x62,   // KeyF7
		"F8":     0x64,   // KeyF8
		"F9":     0x65,   // KeyF9
		"F10":    0x6D,   // KeyF10
		"F11":    0x67,   // KeyF11
		"F12":    0x6F,   // KeyF12
		"F13":    0x69,   // KeyF13
		"F14":    0x6B,   // KeyF14
		"F15":    0x71,   // KeyF15
		"F16":    0x6A,   // KeyF16
		"F17":    0x40,   // KeyF17
		"F18":    0x4F,   // KeyF18
		"F19":    0x50,   // KeyF19
		"F20":    0x5A,   // KeyF20
	}
	return KeyToIDMapper[strings.ToUpper(keyboardKey)]
}

func MapIDToKey(id int) string {
	var IDToKeyMapper = map[int]string{
		0x1000: "CTRL",   // ModCtrl
		0x200:  "SHIFT",  // ModShift
		0x800:  "OPTION", // ModOption
		0x100:  "CMD",    // ModCmd
		49:     "SPACE",  // KeySpace
		18:     "1",      // Key1
		19:     "2",      // Key2
		20:     "3",      // Key3
		21:     "4",      // Key4
		23:     "5",      // Key5
		22:     "6",      // Key6
		26:     "7",      // Key7
		28:     "8",      // Key8
		25:     "9",      // Key9
		29:     "0",      // Key0
		0:      "A",      // KeyA
		11:     "B",      // KeyB
		8:      "C",      // KeyC
		2:      "D",      // KeyD
		14:     "E",      // KeyE
		3:      "F",      // KeyF
		5:      "G",      // KeyG
		4:      "H",      // KeyH
		34:     "I",      // KeyI
		38:     "J",      // KeyJ
		40:     "K",      // KeyK
		37:     "L",      // KeyL
		46:     "M",      // KeyM
		45:     "N",      // KeyN
		31:     "O",      // KeyO
		35:     "P",      // KeyP
		12:     "Q",      // KeyQ
		15:     "R",      // KeyR
		1:      "S",      // KeyS
		17:     "T",      // KeyT
		32:     "U",      // KeyU
		9:      "V",      // KeyV
		13:     "W",      // KeyW
		7:      "X",      // KeyX
		16:     "Y",      // KeyY
		6:      "Z",      // KeyZ
		0x24:   "RETURN", // KeyReturn
		0x35:   "ESCAPE", // KeyEscape
		0x33:   "DELETE", // KeyDelete
		0x30:   "TAB",    // KeyTab
		0x7B:   "LEFT",   // KeyLeft
		0x7C:   "RIGHT",  // KeyRight
		0x7E:   "UP",     // KeyUp
		0x7D:   "DOWN",   // KeyDown
		0x7A:   "F1",     // KeyF1
		0x78:   "F2",     // KeyF2
		0x63:   "F3",     // KeyF3
		0x76:   "F4",     // KeyF4
		0x60:   "F5",     // KeyF5
		0x61:   "F6",     // KeyF6
		0x62:   "F7",     // KeyF7
		0x64:   "F8",     // KeyF8
		0x65:   "F9",     // KeyF9
		0x6D:   "F10",    // KeyF10
		0x67:   "F11",    // KeyF11
		0x6F:   "F12",    // KeyF12
		0x69:   "F13",    // KeyF13
		0x6B:   "F14",    // KeyF14
		0x71:   "F15",    // KeyF15
		0x6A:   "F16",    // KeyF16
		0x40:   "F17",    // KeyF17
		0x4F:   "F18",    // KeyF18
		0x50:   "F19",    // KeyF19
		0x5A:   "F20",    // KeyF20
	}
	return IDToKeyMapper[id]
}
