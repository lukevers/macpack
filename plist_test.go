package main

import "testing"

func TestCreatePlist(t *testing.T) {
	cfg.Icon = "murlok.png"
	cfg.SupportedFiles = []string{
		"public.gif",
	}

	createPackage()
	defer removePackage()

	if err := createPlist(defaultConfig()); err != nil {
		t.Error(err)
	}
}

func TestCreatePlistError(t *testing.T) {
	if err := createPlist(defaultConfig()); err == nil {
		t.Error("createPlist should return an error")
	}
}
