package main

import "path/filepath"
import "os"

func launchSass() error {
	scssName := filepath.Join("resources", "scss")

	if err := os.MkdirAll(scssName, os.ModeDir|0755); err != nil {
		return err
	}

	args := []string{
		"--watch",
		"resources/scss/:resources/css/",
	}
	execCmd("sass", args...)
	return nil
}
