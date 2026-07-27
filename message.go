package main

import (
	"fmt"
	"strings"
)

const (
	reset = "\033[0m"
	bold  = "\033[1m"
	dim   = "\033[2m"

	white  = "\033[97m"
	gray   = "\033[90m"
	green  = "\033[32m"
	blue   = "\033[34m"
	yellow = "\033[33m"
	red    = "\033[31m"
)

func printMessage(levelColor, level string, msg ...any) {
	text := strings.TrimSpace(fmt.Sprint(msg...))

	fmt.Printf(
		"%sminmod%s %s[%s]%s %s%s%s\n",
		bold,
		reset,
		dim,
		level,
		reset,
		levelColor,
		text,
		reset,
	)
}

func MessageOK(msg ...any) {
	printMessage(green, "ok", msg...)
}

func MessageInfo(msg ...any) {
	printMessage(blue, "info", msg...)
}

func MessageWarn(msg ...any) {
	printMessage(yellow, "warn", msg...)
}

func MessageError(msg ...any) {
	printMessage(red, "error", msg...)
}
