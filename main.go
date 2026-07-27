package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"time"
	"unicode/utf8"

	"golang.org/x/mod/modfile"
)

func main() {
	if len(os.Args) > 1 && (os.Args[1] == "-v" || os.Args[1] == "--version") {
		println(Version())
		return
	}

	if len(os.Args) > 1 && (os.Args[1] == "--go" || os.Args[1] == "-g") {
		v, e := GoVersion()
		if e != nil {
			MessageError(e.Error())
			return
		}
		println(v)
		return
	}

	if len(os.Args) > 1 && (os.Args[1] == "-i" || os.Args[1] == "--info") {
		v, e := ModInfo()
		if e != nil {
			MessageError(e.Error())
			return
		}
		println(v)
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

var validModulePath = regexp.MustCompile(`^[a-zA-Z0-9._/-]+$`)

func GoVersion() (string, error) {
	data, err := os.ReadFile("go.mod")
	if err != nil {
		return "", err
	}

	file, err := modfile.Parse("go.mod", data, nil)
	if err != nil {
		return "", err
	}

	return file.Go.Version, nil
}

func ModInfo() (string, error) {
	data, err := os.ReadFile("go.mod")
	if err != nil {
		return "", err
	}

	mod, err := modfile.Parse("go.mod", data, nil)
	if err != nil {
		return "", err
	}

	if mod.Module == nil {
		return "", fmt.Errorf("module not found")
	}

	return mod.Module.Mod.Path, nil
}

func Create(packageName string, flags CLIFlags) error {

	if !utf8.ValidString(packageName) {
		return NewMinimodError("package name is not valid utf8!")
	}

	// Agora o MatchString vai funcionar normalmente
	if !validModulePath.MatchString(packageName) {
		return NewMinimodError("package name contains invalid characters! Use only alphanumeric characters, hyphens, underscores, dots or slashes.")
	}

	start := time.Now()
	VerboseLog(flags.Verbose, "starting create:", packageName)

	pwd, err := os.Getwd()
	if err != nil {
		return err
	}

	absPath := filepath.Join(pwd, packageName)
	VerboseLog(flags.Verbose, "working directory:", pwd)
	VerboseLog(flags.Verbose, "target path:", absPath)

	VerboseLog(flags.Verbose, "creating directory")
	if err := CreateDir(packageName); err != nil {
		return err
	}

	cleanup := true
	defer func() {
		if cleanup {
			VerboseLog(flags.Verbose, "cleanup:", packageName)
			_ = DeleteDir(packageName)
		}
	}()

	VerboseLog(flags.Verbose, "initializing go.mod")
	t0 := time.Now()
	if err := GoModInit(packageName); err != nil {
		return err
	}
	VerboseLog(flags.Verbose, "go.mod initialized in", time.Since(t0))

	if flags.Main {
		VerboseLog(flags.Verbose, "writing main.go")
		t0 = time.Now()
		if err := WriteFile(filepath.Join(packageName, "main.go"), MainContent()); err != nil {
			return err
		}
		VerboseLog(flags.Verbose, "main.go written in", time.Since(t0))
	}

	if flags.GitIgnore {
		VerboseLog(flags.Verbose, "writing .gitignore")
		t0 = time.Now()
		if err := WriteFile(filepath.Join(packageName, ".gitignore"), GitIgnoreContent()); err != nil {
			return err
		}
		VerboseLog(flags.Verbose, ".gitignore written in", time.Since(t0))
	}

	if flags.Readme {
		VerboseLog(flags.Verbose, "writing README.md")
		t0 = time.Now()
		if err := WriteFile(filepath.Join(packageName, "README.md"), ReadmeContent(packageName)); err != nil {
			return err
		}
		VerboseLog(flags.Verbose, "README.md written in", time.Since(t0))
	}

	cleanup = false
	VerboseLog(flags.Verbose, "finished in", time.Since(start))

	MessageOK("created module in", absPath)
	MessageOK("done!")

	if flags.Duration {
		MessageOK("duration:", time.Since(start))
	}

	return nil
}
