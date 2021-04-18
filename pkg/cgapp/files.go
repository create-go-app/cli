// Copyright 2019-present Vic Shóstak. All rights reserved.
// Use of this source code is governed by Apache 2.0 license
// that can be found in the LICENSE file.

package cgapp

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// MakeFiles function for massively create folders.
func MakeFiles(rootFolder string, filesToMake map[string][]byte) error {
	for file, data := range filesToMake {
		folder := filepath.Join(rootFolder, file)

		// Write to created file.
		if err := ioutil.WriteFile(folder, data, 0600); err != nil {
			return throwError("File `" + file + "` was not created!")
		}

		// Show report for file.
		SendMsg(false, "[OK]", "File `"+file+"` was created!", "cyan", false)
	}

	return nil
}

// MakeFolder function for create folder.
func MakeFolder(folderName string, chmod os.FileMode) error {
	// Check if folder exists, fail if it does.
	if _, err := os.Stat(folderName); !os.IsNotExist(err) {
		return throwError("Folder `" + folderName + "` exists!")
	}

	// Create folder.
	if err := os.Mkdir(folderName, chmod); err != nil {
		return throwError("Folder `" + folderName + "` was not created!")
	}

	// Show report for folder.
	SendMsg(false, "OK", "Folder `"+folderName+"` was created!", "", false)

	return nil
}

// RemoveFolders function for massively remove folders.
func RemoveFolders(rootFolder string, foldersToRemove []string) {
	for _, folder := range foldersToRemove {
		_ = os.RemoveAll(filepath.Join(rootFolder, folder))
	}
}
