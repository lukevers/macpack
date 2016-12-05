package main

import "testing"

func TestExecCmd(t *testing.T) {
	if err := execCmd("ls", "-la"); err != nil {
		t.Fatal(err)
	}
}
