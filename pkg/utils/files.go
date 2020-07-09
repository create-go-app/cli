/*
Package utils includes helpful utilities for the Create Go App CLI.

Copyright © 2020 Vic Shóstak <truewebartisans@gmail.com> (https://1wa.co)

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package utils

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// MakeFiles function for massively create folders
func MakeFiles(rootFolder string, filesToMake map[string][]byte) error {
	for file, data := range filesToMake {
		folder := filepath.Join(rootFolder, file)

		// Write to created file
		if err := ioutil.WriteFile(folder, data, 0755); err != nil {
			return throwError("File `" + file + "` was not created!")
		}

		// Show report for file
		SendMsg(false, "[OK]", "File `"+file+"` was created!", "cyan", false)
	}

	return nil
}

// MakeFolder function for create folder
func MakeFolder(folderName string, chmod os.FileMode) error {
	// Check if folder exists, fail if it does
	if _, err := os.Stat(folderName); !os.IsNotExist(err) {
		return throwError("Folder `" + folderName + "` exists!")
	}

	// Create folder
	if err := os.Mkdir(folderName, chmod); err != nil {
		return throwError("Folder `" + folderName + "` was not created!")
	}

	// Show report for folder
	SendMsg(false, "OK", "Folder `"+folderName+"` was created!", "", false)

	return nil
}

// RemoveFolders ...
func RemoveFolders(rootFolder string, foldersToRemove []string) {
	for _, folder := range foldersToRemove {
		_ = os.RemoveAll(filepath.Join(rootFolder, folder))
	}
}
