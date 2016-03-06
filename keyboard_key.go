// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gxui

type KeyboardKey int

const (
	KeyUnknown KeyboardKey = iota
	KeySpace
	KeyApostrophe
	KeyComma
	KeyMinus
	KeyPeriod
	KeySlash
	Key0
	Key1
	Key2
	Key3
	Key4
	Key5
	Key6
	Key7
	Key8
	Key9
	KeySemicolon
	KeyEqual
	KeyA
	KeyB
	KeyC
	KeyD
	KeyE
	KeyF
	KeyG
	KeyH
	KeyI
	KeyJ
	KeyK
	KeyL
	KeyM
	KeyN
	KeyO
	KeyP
	KeyQ
	KeyR
	KeyS
	KeyT
	KeyU
	KeyV
	KeyW
	KeyX
	KeyY
	KeyZ
	KeyLeftBracket
	KeyBackslash
	KeyRightBracket
	KeyGraveAccent
	KeyWorld1
	KeyWorld2
	KeyEscape
	KeyEnter
	KeyTab
	KeyBackspace
	KeyInsert
	KeyDelete
	KeyRight
	KeyLeft
	KeyDown
	KeyUp
	KeyPageUp
	KeyPageDown
	KeyHome
	KeyEnd
	KeyCapsLock
	KeyScrollLock
	KeyNumLock
	KeyPrintScreen
	KeyPause
	KeyF1
	KeyF2
	KeyF3
	KeyF4
	KeyF5
	KeyF6
	KeyF7
	KeyF8
	KeyF9
	KeyF10
	KeyF11
	KeyF12
	KeyKp0
	KeyKp1
	KeyKp2
	KeyKp3
	KeyKp4
	KeyKp5
	KeyKp6
	KeyKp7
	KeyKp8
	KeyKp9
	KeyKpDecimal
	KeyKpDivide
	KeyKpMultiply
	KeyKpSubtract
	KeyKpAdd
	KeyKpEnter
	KeyKpEqual
	KeyLeftShift
	KeyLeftControl
	KeyLeftAlt
	KeyLeftSuper
	KeyRightShift
	KeyRightControl
	KeyRightAlt
	KeyRightSuper
	KeyMenu
	KeyLast
)

func (k KeyboardKey) String() string {
	switch k {
	case KeySpace:
		return "Space"
	case KeyApostrophe:
		return "'"
	case KeyComma:
		return ","
	case KeyMinus:
		return "-"
	case KeyPeriod:
		return "."
	case KeySlash:
		return "/"
	case Key0:
		return "0"
	case Key1:
		return "1"
	case Key2:
		return "2"
	case Key3:
		return "3"
	case Key4:
		return "4"
	case Key5:
		return "5"
	case Key6:
		return "6"
	case Key7:
		return "7"
	case Key8:
		return "8"
	case Key9:
		return "9"
	case KeySemicolon:
		return ";"
	case KeyEqual:
		return "="
	case KeyA:
		return "A"
	case KeyB:
		return "B"
	case KeyC:
		return "C"
	case KeyD:
		return "D"
	case KeyE:
		return "E"
	case KeyF:
		return "F"
	case KeyG:
		return "G"
	case KeyH:
		return "H"
	case KeyI:
		return "I"
	case KeyJ:
		return "J"
	case KeyK:
		return "K"
	case KeyL:
		return "L"
	case KeyM:
		return "M"
	case KeyN:
		return "N"
	case KeyO:
		return "O"
	case KeyP:
		return "P"
	case KeyQ:
		return "Q"
	case KeyR:
		return "R"
	case KeyS:
		return "S"
	case KeyT:
		return "T"
	case KeyU:
		return "U"
	case KeyV:
		return "V"
	case KeyW:
		return "W"
	case KeyX:
		return "X"
	case KeyY:
		return "Y"
	case KeyZ:
		return "Z"
	case KeyLeftBracket:
		return "["
	case KeyBackslash:
		return "\\"
	case KeyRightBracket:
		return "]"
	case KeyGraveAccent:
		return "`"
	case KeyWorld1:
		return "World1"
	case KeyWorld2:
		return "World2"
	case KeyEscape:
		return "Escape"
	case KeyEnter:
		return "Enter"
	case KeyTab:
		return "Tab"
	case KeyBackspace:
		return "Backspace"
	case KeyInsert:
		return "Insert"
	case KeyDelete:
		return "Delete"
	case KeyRight:
		return "Right"
	case KeyLeft:
		return "Left"
	case KeyDown:
		return "Down"
	case KeyUp:
		return "Up"
	case KeyPageUp:
		return "PageUp"
	case KeyPageDown:
		return "PageDown"
	case KeyHome:
		return "Home"
	case KeyEnd:
		return "End"
	case KeyCapsLock:
		return "CapsLock"
	case KeyScrollLock:
		return "ScrollLock"
	case KeyNumLock:
		return "NumLock"
	case KeyPrintScreen:
		return "PrintScreen"
	case KeyPause:
		return "Pause"
	case KeyF1:
		return "F1"
	case KeyF2:
		return "F2"
	case KeyF3:
		return "F3"
	case KeyF4:
		return "F4"
	case KeyF5:
		return "F5"
	case KeyF6:
		return "F6"
	case KeyF7:
		return "F7"
	case KeyF8:
		return "F8"
	case KeyF9:
		return "F9"
	case KeyF10:
		return "F10"
	case KeyF11:
		return "F11"
	case KeyF12:
		return "F12"
	case KeyKp0:
		return "Kp0"
	case KeyKp1:
		return "Kp1"
	case KeyKp2:
		return "Kp2"
	case KeyKp3:
		return "Kp3"
	case KeyKp4:
		return "Kp4"
	case KeyKp5:
		return "Kp5"
	case KeyKp6:
		return "Kp6"
	case KeyKp7:
		return "Kp7"
	case KeyKp8:
		return "Kp8"
	case KeyKp9:
		return "Kp9"
	case KeyKpDecimal:
		return "KpDecimal"
	case KeyKpDivide:
		return "KpDivide"
	case KeyKpMultiply:
		return "KpMultiply"
	case KeyKpSubtract:
		return "KpSubtract"
	case KeyKpAdd:
		return "KpAdd"
	case KeyKpEnter:
		return "KpEnter"
	case KeyKpEqual:
		return "KpEqual"
	case KeyLeftShift:
		return "LeftShift"
	case KeyLeftControl:
		return "LeftControl"
	case KeyLeftAlt:
		return "LeftAlt"
	case KeyLeftSuper:
		return "LeftSuper"
	case KeyRightShift:
		return "RightShift"
	case KeyRightControl:
		return "RightControl"
	case KeyRightAlt:
		return "RightAlt"
	case KeyRightSuper:
		return "RightSuper"
	case KeyMenu:
		return "Menu"
	case KeyLast:
		return "Last"
	default:
		return "Unknown"
	}
}
