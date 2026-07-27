package main

import (
	"os"
	"path/filepath"
)

func FileExists(path string) bool {
	_, e := os.Stat(path)
	if e != nil {
		return false
	}
	return true
}

func GetAbsolute(path string) (string, error) {
	abs, e := filepath.Abs(path)
	if e != nil {
		NewMinimodError(e.Error())
	}
	return abs, nil
}

func CreateDir(packageName string) error {
	if FileExists(packageName) {
		return NewMinimodError("File already exists!")
	}
	e := os.Mkdir(packageName, os.FileMode(os.O_RDWR))
	if e != nil {
		return NewMinimodError(e.Error())
	}
	return nil

}

func WriteFile(path string, content string) error {
	if FileExists(path) {
		return NewMinimodError("File already exists!")
	}
	e := os.WriteFile(path, []byte(content), os.FileMode(os.O_RDWR))
	if e != nil {
		return NewMinimodError(e.Error())
	}
	return nil
}

func DeleteDir(path string) error {
	_, e := os.Stat(path)
	if e != nil {
		return NewMinimodError("file does not exist")
	}

	e = os.RemoveAll(path)
	if e != nil {
		// Só chama e.Error() se realmente existir um erro
		return NewMinimodError(e.Error())
	}

	return nil // Retorna nil indicando que tudo deu certo
}
