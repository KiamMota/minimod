package main

import (
	"fmt"
	"strings"
)

const (
	reset = "\033[0m"
	bold  = "\033[1m"
	white = "\033[97m"

	bgBlue   = "\033[44m"
	bgGreen  = "\033[42m"
	bgYellow = "\033[43m"
	bgRed    = "\033[41m"

	gray = "\033[90m"
)

func log(bg string, msg ...any) {
	fmt.Printf(
		"%s%s%sminimod%s %s|%s %s\n",
		bold,
		bg,
		white,
		reset,
		gray,
		reset,
		strings.TrimSpace(fmt.Sprint(msg...)),
	)
}

func MessageOK(msg ...any) {
	log(bgGreen, msg...)
}

func MessageInfo(msg ...any) {
	log(bgBlue, msg...)
}

func MessageWarn(msg ...any) {
	log(bgYellow, msg...)
}

func MessageError(msg ...any) {
	log(bgRed, msg...)
}
