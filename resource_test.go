package main

import (
	"os"
	"testing"
)

func TestSyncResources(t *testing.T) {
	conf := defaultConfig()

	createPackage(conf)
	defer os.RemoveAll(conf.Name + ".app")

	if err := syncResources(conf, true, true); err != nil {
		t.Error(err)
	}
}

func TestSyncResourcesNoResources(t *testing.T) {
	conf := defaultConfig()
	if err := syncResources(conf, false, false); err == nil {
		t.Error(err)
	}
}
