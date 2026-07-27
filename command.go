package main

import (
	"os/exec"
)

func GoModInit(packageName string) error {
	dir, err := GetAbsolute(packageName)
	if err != nil {
		return err
	}

	cmd := exec.Command("go", "mod", "init", packageName)
	cmd.Dir = dir

	out, err := cmd.CombinedOutput()
	if err != nil {
		return NewMinimodError(
			"error executing go mod init: " + err.Error() + "\n" + string(out),
		)
	}

	return nil
}
