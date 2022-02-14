// Copyright 2022 Vic Shóstak and Create Go App Contributors. All rights reserved.
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
func ExecCommand(command string, options []string, silentMode bool) error {
	// Checking for nil.
	if command == "" || options == nil {
		return fmt.Errorf("no command to execute")
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
		return ShowError(errStdoutPipe.Error())
	}

	// Start executing command.
	if errStart := cmd.Start(); errStart != nil {
		return ShowError(errStart.Error())
	}

	// Create a new scanner and run goroutine func with output, if not in silent mode.
	if !silentMode {
		scanner := bufio.NewScanner(cmdReader)
		go func() {
			for scanner.Scan() {
				ShowMessage("", scanner.Text(), false, false)
			}
		}()
	}

	// Wait for executing command.
	if errWait := cmd.Wait(); errWait != nil {
		return ShowError(errWait.Error())
	}

	return nil
}
