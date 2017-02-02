package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func launchSass() error {
	scssName := filepath.Join("resources", "scss")
	if err := os.MkdirAll(scssName, os.ModeDir|0755); err != nil {
		return err
	}

	args := []string{
		"--watch",
		"resources/scss/:resources/css/",
	}
	if err := execCmd("sass", args...); err != nil {
		return fmt.Errorf("\033[91m%v\033[00m\ninstall it with \033[94msudo gem install sass\033[00m", err)
	}
	return nil
}
