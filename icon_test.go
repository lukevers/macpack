package main

import "testing"

func TestGenerateIcon(t *testing.T) {
	config.Icon = "logo.png"

	createPackage()
	defer removePackage()

	syncResources()

	if err := generateIcon(); err != nil {
		t.Error(err)
	}
}
