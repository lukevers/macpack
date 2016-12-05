package main

import "path/filepath"
import "os"

func syncResources(conf Config, force bool, verbose bool) error {
	if _, err := os.Stat("resources"); err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	resourcesName := filepath.Join(conf.Name+".app", "Contents", "Resources")
	args := []string{
		"resources/",
		resourcesName,
		"-r",
		"--delete",
	}

	if !force {
		args = append(args, "--update")
	}

	if verbose {
		args = append(args, "-v")
	}

	return execCmd("rsync", args...)
}
