package color

import (
	"fmt"
	"strings"
)

type Color string

const (
	Black   = "\x1b[40m"
	Red     = "\x1b[41m"
	Green   = "\x1b[42m"
	Yellow  = "\x1b[43m"
	Blue    = "\x1b[44m"
	Magenta = "\x1b[45m"
	Cyan    = "\x1b[46m"
	White   = "\x1b[47m"
	Reset   = "\x1b[0m"
)

func GetColor(str string) Color {
	switch strings.ToLower(str) {
	case "black":
		return Black
	case "red":
		return Red
	case "green":
		return Green
	case "yellow":
		return Yellow
	case "blue":
		return Blue
	case "magenta":
		return Magenta
	case "cyan":
		return Cyan
	case "white":
		return White
	default:
		return Reset // Default to reset color
	}
}

func Colorf(color Color, format string, a ...interface{}) string {
	return fmt.Sprintf("%s%s%s", color, fmt.Sprintf(format, a...), Reset)
}

func Colorln(color Color, a ...interface{}) string {
	return fmt.Sprintf("%s%s%s\n", color, fmt.Sprint(a...), Reset)
}
