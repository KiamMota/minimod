package main

import "flag"

var version string = "1.0.0"

type CLIFlags struct {
	GitIgnore bool
	License   string
	Readme    bool
	Main      bool
}

func Version() string {
	return "minimod: " + version
}

func GetCLIFlags() CLIFlags {
	flags := CLIFlags{
		GitIgnore: true,
		Main:      true,
		License:   "MIT",
	}

	noGitIgnore := flag.Bool(
		"no-gitignore",
		false,
		"do not create .gitignore",
	)

	noMain := flag.Bool(
		"no-main",
		false,
		"do not create main.go",
	)

	readme := flag.Bool(
		"readme",
		false,
		"create README.md",
	)

	license := flag.String(
		"license",
		"MIT",
		"license type",
	)

	flag.Parse()

	flags.GitIgnore = !*noGitIgnore
	flags.Main = !*noMain
	flags.Readme = *readme
	flags.License = *license

	return flags
}

func GetArgs(flags CLIFlags) []string {
	return flag.Args()
}
