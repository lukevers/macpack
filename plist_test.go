package main

import "testing"

func TestCreatePlist(t *testing.T) {
	config.Icon = "murlok.png"
	config.SupportedFiles = []string{
		"public.gif",
	}

	createPackage()
	defer removePackage()

	if err := createPlist(); err != nil {
		t.Error(err)
	}
}

func TestCreatePlistError(t *testing.T) {
	if err := createPlist(); err == nil {
		t.Error("createPlist should return an error")
	}
}
