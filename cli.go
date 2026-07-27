package main

var version string = "1.0.0"

type CliFlags struct {
	GitIgnore bool
	License   bool
	Readme    bool
}

func Version() string {
	return "minimod: " + version
}
