// Copyright 2022 Vic Shóstak and Create Go App Contributors. All rights reserved.
// Use of this source code is governed by Apache 2.0 license
// that can be found in the LICENSE file.

package cgapp

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-git/go-git/v5"
)

// GitClone function for `git clone` defined project template.
func GitClone(templateType, templateURL string) error {
	// Checking for nil.
	if templateType == "" || templateURL == "" {
		return fmt.Errorf("project template not found")
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
			URL: getAbsoluteURL(templateURL),
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

// getAbsolutURL func for help define correct HTTP protocol.
func getAbsoluteURL(templateURL string) string {
	templateURL = strings.TrimSpace(templateURL)
	u, _ := url.Parse(templateURL)

	if u.Scheme == "" {
		u.Scheme = "https"
	}

	return u.String()
}
