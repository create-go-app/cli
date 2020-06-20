package cgapp

import (
	"io/ioutil"
	"path/filepath"
)

// File function will print any string of text to a file safely
// by checking for errors and syncing at the end
func File(name string, data []byte) error {
	// Define folder
	folder := filepath.Join(appPath, name)

	// Write to created file
	err := ioutil.WriteFile(folder, data, 0755)
	ErrChecker(err)

	// Show report for file
	SendMessage("[OK] File `"+name+"` was created!", "")

	return nil
}
