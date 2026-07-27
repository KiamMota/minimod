package main

import (
	"flag"
	"os"
	"strings"
)

var version string = "1.0.0"

// cliArgs guarda os argumentos posicionais (ex: nome do módulo) extraídos
// pela reorderArgs, já que agora eles podem vir antes OU depois das flags.
var cliArgs []string

type CLIFlags struct {
	GitIgnore bool
	License   string
	Readme    bool
	Main      bool
}

func Version() string {
	return "minimod: " + version
}

// valueFlagNames são as flags que recebem um valor separado (ex: -license MIT)
var valueFlagNames = map[string]bool{
	"license": true,
}

// flagBaseName remove "-" ou "--" do início e retorna o nome puro da flag
// e, se houver "=", o nome antes do "=" (para suportar -license=MIT)
func flagBaseName(arg string) (name string, hasEquals bool) {
	trimmed := strings.TrimLeft(arg, "-")
	if idx := strings.Index(trimmed, "="); idx != -1 {
		return trimmed[:idx], true
	}
	return trimmed, false
}

// reorderArgs separa os argumentos em posicionais e flags, permitindo que
// flags apareçam em qualquer posição na linha de comando (antes ou depois
// dos argumentos posicionais), algo que o pacote "flag" da stdlib não
// suporta nativamente (ele para de parsear flags no primeiro não-flag).
func reorderArgs(args []string) (reordered []string, positional []string) {
	i := 0
	for i < len(args) {
		arg := args[i]

		isFlag := strings.HasPrefix(arg, "-")
		if !isFlag {
			positional = append(positional, arg)
			i++
			continue
		}

		name, hasEquals := flagBaseName(arg)

		reordered = append(reordered, arg)

		// Se a flag espera um valor separado (ex: -license MIT) e não veio
		// no formato -license=MIT, o próximo token pertence a essa flag.
		if valueFlagNames[name] && !hasEquals {
			if i+1 < len(args) {
				reordered = append(reordered, args[i+1])
				i += 2
				continue
			}
		}

		i++
	}

	return reordered, positional
}

func GetCLIFlags() CLIFlags {
	flags := CLIFlags{
		GitIgnore: true,
		Main:      true,
		License:   "MIT",
	}

	fs := flag.NewFlagSet("minimod", flag.ExitOnError)

	noGitIgnore := fs.Bool(
		"no-gitignore",
		false,
		"do not create .gitignore",
	)

	noMain := fs.Bool(
		"no-main",
		false,
		"do not create main.go",
	)

	readme := fs.Bool(
		"readme",
		false,
		"create README.md",
	)

	license := fs.String(
		"license",
		"MIT",
		"license type",
	)

	reorderedFlags, positional := reorderArgs(os.Args[1:])
	fs.Parse(reorderedFlags)

	cliArgs = positional

	flags.GitIgnore = !*noGitIgnore
	flags.Main = !*noMain
	flags.Readme = *readme
	flags.License = *license

	return flags
}

func GetArgs(flags CLIFlags) []string {
	return cliArgs
}
