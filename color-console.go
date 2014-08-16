package lib

import (
	"fmt"
)

type Color int

const (
	// No change of color
	None    = 0
	Black   = 1
	Red     = 2
	Green   = 3
	Yellow  = 4
	Blue    = 5
	Magenta = 6
	Cyan    = 7
	White   = 8
)

const (
	// No change of color
	NoBg      = 0
	BlackBg   = Black << 4
	RedBg     = Red << 4
	GreenBg   = Green << 4
	YellowBg  = Yellow << 4
	BlueBg    = Blue << 4
	MagentaBg = Magenta << 4
	CyanBg    = Cyan << 4
	WhiteBg   = White << 4
)

const (
	// Normal     = 0
	Bold       = 1 << 8
	Bright     = 1 << 8
	Dim        = 2 << 8
	Underlined = 4 << 8
	Blink      = 5 << 8
	Inverted   = 7 << 8
	Hidden     = 8 << 8
)

func getColorEscapeSequence(c Color) string {

	fg := c & 0xF
	bg := (c & 0xF0) >> 4
	op := (c & 0xF00) >> 8

	s := fmt.Sprintf("[%d", op)

	if fg != 0 {
		s += fmt.Sprintf(";%d", (30 + fg - 1))
	} else {
		s += fmt.Sprintf(";%d", 39)
	}

	if bg != 0 {
		s += fmt.Sprintf(";%d", (40 + bg - 1))
	} else {
		s += fmt.Sprintf(";%d", 49)
	}

	return fmt.Sprintf("\x1b%sm", s)
}

func ColorPrintf(c Color, format string, a ...interface{}) (int, error) {
	return fmt.Printf("%s%s%s", getColorEscapeSequence(c), fmt.Sprintf(format, a...), "\x1b[0m")
}

func ColorPrintln(c Color, a ...interface{}) (int, error) {
	return fmt.Printf("%s%s%s", getColorEscapeSequence(c), fmt.Sprintln(a...), "\x1b[0m")
}
