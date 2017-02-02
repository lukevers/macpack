package main

import "testing"

func TestBuild(t *testing.T) {
	cfg = defaultConfig()
	if err := build(); err != nil {
		removePackage()
		t.Error(err)
	}
}
