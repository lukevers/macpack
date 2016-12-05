package main

import (
	"log"
	"os"
	"path/filepath"
)

func build(verbose bool) error {
	args := []string{"build"}

	if verbose {
		args = append(args, "-v")
	}

	return execCmd("go", args...)
}

func createExec(conf Config) error {
	execName := filepath.Join(conf.Name+".app", "Contents", "MacOS", conf.Name)

	wd, err := os.Getwd()
	if err != nil {
		log.Panic(err)
	}

	currentExecName := filepath.Base(wd)
	return os.Rename(currentExecName, execName)
}
