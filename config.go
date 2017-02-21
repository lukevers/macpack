package main

import (
	"bytes"
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
	Name             string       `json:"name"              help:"Package name, menu bar/dock display name."`
	ID               string       `json:"id"                help:"UTI representing the app."`
	Version          string       `json:"version"           help:"Version of the app (minified form eg 1.42)."`
	BuildNumber      int          `json:"build-number"      help:"Build number."`
	Icon             string       `json:"icon"              help:"The app icon as .png file. Provide a big one! Other required icon sizes will be auto generated."`
	DevRegion        string       `json:"dev-region"        help:"Development region."`
	DeploymentTarget string       `json:"deployment-target" help:"MacOS version."`
	Copyright        string       `json:"copyright"         help:"Human readable copyright."`
	Role             string       `json:"role"              help:"Application role: Editor|Viewer|Shell|None."`
	Category         string       `json:"category"          help:"Applicaton category type.\nSee https://developer.apple.com/library/content/documentation/General/Reference/InfoPlistKeyReference/Articles/LaunchServicesKeys.html#//apple_ref/doc/uid/TP40009250-SW8."`
	Sandbox          bool         `json:"sandbox"           help:"Defines if the app will run in sandbox mode."`
	Capabilities     capabilities `json:"capabilities"      help:"App capabilities. Required by the Mac App Store. Requires sandbox mode."`
	Store            bool         `json:"store"             help:"Creates a .pkg ready to be uploaded with Application Loader."`
	SignID           string       `json:"sign-id"           help:"signing id. security find-identity -v -p codesigning (to see available ids)."`
	SupportedFiles   []string     `json:"supported-files"   help:"List of UTI representing the file types the app can open."`
}

type capabilities struct {
	Network    networkCap    `json:"network"     help:"Network capabilities."`
	Hardware   harwareCap    `json:"hardware"    help:"Hardware capabilities."`
	AppData    appDataCap    `json:"app-data"    help:"Application data capabilities."`
	FileAccess fileAccessCap `json:"file-access" help:"File access capabilities."`
}

type networkCap struct {
	In  bool `json:"in"  help:"Incoming connections (Server)."`
	Out bool `json:"out" help:"Outgoing connections (Client)."`
}

type harwareCap struct {
	Camera     bool `json:"camera"     help:"Use of camera."`
	Microphone bool `json:"microphone" help:"Use of microphone."`
	USB        bool `json:"usb"        help:"Use of USB."`
	Printing   bool `json:"printing"   help:"Use of printer."`
	Bluetooth  bool `json:"bluetooth"  help:"Use of bluetooth."`
}

type appDataCap struct {
	Contacts bool `json:"contacts" help:"Access contacts."`
	Location bool `json:"location" help:"Access location."`
	Calendar bool `json:"calendar" help:"Access calendar."`
}

type fileAccessCap struct {
	UserSelected fileAccess `json:"user-selected" help:"Access files from file pickers: read-only|read-write|\"\"."`
	Downloads    fileAccess `json:"downloads"     help:"Access default Downloads directory: read-only|read-write|\"\"."`
	Pictures     fileAccess `json:"pictures"      help:"Access default Pictures directory: read-only|read-write|\"\"."`
	Music        fileAccess `json:"music"         help:"Access default Music directory: read-only|read-write|\"\"."`
	Movies       fileAccess `json:"movies"        help:"Access default Movies directory: read-only|read-write|\"\"."`
}

type fileAccess string

const (
	fileNoAccess        fileAccess = ""
	fileReadAccess      fileAccess = "read-only"
	fileReadWriteAccess fileAccess = "read-write"
)

func (c Config) appName() string {
	return c.Name + ".app"
}

func commandsString() string {
	b := bytes.Buffer{}
	fmt.Fprintf(&b, "build\t Builds the .app.\n")
	fmt.Fprintf(&b, "sass\t Launches sass --watch resources/scss/:resources/css/.\n")
	return b.String()
}

func defaultConfig() Config {
	wd, err := os.Getwd()
	if err != nil {
		log.Panic(errors.New(err))
	}
	name := filepath.Base(wd)

	return Config{
		Name:             name,
		ID:               fmt.Sprintf("%v.%v", os.Getenv("USER"), name),
		Version:          "1.0",
		BuildNumber:      1,
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
	validUTI := regexp.MustCompile(`^([A-Za-z0-9]|[\-]|[\.])+$`)
	validMinVersion := regexp.MustCompile(`^[0-9]+([\.][0-9]+){1,2}$`)

	if !validName.MatchString(c.Name) {
		return fmt.Errorf("name from config must contain alphanumeric characters, '_' or '-' : %v", c.Name)
	}

	if !validUTI.MatchString(c.ID) {
		return fmt.Errorf("id from config must contain alphanumeric characters, '-' or '.' : %v", c.ID)
	}

	if !validMinVersion.MatchString(c.Version) {
		return fmt.Errorf("version from config must follow the pattern x.x.x where x is a non negative number: %v", c.Version)
	}

	if !validMinVersion.MatchString(c.DeploymentTarget) {
		return fmt.Errorf("deployment-target from config must follow the pattern x.x.x where x is a non negative number: %v", c.DeploymentTarget)
	}

	macOSDeploymentTarget := strings.Split(c.DeploymentTarget, ".")
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
