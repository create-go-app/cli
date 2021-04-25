// Copyright 2019-present Vic Shóstak. All rights reserved.
// Use of this source code is governed by Apache 2.0 license
// that can be found in the LICENSE file.

package cgapp

import (
	"bufio"
	"bytes"
	"fmt"
	"os/exec"
)

// ExecCommand function to execute a given command.
func ExecCommand(command string, options []string) error {
	//
	if command == "" {
		return fmt.Errorf("No command to execute!")
	}

	// Create buffer for stderr.
	stderr := &bytes.Buffer{}

	// Collect command line
	cmd := exec.Command(command, options...) // #nosec G204

	// Set buffer for stderr from cmd
	cmd.Stderr = stderr

	// Create a new reader
	cmdReader, err := cmd.StdoutPipe()
	catchError("", err)

	// Start executing command.
	errStart := cmd.Start()
	catchError(stderr.String(), errStart)

	// Create a new scanner and run goroutine func with output.
	scanner := bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan() {
			SendMsg(false, "*", scanner.Text(), "cyan", false)
		}
	}()

	// Wait for executing command.
	errWait := cmd.Wait()
	catchError(stderr.String(), errWait)

	return nil
}
