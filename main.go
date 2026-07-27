package main

import (
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) > 1 && (os.Args[1] == "-v" || os.Args[1] == "--version") {
		MessageOK(Version())
		return
	}

	cliFlags := GetCLIFlags()
	args := GetArgs(cliFlags)

	if len(args) == 0 {
		MessageError("usage: minimod <module-name>")
		os.Exit(1)
	}

	if err := Create(args[0], cliFlags); err != nil {
		MessageError(err.Error())
		os.Exit(1)
	}
}

func Create(packageName string, flags CLIFlags) error {

	pwd, err := os.Getwd()
	if err != nil {
		return err
	}

	absPath := filepath.Join(pwd, packageName)

	if err := CreateDir(packageName); err != nil {
		return err
	}

	cleanup := true

	defer func() {
		if cleanup {
			DeleteDir(packageName)
		}
	}()

	if err := GoModInit(packageName); err != nil {
		cleanup = true
		return err
	}

	if flags.Main {
		if err := WriteFile(filepath.Join(packageName, "main.go"), MainContent()); err != nil {
			cleanup = true
			return err
		}
	}

	if flags.GitIgnore {
		if err := WriteFile(filepath.Join(packageName, ".gitignore"), GitIgnoreContent()); err != nil {
			cleanup = true
			return err
		}
	}

	if flags.Readme {
		if err := WriteFile(filepath.Join(packageName, "README.md"), ReadmeContent(packageName)); err != nil {
			cleanup = true
			return err
		}
	}

	cleanup = false

	MessageOK("created module in", absPath)
	MessageOK("done!")

	return nil
}
