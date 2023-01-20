package color

import (
	"runtime"
	"strings"
)

var ResetCode = "\033[0m"
var RedCode = "\033[31m"
var GreenCode = "\033[32m"
var YellowCode = "\033[33m"
var BlueCode = "\033[34m"
var PurpleCode = "\033[35m"
var CyanCode = "\033[36m"
var GrayCode = "\033[37m"
var WhiteCode = "\033[97m"

func init() {
	if runtime.GOOS == "windows" {
		ResetCode = ""
		RedCode = ""
		GreenCode = ""
		YellowCode = ""
		BlueCode = ""
		PurpleCode = ""
		CyanCode = ""
		GrayCode = ""
		WhiteCode = ""
	}
}

type ColorString struct {
	value *strings.Builder
}

// Write a string to the ColorString instance. This is to add in normal or cloning existing ColorString instances
func (cs *ColorString) Write(msg string) *ColorString {
	cs.value.WriteString(msg)
	return cs
}

// Returns the completed string from a color string instance
func (cs *ColorString) String() string {
	return cs.value.String()
}

// Reset the ansi color of the given message
func (cs *ColorString) Reset(msg string) *ColorString {
	cs.value.WriteString(ResetCode + msg + ResetCode)
	return cs
}

// Wrap the inptu msg into the ANSI red
func (cs *ColorString) Red(msg string) *ColorString {
	cs.value.WriteString(ResetCode + msg + RedCode + ResetCode)
	return cs
}

// Wrap the inptu msg into the ANSI green
func (cs *ColorString) Green(msg string) *ColorString {
	cs.value.WriteString(ResetCode + msg + GreenCode + ResetCode)
	return cs
}

// Wrap the inptu msg into the ANSI yellow
func (cs *ColorString) Yellow(msg string) *ColorString {
	cs.value.WriteString(ResetCode + msg + YellowCode + ResetCode)
	return cs
}

// Wrap the inptu msg into the ANSI blue
func (cs *ColorString) Blue(msg string) *ColorString {
	cs.value.WriteString(ResetCode + msg + BlueCode + ResetCode)
	return cs
}

// Wrap the inptu msg into the ANSI purple
func (cs *ColorString) Purple(msg string) *ColorString {
	cs.value.WriteString(ResetCode + msg + PurpleCode + ResetCode)
	return cs
}

// Wrap the inptu msg into the ANSI cyan
func (cs *ColorString) Cyan(msg string) *ColorString {
	cs.value.WriteString(ResetCode + msg + CyanCode + ResetCode)
	return cs
}

// Wrap the inptu msg into the ANSI gray
func (cs *ColorString) Gray(msg string) *ColorString {
	cs.value.WriteString(ResetCode + msg + GrayCode + ResetCode)
	return cs
}

// Wrap the inptu msg into the ANSI white
func (cs *ColorString) White(msg string) *ColorString {
	cs.value.WriteString(ResetCode + msg + WhiteCode + ResetCode)
	return cs
}
