package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if err := Create(); err != nil {
		MessageError(err.Error())
		os.Exit(1)
	}
}

func Create() error {
	args := os.Args[1:]

	if len(args) == 0 {
		return fmt.Errorf("usage: minimod <module-name>")
	}

	switch args[0] {
	case "-v", "--version":
		fmt.Println(Version())
		return nil

	case "-h", "--help":
		fmt.Println("usage: minimod <module-name>")
		return nil
	}

	packageName := args[0]

	pwd, err := os.Getwd()
	if err != nil {
		return err
	}

	absPath := filepath.Join(pwd, packageName)

	if err := CreateDir(packageName); err != nil {
		return err
	}

	cleanup := false
	defer func() {
		if cleanup {
			DeleteDir(packageName)
		}
	}()

	if err := GoModInit(packageName); err != nil {
		cleanup = true
		return err
	}

	if err := WriteFile(filepath.Join(packageName, "main.go"), MainContent()); err != nil {
		cleanup = true
		return err
	}

	if err := WriteFile(filepath.Join(packageName, ".gitignore"), GitIgnoreContent()); err != nil {
		cleanup = true
		return err
	}

	if err := WriteFile(filepath.Join(packageName, "README.md"), ReadmeContent(packageName)); err != nil {
		cleanup = true
		return err
	}

	MessageOK("created module in", absPath)
	MessageOK("done!")

	return nil
}
