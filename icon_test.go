package main

import "testing"

func TestGenerateIcon(t *testing.T) {
	cfg.Icon = "logo.png"

	createPackage()
	defer removePackage()

	syncResources()

	if err := generateIcon(); err != nil {
		t.Error(err)
	}
}
