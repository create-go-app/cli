// Copyright 2019-present Vic Shóstak. All rights reserved.
// Use of this source code is governed by Apache 2.0 license
// that can be found in the LICENSE file.

package cgapp

import (
	"embed"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// EmbeddedFileSystems struct contains embedded file system fields.
type EmbeddedFileSystem struct {
	Name       embed.FS
	RootFolder string
	SkipDir    bool
}

// CopyFromEmbeddedFS function for copy files from embedded file system.
func CopyFromEmbeddedFS(efs *EmbeddedFileSystem) error {
	// Return copied folders and files.
	if err := fs.WalkDir(efs.Name, efs.RootFolder, func(path string, entry fs.DirEntry, err error) error {
		// Checking embed path.
		if err != nil {
			log.Fatal(BeautifyText("Can't make structure from embedded path `"+efs.RootFolder+"`!", "red"))
		}

		// Checking, if embedded file is a folder.
		if entry.IsDir() && !efs.SkipDir {
			// Create folders structure from embedded.
			if err := MakeFolder(path); err != nil {
				log.Fatal(BeautifyText(err.Error(), "red"))
			}
		}

		// Checking, if embedded file is not a folder.
		if !entry.IsDir() {
			// Set file data.
			fileData, errReadFile := fs.ReadFile(efs.Name, path)
			if errReadFile != nil {
				log.Fatal(BeautifyText("File `"+path+"/"+entry.Name()+"` was broken!", "red"))
			}

			// Path to file, if skipped folders.
			if efs.SkipDir {
				path = entry.Name()
			}

			// Create file from embedded.
			if errMakeFile := MakeFile(path, fileData); errMakeFile != nil {
				log.Fatal(BeautifyText(err.Error(), "red"))
			}
		}

		return nil
	}); err != nil {
		log.Fatal(BeautifyText(err.Error(), "red"))
	}

	return nil
}

// MakeFile function for single file create.
func MakeFile(fileName string, fileData []byte) error {
	// Write to created file.
	if err := ioutil.WriteFile(fileName, fileData, 0600); err != nil {
		log.Fatal(BeautifyText("File `"+fileName+"` was not created!", "red"))
	}

	// Show report for file.
	SendMsg(false, "[OK]", "File `"+fileName+"` was created!", "cyan", false)

	return nil
}

// MakeFolder function for create folder.
func MakeFolder(folderName string) error {
	// Check if folder exists, fail if it does.
	if err := os.Mkdir(folderName, 0750); err != nil {
		log.Fatal(BeautifyText("Folder `"+folderName+"` is exists!", "red"))
	}

	// Show report for folder.
	SendMsg(false, "[OK]", "Folder `"+folderName+"` was created!", "cyan", false)

	return nil
}

// RemoveFolders function for massively remove folders.
func RemoveFolders(rootFolder string, foldersToRemove []string) {
	for _, folder := range foldersToRemove {
		_ = os.RemoveAll(filepath.Join(rootFolder, folder))
	}
}
