package main

import (
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	conf := defaultConfig()
	defer os.RemoveAll(conf.Name + ".app")
	main()
}
