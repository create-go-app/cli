// Copyright 2023 Vic Shóstak and Create Go App Contributors. All rights reserved.
// Use of this source code is governed by Apache 2.0 license
// that can be found in the LICENSE file.

package helpers

import (
	"embed"
	"errors"
	"html/template"
	"io/fs"
	"os"

	"github.com/create-go-app/cli/v5/internal/config"
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
			return errors.New("can't copy files from embedded path")
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
func GenerateFileFromTemplate(fs embed.FS, templateName, outputName string, vars *config.Config) error {
	// Checking template and output file names.
	if templateName == "" || outputName == "" {
		return errors.New("empty template or output file name")
	}

	// Parse template from embed file system.
	tmpl, err := template.ParseFS(fs, templateName)
	if err != nil {
		return err
	}

	// Create a new temp file with the given data.
	file, err := os.CreateTemp("", "*")
	if err != nil {
		return err
	}

	// Rename temp file.
	if err = os.Rename(file.Name(), outputName); err != nil {
		return err
	}

	// Set variables to the given.
	if err = tmpl.Execute(file, vars); err != nil {
		return err
	}

	return file.Close()
}
