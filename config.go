package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

const (
	confName = "mac.json"
)

// Config represents the configuration of the app to be packaged.
type Config struct {
	Name           string   `json:"name"`            // Name displayed in menu and dock
	Version        string   `json:"version"`         // 0.0.0.0
	Icon           string   `json:"icon"`            // .png
	ID             string   `json:"id"`              // UTI in reverse-DNS format with (A-Za-z0-9), (-) and (.) eg com.murlok.Hello-World
	OSMinVersion   string   `json:"os-min-version"`  // >= 10.12
	Role           string   `json:"role"`            // Editor | Viewer | Shell | None
	Sandbox        bool     `json:"sandbox"`         // Sandbox mode
	SupportedFiles []string `json:"supported-files"` // Slice of UTI representing types of the supported files
}

func defaultConfig() Config {
	wd, err := os.Getwd()
	if err != nil {
		log.Panic(err)
	}

	name := filepath.Base(wd)

	return Config{
		Name:           name,
		Version:        "1.0.0.0",
		ID:             fmt.Sprintf("%v.%v", os.Getenv("USER"), name),
		OSMinVersion:   "10.12",
		Role:           "None",
		Sandbox:        true,
		SupportedFiles: []string{},
	}
}

func saveConfig(conf Config, name string) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}

	defer f.Close()

	enc := json.NewEncoder(f)
	enc.SetIndent("", "    ")
	return enc.Encode(conf)
}

func readConfig(name string) (conf Config, err error) {
	f, err := os.Open(name)
	if err != nil {
		return
	}

	defer f.Close()

	dec := json.NewDecoder(f)
	err = dec.Decode(&conf)
	return
}

func checkConfig(conf Config) error {
	validName := regexp.MustCompile(`^([A-Za-z0-9_]|[\-])+$`)
	validVersion := regexp.MustCompile(`^[0-9]+([\.][0-9]+){3}$`)
	validUTI := regexp.MustCompile(`^([A-Za-z0-9]|[\-]|[\.])+$`)
	validOSVersion := regexp.MustCompile(`^[0-9]+[\.][0-9]+$`)

	if !validName.MatchString(conf.Name) {
		return fmt.Errorf("name from config must contain alphanumeric characters, '_' or '-' : %v", conf.Name)
	}

	if !validVersion.MatchString(conf.Version) {
		return fmt.Errorf("version from config must follow the pattern x.x.x.x where x is a non negative number: %v", conf.Version)
	}

	if !validUTI.MatchString(conf.ID) {
		return fmt.Errorf("id from config must contain alphanumeric characters, '-' or '.' : %v", conf.ID)
	}

	if role := conf.Role; role != "Editor" && role != "Viewer" && role != "Shell" && role != "None" {
		return fmt.Errorf("role from config should be Editor | Viewer | Shell | None:%v", conf.Role)
	}

	if !validOSVersion.MatchString(conf.OSMinVersion) {
		return fmt.Errorf("os-min-version from config must follow the pattern x.x where x is a non negative number: %v", conf.OSMinVersion)
	}

	osMinVersion := strings.Split(conf.OSMinVersion, ".")

	if maj, _ := strconv.Atoi(osMinVersion[0]); maj < 10 {
		return fmt.Errorf("major revision in os-min-version from config cannot be under 10:%v", maj)
	}

	if min, _ := strconv.Atoi(osMinVersion[1]); min < 12 {
		return fmt.Errorf("minor revision in os-min-version from config cannot be under 12:%v", min)
	}

	for _, uti := range conf.SupportedFiles {
		if !validUTI.MatchString(uti) {
			return fmt.Errorf("supported-files from config contains a non valid UTI: %v", uti)
		}
	}
	return nil
}
