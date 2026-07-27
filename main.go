package main

import (
	"os"
	"path/filepath"
)

func main() {

	if len(os.Args) < 2 {
		MessageError("usage: minimod <module-name>")
		os.Exit(1)
	}

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "-v", "--version":
			println(Version())
			return

		case "-h", "--help":
			println("just use minimod <module-name>")
			return
		}
	}

	packageName := os.Args[1]

	pwd, err := os.Getwd()
	if err != nil {
		MessageError(err.Error())
		os.Exit(1)
	}

	absPath := filepath.Join(pwd, packageName)

	if err := CreateDir(packageName); err != nil {
		MessageError(err.Error())
		os.Exit(1)
	}

	if err := GoModInit(packageName); err != nil {
		MessageError(err.Error())
		goto delete
	}

	if err := WriteFile(filepath.Join(packageName, "main.go"), MainContent()); err != nil {
		MessageError(err.Error())
		goto delete
	}

	if err := WriteFile(filepath.Join(packageName, ".gitignore"), GitIgnoreContent()); err != nil {
		MessageError(err.Error())
		goto delete
	}

	if err := WriteFile(filepath.Join(packageName, "README.md"), ReadmeContent(packageName)); err != nil {
		MessageError(err.Error())
		goto delete
	}

	MessageOK("created module in", absPath)
	MessageOK("done!")

delete:
	DeleteDir(packageName)
	os.Exit(1)
}
