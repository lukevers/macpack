package main

import "testing"
import "os"

func TestDefaultConfig(t *testing.T) {
	t.Logf("%+v", defaultConfig())
}

func TestSaveConfig(t *testing.T) {
	// Create.
	if err := saveConfig(defaultConfig(), confName); err != nil {
		t.Fatal(err)
	}

	defer os.Remove(confName)

	// Override
	if err := saveConfig(defaultConfig(), confName); err != nil {
		t.Fatal(err)
	}
}

func TestReadConfig(t *testing.T) {
	// Existing conf.
	saveConfig(defaultConfig(), confName)
	defer os.Remove(confName)

	conf, err := readConfig(confName)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", conf)

	// Nonexistent conf.
	if _, err = readConfig("foo.json"); err == nil {
		t.Error("readConfig should return an error")
	}
}

func TestCheckConfigName(t *testing.T) {
	conf := defaultConfig()

	conf.Name = "foo-bar_boo"
	if err := checkConfig(conf); err != nil {
		t.Error(err)
	}

	conf.Name = ""
	if err := checkConfig(conf); err == nil {
		t.Error("checkConfig should return an error")
	}

	conf.Name = "foo.bar"
	if err := checkConfig(conf); err == nil {
		t.Error("checkConfig should return an error")
	}
}

func TestCheckConfigVersion(t *testing.T) {
	conf := defaultConfig()
	conf.Version = "1.0.0.42"

	if err := checkConfig(conf); err != nil {
		t.Error(err)
	}

	conf.Version = ""
	if err := checkConfig(conf); err == nil {
		t.Error("checkConfig should return an error")
	}

	conf.Version = "1.42"
	if err := checkConfig(conf); err == nil {
		t.Error("checkConfig should return an error")
	}

	conf.Version = "1.42.f.1"
	if err := checkConfig(conf); err == nil {
		t.Error("checkConfig should return an error")
	}
}

func TestCheckConfigID(t *testing.T) {
	conf := defaultConfig()

	conf.ID = "com.murlok.Test-ID"
	if err := checkConfig(conf); err != nil {
		t.Error(err)
	}

	conf.ID = ""
	if err := checkConfig(conf); err == nil {
		t.Error("checkConfig should return an error")
	}

	conf.ID = "foo_bar"
	if err := checkConfig(conf); err == nil {
		t.Error("checkConfig should return an error")
	}
}

func TestCheckConfigRole(t *testing.T) {
	conf := defaultConfig()

	conf.Role = "Editor"
	if err := checkConfig(conf); err != nil {
		t.Error(err)
	}

	conf.Role = "Viewer"
	if err := checkConfig(conf); err != nil {
		t.Error(err)
	}

	conf.Role = "Shell"
	if err := checkConfig(conf); err != nil {
		t.Error(err)
	}

	conf.Role = "None"
	if err := checkConfig(conf); err != nil {
		t.Error(err)
	}

	conf.Role = "Logger"
	if err := checkConfig(conf); err == nil {
		t.Error("checkConfig should return an error")
	}
}

func TestCheckConfigOSMinVersion(t *testing.T) {
	conf := defaultConfig()
	conf.OSMinVersion = "10.15"

	if err := checkConfig(conf); err != nil {
		t.Error(err)
	}

	conf.OSMinVersion = ""
	if err := checkConfig(conf); err == nil {
		t.Error("checkConfig should return an error")
	}

	conf.OSMinVersion = "1.a"
	if err := checkConfig(conf); err == nil {
		t.Error("checkConfig should return an error")
	}

	conf.OSMinVersion = "9.12"
	if err := checkConfig(conf); err == nil {
		t.Error("checkConfig should return an error")
	}

	conf.OSMinVersion = "10.11"
	if err := checkConfig(conf); err == nil {
		t.Error("checkConfig should return an error")
	}
}

func TestCheckConfigSupportedFiles(t *testing.T) {
	conf := defaultConfig()
	conf.SupportedFiles = []string{
		"public.jpeg",
	}

	if err := checkConfig(conf); err != nil {
		t.Error(err)
	}

	conf.SupportedFiles = append(conf.SupportedFiles, "public_png")
	if err := checkConfig(conf); err == nil {
		t.Error("checkConfig should return an error")
	}
}
