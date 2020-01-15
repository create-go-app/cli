package cgapp

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// File function will print any string of text to a file safely
// by checking for errors and syncing at the end
func File(name, data string) error {
	// Define folder
	folder := filepath.Join(appPath, name)

	// Create file
	file, err := os.Create(folder)
	ErrChecker(err)
	defer file.Close()

	// Write to created file
	_, err = io.WriteString(file, data)
	ErrChecker(err)

	// Show report for file
	fmt.Printf("%v[OK]%v File '%v' was created!\n", green, noColor, name)

	// Sync & return
	return file.Sync()
}
