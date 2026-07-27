package main

import "errors"

func NewMinimodError(msg string) error {
	return errors.New("minmod error: " + msg)
}
