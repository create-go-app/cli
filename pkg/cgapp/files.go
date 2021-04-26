// Copyright 2019-present Vic Shóstak. All rights reserved.
// Use of this source code is governed by Apache 2.0 license
// that can be found in the LICENSE file.

package cgapp

import (
	"embed"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"
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
			return fmt.Errorf(
				ShowMessage("error", "Can't copy files from embedded path `"+efs.RootFolder+"`!", true, true),
			)
		}

		// Checking, if embedded file is a folder.
		if entry.IsDir() && !efs.SkipDir {
			// Create folders structure from embedded.
			if err := MakeFolder(path); err != nil {
				return fmt.Errorf(
					ShowMessage("error", err.Error(), true, true),
				)
			}
		}

		// Checking, if embedded file is not a folder.
		if !entry.IsDir() {
			// Set file data.
			fileData, errReadFile := fs.ReadFile(efs.Name, path)
			if errReadFile != nil {
				return fmt.Errorf(
					ShowMessage("error", "File `"+path+"/"+entry.Name()+"` was broken!", true, true),
				)
			}

			// Path to file, if skipped folders.
			if efs.SkipDir {
				path = entry.Name()
			}

			// Create file from embedded.
			if errMakeFile := MakeFile(path, fileData); errMakeFile != nil {
				return fmt.Errorf(
					ShowMessage("error", errMakeFile.Error(), true, true),
				)
			}
		}

		return nil
	}); err != nil {
		return fmt.Errorf(
			ShowMessage("error", err.Error(), true, true),
		)
	}

	return nil
}

// GenerateFileFromTemplate ...
func GenerateFileFromTemplate(fileName string, variables map[string]interface{}) error {
	//
	tmpl, errParseFiles := template.ParseFiles(fileName)
	if errParseFiles != nil {
		return fmt.Errorf(
			ShowMessage("error", errParseFiles.Error(), true, true),
		)
	}

	//
	file, errCreate := os.Create(fileName)
	if errCreate != nil {
		return fmt.Errorf(
			ShowMessage("error", errCreate.Error(), true, true),
		)
	}

	//
	if errExecute := tmpl.Execute(file, variables); errExecute != nil {
		return fmt.Errorf(
			ShowMessage("error", errExecute.Error(), true, true),
		)
	}
	_ = file.Close()

	resultFile, errOpen := os.Open(filepath.Clean(fileName))
	if errOpen != nil {
		return fmt.Errorf(
			ShowMessage("error", errOpen.Error(), true, true),
		)
	}

	if _, errCopy := io.Copy(os.Stdout, resultFile); errCopy != nil {
		return fmt.Errorf(
			ShowMessage("error", errCopy.Error(), true, true),
		)
	}
	_ = resultFile.Close()

	//
	newFileName := strings.Replace(fileName, ".tmpl", "", -1)
	if errRename := os.Rename(fileName, newFileName); errRename != nil {
		return fmt.Errorf(
			ShowMessage("error", errRename.Error(), true, true),
		)
	}

	return nil
}

// MakeFile function for single file create.
func MakeFile(fileName string, fileData []byte) error {
	// Write to created file.
	if err := ioutil.WriteFile(fileName, fileData, 0600); err != nil {
		return fmt.Errorf(
			ShowMessage("error", "File `"+fileName+"` was not created!", true, true),
		)
	}

	// Show report for file.
	_ = ShowMessage("success", "File `"+fileName+"` was created!", false, true)

	return nil
}

// MakeFolder function for create folder.
func MakeFolder(folderName string) error {
	// Check if folder exists, fail if it does.
	if err := os.Mkdir(folderName, 0750); err != nil {
		return fmt.Errorf(
			ShowMessage("error", "Folder `"+folderName+"` is exists!", true, true),
		)
	}

	// Show report for folder.
	_ = ShowMessage("success", "Folder `"+folderName+"` was created!", false, true)

	return nil
}

// RemoveFolders function for massively remove folders.
func RemoveFolders(rootFolder string, foldersToRemove []string) {
	for _, folder := range foldersToRemove {
		_ = os.RemoveAll(filepath.Join(rootFolder, folder))
	}
}
