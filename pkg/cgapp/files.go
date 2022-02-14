// Copyright 2022 Vic Shóstak and Create Go App Contributors. All rights reserved.
// Use of this source code is governed by Apache 2.0 license
// that can be found in the LICENSE file.

package cgapp

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

// EmbeddedFileSystem struct contains embedded file system fields.
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
			return ShowError(
				fmt.Sprintf("Can't copy files from embedded path `%v`!", efs.RootFolder),
			)
		}

		// Checking, if embedded file is a folder.
		if entry.IsDir() && !efs.SkipDir {
			// Create folders structure from embedded.
			if err := MakeFolder(path); err != nil {
				return err
			}
		}

		// Checking, if embedded file is not a folder.
		if !entry.IsDir() {
			// Set file data.
			fileData, errReadFile := fs.ReadFile(efs.Name, path)
			if errReadFile != nil {
				return errReadFile
			}

			// Path to file, if skipped folders.
			if efs.SkipDir {
				path = entry.Name()
			}

			// Create file from embedded.
			if errMakeFile := MakeFile(path, fileData); errMakeFile != nil {
				return errMakeFile
			}
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

// GenerateFileFromTemplate func to generate files from templates.
func GenerateFileFromTemplate(fileName string, variables map[string]interface{}) error {
	// Checking file name.
	if fileName == "" {
		return ShowError(
			fmt.Sprintf("Not correct or empty file name (given: `%s`)!", fileName),
		)
	}

	// Clean file name.
	cleanFileName := filepath.Clean(fileName)

	// Parse template.
	tmpl, errParseFiles := template.ParseFiles(cleanFileName)
	if errParseFiles != nil {
		return ShowError(errParseFiles.Error())
	}

	// Create a new file with template data.
	file, errCreate := os.Create(cleanFileName)
	if errCreate != nil {
		return ShowError(errCreate.Error())
	}

	// Execute template with variables.
	if errExecute := tmpl.Execute(file, variables); errExecute != nil {
		return ShowError(errExecute.Error())
	}
	_ = file.Close()

	// Rename output file.
	newFileName := strings.ReplaceAll(cleanFileName, ".tmpl", "")
	if errRename := os.Rename(cleanFileName, newFileName); errRename != nil {
		return ShowError(errRename.Error())
	}

	return nil
}

// MakeFile function for single file create.
func MakeFile(fileName string, fileData []byte) error {
	// Write to created file.
	if err := os.WriteFile(fileName, fileData, 0o600); err != nil {
		return err
	}

	return nil
}

// MakeFolder function for create folder.
func MakeFolder(folderName string) error {
	// Check if folder exists, fail if it does.
	if err := os.Mkdir(folderName, 0o750); err != nil {
		return err
	}

	return nil
}

// RemoveFolders function for massively remove folders.
func RemoveFolders(rootFolder string, foldersToRemove []string) {
	for _, folder := range foldersToRemove {
		_ = os.RemoveAll(filepath.Join(rootFolder, folder))
	}
}
