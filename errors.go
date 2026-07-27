package main

import "errors"

func NewMinimodError(msg string) error {
	return errors.New("minimod error: " + msg)
}
