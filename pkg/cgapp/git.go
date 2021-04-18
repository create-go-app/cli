// Copyright 2019-present Vic Shóstak. All rights reserved.
// Use of this source code is governed by Apache 2.0 license
// that can be found in the LICENSE file.

package cgapp

import "github.com/go-git/go-git/v5"

// GitClone function for `git clone` defined project template.
func GitClone(rootFolder, templateName string) error {
	// Clone project template.
	_, err := git.PlainClone(rootFolder, false, &git.CloneOptions{
		URL: "https://" + templateName,
	})
	if err != nil {
		return throwError("Repository was not cloned!")
	}

	return nil
}
