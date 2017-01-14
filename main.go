package main

import (
	"fmt"

	"github.com/segmentio/conf"
)

var (
	config = defaultConfig()
)

func main() {
	const cmdFormat = "\033[90muse\033[00m gomac [-h] [-help] [options...] [build | sass]"
	cmds := conf.Load(&config)
	if len(cmds) != 1 {
		fmt.Println("\033[91mbad cmd format\033[00m")
		fmt.Println(cmdFormat)
		return
	}
	if err := config.check(); err != nil {
		fmt.Printf("\033[91m%v\033[00m\n", err)
	}

	switch cmd := cmds[0]; cmd {
	case "build":
		if err := build(); err != nil {
			fmt.Println(err)
			return
		}

	case "sass":
		if err := launchSass(); err != nil {
			fmt.Println(err)
			return
		}

	default:
		fmt.Println("\033[91munknown cmd:\033[00m", cmd)
		fmt.Println(cmdFormat)
		return
	}
}

func build() error {
	if err := gobuild(); err != nil {
		return err
	}

	if err := createPackage(); err != nil {
		return err
	}

	if err := createExec(); err != nil {
		return err
	}

	if err := createResources(); err != nil {
		return err
	}

	if err := syncResources(); err != nil {
		return err
	}

	if err := generateIcon(); err != nil {
		return err
	}

	if err := createPlist(); err != nil {
		return err
	}
	return nil
}
