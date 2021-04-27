// Copyright 2019-present Vic Shóstak. All rights reserved.
// Use of this source code is governed by Apache 2.0 license
// that can be found in the LICENSE file.

package cgapp

import "fmt"

// CreateProjectFromCmd function for create a new frontend project from cmd.
func CreateProjectFromCmd(options []string) error {
	// Checking for nil.
	if options == nil {
		return fmt.Errorf("Options not found!")
	}

	// Run execution command.
	if errExecCommand := ExecCommand("npm", options); errExecCommand != nil {
		return fmt.Errorf(
			ShowMessage("error", errExecCommand.Error(), true, true),
		)
	}

	return nil
}

// CreateProjectFromGit function for create a new project from a comand line.
func CreateProjectFromGit(projectType, projectRepository string) error {
	// Checking for nil.
	if projectRepository == "" {
		return fmt.Errorf("Project template not found!")
	}

	// Create frontend from given repository (GitHub, etc).
	if err := GitClone(projectType, projectRepository); err != nil {
		return fmt.Errorf(
			ShowMessage("error", err.Error(), true, true),
		)
	}

	return nil
}
