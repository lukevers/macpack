package main

import (
	"os"
	"path/filepath"
)

func createPackage(conf Config) error {
	name := conf.Name + ".app"
	macOSName := filepath.Join(name, "Contents", "MacOS")
	resourcesName := filepath.Join(name, "Contents", "Resources")

	if err := os.MkdirAll(macOSName, os.ModeDir|0755); err != nil {
		return err
	}
	return os.MkdirAll(resourcesName, os.ModeDir|0755)
}
