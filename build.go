package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

func gobuild() error {
	args := []string{"build", "-ldflags", "-s"}
	return execCmd("go", args...)
}

func createPackage() error {
	name := cfg.Name + ".app"
	macOSName := filepath.Join(name, "Contents", "MacOS")
	resourcesName := filepath.Join(name, "Contents", "Resources")

	if err := os.MkdirAll(macOSName, os.ModeDir|0755); err != nil {
		return err
	}
	return os.MkdirAll(resourcesName, os.ModeDir|0755)
}

func removePackage() {
	name := cfg.Name + ".app"
	os.RemoveAll(name)
}

func createExec() error {
	execName := filepath.Join(cfg.Name+".app", "Contents", "MacOS", cfg.Name)
	currentExecName := goExecName()

	fi, err := os.Stat(currentExecName)
	if err != nil {
		return err
	}
	if fi.IsDir() {
		return errors.New("%v should be a binary: dir")
	}
	return os.Rename(currentExecName, execName)
}

func goExecName() string {
	wd, err := os.Getwd()
	if err != nil {
		log.Panic(err)
	}
	return filepath.Base(wd)
}

func createResources() error {
	cssName := filepath.Join("resources", "css")
	return os.MkdirAll(cssName, os.ModeDir|0755)
}

func syncResources() error {
	cssName := filepath.Join("resources", "css")
	if err := os.MkdirAll(cssName, os.ModeDir|0755); err != nil {
		return err
	}

	resourcesName := filepath.Join(cfg.Name+".app", "Contents", "Resources")
	args := []string{
		"resources/",
		resourcesName,
		"-r",
		"--delete",
		"--update",
	}
	return execCmd("rsync", args...)
}
