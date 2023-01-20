package color_test

import (
	"testing"

	"github.com/barelyhuman/go/color"
)

func TestWrite(t *testing.T) {
	testableString := "Hello World"
	cs := color.ColorString{}
	cs.Write("Hello World")
	val := cs.String()

	if val != testableString {
		t.Fail()
	}
}

func TestColorBuilder(t *testing.T) {
	testingColor := color.RedCode
	testableString := "Hello World"
	cs := color.ColorString{}
	cs.Red("Hello World")
	val := cs.String()

	if val != color.ResetCode+testingColor+testableString+color.ResetCode {
		t.Fail()
	}
}
