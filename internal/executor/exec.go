// Copyright 2023 Vic Shóstak and Create Go App Contributors. All rights reserved.
// Use of this source code is governed by Apache 2.0 license
// that can be found in the LICENSE file.

package executor

import (
	"errors"
	"fmt"
	"os/exec"
)

// Execute function to execute a given command with options.
func Execute(command string, options ...string) error {
	// Checking for nil.
	if command == "" || options == nil {
		return errors.New("no command or options to execute")
	}

	// Create a process for execute the current command.
	cmd := exec.Command(command, options...)

	// Run execution of the current command.
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("command '%s' ended with an error: %s", command, err.Error())
	}

	return nil
}
