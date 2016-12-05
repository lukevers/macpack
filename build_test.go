package main

import "testing"
import "os"

func TestBuild(t *testing.T) {
	defer os.Remove("macpack")

	if err := build(true); err != nil {
		t.Fatal(err)
	}
}

func TestCreateExec(t *testing.T) {
	conf := defaultConfig()
	build(false)

	createPackage(conf)
	defer os.RemoveAll(conf.Name + ".app")

	if err := createExec(conf); err != nil {
		t.Error(err)
	}
}
