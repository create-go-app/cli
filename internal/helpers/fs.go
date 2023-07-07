package helpers

import (
	"os"
)

// MakeFile function for single file create.
func MakeFile(fileName string, fileData []byte) error {
	// Write to created file.
	if err := os.WriteFile(fileName, fileData, 0o644); err != nil {
		return err
	}

	return nil
}
