package main

import (
	"os"
	"path/filepath"
	"strings"
)

func message(msg ...string) string {
	return "minimod: " + strings.Join(msg, " ")
}

func main() {
	if len(os.Args) < 2 {
		println(message("usage: minimod <module-name>"))
		os.Exit(1)
	}

	packageName := os.Args[1]

	pwd, err := os.Getwd()
	if err != nil {
		println(message(err.Error()))
		os.Exit(1)
	}

	absPath := filepath.Join(pwd, packageName)

	if err := CreateDir(packageName); err != nil {
		println(message(err.Error()))
		os.Exit(1)
	}

	if err := GoModInit(packageName); err != nil {
		println(message(err.Error()))
		os.Exit(1)
	}

	if err := WriteFile(filepath.Join(packageName, "main.go"), MainContent()); err != nil {
		println(message(err.Error()))
		os.Exit(1)
	}

	if err := WriteFile(filepath.Join(packageName, ".gitignore"), GitIgnoreContent()); err != nil {
		println(message(err.Error()))
		os.Exit(1)
	}

	if err := WriteFile(filepath.Join(packageName, "README.md"), ReadmeContent(packageName)); err != nil {
		println(message(err.Error()))
		os.Exit(1)
	}

	println(message("created module in", absPath))
	println(message("done!"))
}
