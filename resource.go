package main

import "path/filepath"
import "os"

func syncResources(conf Config, force bool, verbose bool) error {
	cssName := filepath.Join("resources", "css")

	if err := os.MkdirAll(cssName, os.ModeDir|0755); err != nil {
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
