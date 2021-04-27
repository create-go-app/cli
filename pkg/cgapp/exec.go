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
	// Checking for nil.
	if command == "" || options == nil {
		return fmt.Errorf("No command to execute!")
	}

	// Create buffer for stderr.
	stderr := &bytes.Buffer{}

	// Collect command line.
	cmd := exec.Command(command, options...) // #nosec G204

	// Set buffer for stderr from cmd.
	cmd.Stderr = stderr

	// Create a new reader.
	cmdReader, errStdoutPipe := cmd.StdoutPipe()
	if errStdoutPipe != nil {
		return fmt.Errorf(
			ShowMessage("error", errStdoutPipe.Error(), true, true),
		)
	}

	// Start executing command.
	if errStart := cmd.Start(); errStart != nil {
		return fmt.Errorf(
			ShowMessage("error", stderr.String(), true, true),
		)
	}

	// Create a new scanner and run goroutine func with output.
	scanner := bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan() {
			_ = ShowMessage("info", scanner.Text(), false, false)
		}
	}()

	// Wait for executing command.
	if errWait := cmd.Wait(); errWait != nil {
		return fmt.Errorf(
			ShowMessage("error", stderr.String(), true, true),
		)
	}

	return nil
}
