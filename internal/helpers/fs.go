package helpers

import (
	"os"
)

// MakeFile function for single file create.
func MakeFile(name string, data []byte) error {
	// Check, if file is existing.
	_, err := os.Stat(name)
	if err != nil {
		// Write to created file.
		if err = os.WriteFile(name, data, 0o644); err != nil {
			return err
		}
	}

	return nil
}

// MakeFolder function for create folder.
func MakeFolder(name string) error {
	// Check, if folder is existing.
	folderInfo, err := os.Stat(name)
	if err != nil && folderInfo.IsDir() {
		return err
	}

	// Check if folder exists, fail if it does.
	if err = os.Mkdir(name, 0o644); err != nil {
		return err
	}

	return nil
}

// RemoveFolders function for massively remove folders.
func RemoveFolders(names []string) {
	for _, name := range names {
		_ = os.RemoveAll(name)
	}
}
