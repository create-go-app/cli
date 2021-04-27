// Copyright 2019-present Vic Shóstak. All rights reserved.
// Use of this source code is governed by Apache 2.0 license
// that can be found in the LICENSE file.

package cgapp

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/go-git/go-git/v5"
)

// GitClone function for `git clone` defined project template.
func GitClone(templateType, templateURL string) error {
	// Get current directory.
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(ShowMessage("error", err.Error(), true, true))
	}

	//
	folder := filepath.Join(currentDir, templateType)

	// Clone project template.
	_, errPlainClone := git.PlainClone(
		folder,
		false,
		&git.CloneOptions{
			URL: "https://" + templateURL,
		},
	)
	if errPlainClone != nil {
		return fmt.Errorf(
			ShowMessage("error", "Repository `"+templateURL+"` was not cloned!", true, true),
		)
	}

	// Cleanup project.
	RemoveFolders(folder, []string{".git", ".github"})

	return nil
}
