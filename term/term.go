package term

import (
	"fmt"

	t "github.com/buger/goterm"
)

// HideCursor hides the cursor.
func HideCursor() {
	fmt.Printf("\033[?25l")
}

// ShowCursor shows the cursor.
func ShowCursor() {
	fmt.Printf("\033[?25h")
}

// ClearAll clears the screen.
func ClearAll() {
	fmt.Printf("\033[2J")
	t.MoveTo("", 1, 1)
}
