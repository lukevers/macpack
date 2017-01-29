package main

import (
	"fmt"

	"github.com/murlokswarm/config"
)

var (
	cfg = defaultConfig()
)

func main() {
	config.ConfigName = "mac.json"
	config.CreateIfNotExists = true
	config.Commands = commandsString()

	cmds := config.Load(&cfg)
	if len(cmds) != 1 {
		fmt.Printf("\033[91mInvalid command: %v. use macpack -h for help\n\033[00m", cmds)
		return
	}
	if err := cfg.check(); err != nil {
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

	if err := createPlist(cfg); err != nil {
		return err
	}

	if cfg.Sandbox {
		if len(cfg.SignID) == 0 {
			fmt.Println("\033[1mWarning:\033[00m")
			fmt.Println("  Sandbox mode requires \033[1msign-id\033[00m flag to be set.")
			fmt.Println("  Valid ids can be seen with: \033[1msecurity find-identity -v -p codesigning\033[00m")
			fmt.Println("  See the 'Obtaining Your Signing Identities' section from ")
			fmt.Println("  \033[94mhttps://developer.apple.com/library/content/documentation/Security/Conceptual/CodeSigningGuide/Procedures/Procedures.html\033[00m")
			fmt.Println("  for creating one.")
			return nil
		}

		if err := createEntitlements(cfg.Capabilities); err != nil {
			return err
		}
		defer deleteEntitlements()

		if err := signPackage(cfg); err != nil {
			return err
		}

		if err := signCheck(cfg); err != nil {
			return err
		}
	}

	if cfg.Store {
		if !cfg.Sandbox {
			fmt.Println("\033[1mWarning:\033[00m")
			fmt.Println("  Mac App Store app must run in sandbox mode.")
			fmt.Println("  Use \033[1m-sandbox\033[00m flag.")
			return nil
		}
		if err := packageForAppStore(cfg); err != nil {
			return err
		}
	}

	fmt.Println("\033[92mBuild success!\033[00m")
	return nil
}
