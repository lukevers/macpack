package main

import (
	"bufio"
	"io"
	"os"
	"os/exec"
)

func execCmd(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)

	cmdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	cmderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	go printOutput(cmdout, os.Stdout)
	go printOutput(cmderr, os.Stderr)

	if err = cmd.Start(); err != nil {
		return err
	}
	return cmd.Wait()
}

func printOutput(r io.Reader, output io.Writer) {
	reader := bufio.NewReader(r)

	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				return
			}
		}
		output.Write(line)
	}
}
