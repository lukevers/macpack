package main

import (
	"os"
	"testing"
)

func TestGenerateIcon(t *testing.T) {
	conf := defaultConfig()
	conf.Icon = "logo.png"

	createPackage(conf)
	defer os.RemoveAll(conf.Name + ".app")

	syncResources(conf, true, false)

	if err := generateIcon(conf); err != nil {
		t.Error(err)
	}
}
