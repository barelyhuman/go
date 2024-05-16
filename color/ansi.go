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
var DimCode = "\033[2m"

type ColorString struct {
	value *strings.Builder
}

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
		DimCode = ""
	}

}

func (cs *ColorString) initCheck() {
	if cs.value == nil {
		cs.value = &strings.Builder{}
	}
}

// Write a string to the ColorString instance. This is to add in normal or cloning existing ColorString instances
func (cs *ColorString) Write(msg string) *ColorString {
	cs.initCheck()
	cs.value.WriteString(msg)
	return cs
}

func colorBuilderFunc(cs *ColorString, colorCode string, msg string) *ColorString {
	cs.initCheck()
	cs.value.WriteString(ResetCode + colorCode + msg + ResetCode)
	return cs
}

// Returns the completed string from a color string instance
func (cs *ColorString) String() string {
	cs.initCheck()
	return strings.ReplaceAll(cs.value.String(), ResetCode+ResetCode, ResetCode)
}

// Reset the ansi color of the given message
func (cs *ColorString) Reset(msg string) *ColorString {
	return colorBuilderFunc(cs, ResetCode, msg)
}

// Wrap the input msg into the ANSI blue
func (cs *ColorString) Blue(msg string) *ColorString {
	return colorBuilderFunc(cs, BlueCode, msg)
}

// Wrap the input msg into the ANSI red
func (cs *ColorString) Red(msg string) *ColorString {
	return colorBuilderFunc(cs, RedCode, msg)
}

// Wrap the input msg into the ANSI green
func (cs *ColorString) Green(msg string) *ColorString {
	return colorBuilderFunc(cs, GreenCode, msg)
}

// Wrap the input msg into the ANSI yellow
func (cs *ColorString) Yellow(msg string) *ColorString {
	return colorBuilderFunc(cs, YellowCode, msg)
}

// Wrap the input msg into the ANSI purple
func (cs *ColorString) Purple(msg string) *ColorString {
	return colorBuilderFunc(cs, PurpleCode, msg)
}

// Wrap the input msg into the ANSI cyan
func (cs *ColorString) Cyan(msg string) *ColorString {
	return colorBuilderFunc(cs, CyanCode, msg)
}

// Wrap the input msg into the ANSI gray
func (cs *ColorString) Gray(msg string) *ColorString {
	return colorBuilderFunc(cs, GrayCode, msg)
}

// Wrap the input msg into the ANSI white
func (cs *ColorString) White(msg string) *ColorString {
	return colorBuilderFunc(cs, WhiteCode, msg)
}

func (cs *ColorString) Dim(msg string) *ColorString {
	return colorBuilderFunc(cs, DimCode, msg)
}
