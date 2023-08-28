package color

import (
	"fmt"
	"strings"
)

type Color string

const (
	Black   Color = "\x1b[30m"
	Red     Color = "\x1b[31m"
	Green   Color = "\x1b[32m"
	Yellow  Color = "\x1b[33m"
	Blue    Color = "\x1b[34m"
	Magenta Color = "\x1b[35m"
	Cyan    Color = "\x1b[36m"
	White   Color = "\x1b[37m"
	Reset   Color = "\x1b[0m"
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
