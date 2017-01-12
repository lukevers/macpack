package main

import (
	"errors"
	"log"
	"os"
	"path/filepath"
)

func build(verbose bool) error {
	args := []string{"build", "-ldflags", "-s"}

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

	fi, err := os.Stat(currentExecName)
	if err != nil {
		return err
	}

	if fi.IsDir() {
		return errors.New("%v should be a binary: dir")
	}
	return os.Rename(currentExecName, execName)
}
