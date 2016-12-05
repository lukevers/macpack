package main

import (
	"os"
	"testing"
)

func TestCreatePlist(t *testing.T) {
	conf := defaultConfig()
	conf.Icon = "murlok.png"
	conf.SupportedFiles = []string{
		"public.gif",
	}

	createPackage(conf)
	defer os.RemoveAll(conf.Name + ".app")

	if err := createPlist(conf); err != nil {
		t.Error(err)
	}
}

func TestCreatePlistError(t *testing.T) {
	conf := defaultConfig()
	if err := createPlist(conf); err == nil {
		t.Error("createPlist should return an error")
	}
}
