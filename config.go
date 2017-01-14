package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/murlokswarm/errors"
)

// Config represents the configuration of the app to be packaged.
type Config struct {
	// Context options.
	Verbose   bool `conf:"v,omitempty" help:"Verbose mode"`
	Overwrite bool `conf:"overwrite"   help:"Resources directory will be entirely copied, overwriting existing files"`

	// Build options.
	Name             string       `conf:"name"              help:"Package name, menu bar/dock display name"`
	ID               string       `conf:"id"                help:"UTI representing the app"`
	Version          string       `conf:"version"           help:"Version of the app"`
	Icon             string       `conf:"icon"              help:"The app icon as .png file. Provide a big one, required icon sizes will be auto generated"`
	DevRegion        string       `conf:"dev-region"        help:"Development region"`
	DeploymentTarget string       `conf:"deployment-target" help:"MacOS version"`
	Copyright        string       `conf:"copyright"         help:"Human readable copyright"`
	Role             string       `conf:"role"              help:"Application role: Editor|Viewer|Shell|None"`
	Sandbox          bool         `conf:"sandbox"           help:"Defines if the app will run in sandbox mode"`
	Capabilities     capabilities `conf:"capabilities"      help:"Capabilities" conf:"App capabilities. Required by the Mac App Store. Requires sandbox mode"`
	SupportedFiles   []string     `conf:"supported-files"   help:"List of UTI representing the file types the app can open"`
}

type capabilities struct {
	Network    networkCap    `conf:"network"     help:"Network capabilities"`
	Hardware   harwareCap    `conf:"hardware"    help:"Hardware capabilities"`
	AppData    appDataCap    `conf:"app-data"    help:"Application data capabilities"`
	FileAccess fileAccessCap `conf:"file-access" help:"File access capabilities"`
}

type networkCap struct {
	In  bool `conf:"in"  help:"Incoming connections (Server)"`
	Out bool `conf:"out" help:"Outgoing connections (Client)"`
}

type harwareCap struct {
	Camera     bool `conf:"camera"     help:"Use of camera"`
	Microphone bool `conf:"microphone" help:"Use of microphone"`
	USB        bool `conf:"usb"        help:"Use of USB"`
	Printing   bool `conf:"printing"   help:"Use of printer"`
	Bluetooth  bool `conf:"bluetooth"  help:"Use of bluetooth"`
}

type appDataCap struct {
	Contacts bool `conf:"contacts" help:"Access contacts"`
	Location bool `conf:"location" help:"Access location"`
	Calendar bool `conf:"calendar" help:"Access calendar"`
}

type fileAccessCap struct {
	UserSelected fileAccess `conf:"user-selected" help:"Access files from file pickers"`
	Downloads    fileAccess `conf:"downloads"     help:"Access default Downloads directory"`
	Pictures     fileAccess `conf:"pictures"      help:"Access default Pictures directory"`
	Music        fileAccess `conf:"music"         help:"Access default Music directory"`
	Movies       fileAccess `conf:"movies"        help:"Access default Movies directory"`
}

type fileAccess string

const (
	fileNoAccess        fileAccess = ""
	fileReadAccess      fileAccess = "read-only"
	fileReadWriteAccess fileAccess = "read-write"
)

func defaultConfig() Config {
	wd, err := os.Getwd()
	if err != nil {
		log.Panic(errors.New(err))
	}
	name := filepath.Base(wd)

	return Config{
		Name:             name,
		ID:               fmt.Sprintf("%v.%v", os.Getenv("USER"), name),
		Version:          "1.0.0.0",
		DevRegion:        "en",
		DeploymentTarget: "10.12",
		Copyright:        fmt.Sprintf("Copyright © 2017 %v. All rights reserved", os.Getenv("USER")),
		Role:             "None",
		Sandbox:          true,
		Capabilities: capabilities{
			Network: networkCap{
				Out: true,
			},
		},
		SupportedFiles: []string{},
	}
}

func (c Config) check() error {
	validName := regexp.MustCompile(`^([A-Za-z0-9_]|[\-])+$`)
	validVersion := regexp.MustCompile(`^[0-9]+([\.][0-9]+){3}$`)
	validUTI := regexp.MustCompile(`^([A-Za-z0-9]|[\-]|[\.])+$`)
	validMinVersion := regexp.MustCompile(`^[0-9]+[\.][0-9]+$`)

	if !validName.MatchString(c.Name) {
		return fmt.Errorf("name from config must contain alphanumeric characters, '_' or '-' : %v", c.Name)
	}

	if !validUTI.MatchString(c.ID) {
		return fmt.Errorf("id from config must contain alphanumeric characters, '-' or '.' : %v", c.ID)
	}

	if !validVersion.MatchString(c.Version) {
		return fmt.Errorf("version from config must follow the pattern x.x.x.x where x is a non negative number: %v", c.Version)
	}

	if !validMinVersion.MatchString(c.DeploymentTarget) {
		return fmt.Errorf("deployment-target from config must follow the pattern x.x where x is a non negative number: %v", c.DeploymentTarget)
	}

	macOSDeploymentTarget := strings.Split(c.DeploymentTarget, "")
	if maj, _ := strconv.Atoi(macOSDeploymentTarget[0]); maj < 10 {
		return fmt.Errorf("major revision in deployment-target from config cannot be under 10:%v", maj)
	}
	if min, _ := strconv.Atoi(macOSDeploymentTarget[1]); min < 11 {
		return fmt.Errorf("minor revision in deployment-target from config cannot be under 11:%v", min)
	}

	if role := c.Role; role != "Editor" && role != "Viewer" && role != "Shell" && role != "None" {
		return fmt.Errorf("role from config should be Editor | Viewer | Shell | None:%v", c.Role)
	}

	for _, uti := range c.SupportedFiles {
		if !validUTI.MatchString(uti) {
			return fmt.Errorf("supported-files from config contains a non valid UTI: %v", uti)
		}
	}
	return nil
}
