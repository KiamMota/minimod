package main

import "os/exec"

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
			"Error executing module initialization in Go:" + err.Error() + string(out),
		)

		if e := DeleteDir(packageName); e != nil {
			NewMinimodError(e.Error())
		}

		return err
	}

	return nil
}
