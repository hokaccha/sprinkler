package utils

import (
	"fmt"
)

type ColorCode int

const (
	black ColorCode = iota + 30
	red
	green
	yellow
	blue
	magenta
	cyan
	white
)

func colorize(c ColorCode, str string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", c, str)
}

func Black(str string) string {
	return colorize(black, str)
}

func Red(str string) string {
	return colorize(red, str)
}

func Green(str string) string {
	return colorize(green, str)
}

func Yellow(str string) string {
	return colorize(yellow, str)
}

func Blue(str string) string {
	return colorize(blue, str)
}

func Magenta(str string) string {
	return colorize(magenta, str)
}

func Cyan(str string) string {
	return colorize(cyan, str)
}

func White(str string) string {
	return colorize(white, str)
}
