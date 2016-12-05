package main

import "testing"
import "os"

func TestCreatePackage(t *testing.T) {
	conf := defaultConfig()
	defer os.RemoveAll(conf.Name + ".app")

	if err := createPackage(conf); err != nil {
		t.Fatal(err)
	}
}
