// Copyright 2019-present Vic Shóstak. All rights reserved.
// Use of this source code is governed by Apache 2.0 license
// that can be found in the LICENSE file.

package cgapp

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-git/go-git/v5"
)

// GitClone function for `git clone` defined project template.
func GitClone(templateType, templateURL string) error {
	// Checking for nil.
	if templateType == "" || templateURL == "" {
		return fmt.Errorf("Project template not found!")
	}

	// Get current directory.
	currentDir, _ := os.Getwd()

	// Set project folder.
	folder := filepath.Join(currentDir, templateType)

	// Clone project template.
	_, errPlainClone := git.PlainClone(
		folder,
		false,
		&git.CloneOptions{
			URL: fmt.Sprintf("https://%s", templateURL),
		},
	)
	if errPlainClone != nil {
		return ShowError(
			fmt.Sprintf("Repository `%v` was not cloned!", templateURL),
		)
	}

	// Cleanup project.
	RemoveFolders(folder, []string{".git", ".github"})

	return nil
}
