package main

import "testing"
import "os"

func TestGobuild(t *testing.T) {
	defer os.Remove(goExecName())

	if err := gobuild(); err != nil {
		t.Fatal(err)
	}
}

func TestCreatePackage(t *testing.T) {
	defer removePackage()

	if err := createPackage(); err != nil {
		t.Fatal(err)
	}
}

func TestCreateExec(t *testing.T) {
	gobuild()
	defer os.Remove(goExecName())

	createPackage()
	defer removePackage()

	if err := createExec(); err != nil {
		t.Error(err)
	}
}

func TestCreateResources(t *testing.T) {
	if err := createResources(); err != nil {
		t.Fatal(err)
	}
	removePackage()
}

func TestSyncResources(t *testing.T) {
	createPackage()
	defer removePackage()

	if err := syncResources(); err != nil {
		t.Error(err)
	}
}

func TestSyncResourcesNoResources(t *testing.T) {
	if err := syncResources(); err == nil {
		t.Error(err)
	}
}
