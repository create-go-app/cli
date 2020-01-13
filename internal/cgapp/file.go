package cgapp

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// File will print any string of text to a file safely by
// checking for errors and syncing at the end.
func File(name, data string) error {
	//
	folder := filepath.Join(appPath, name)

	//
	file, err := os.Create(folder)
	ErrChecker(err)
	defer file.Close()

	//
	_, err = io.WriteString(file, data)
	ErrChecker(err)

	// Show report for each file
	fmt.Printf("%v[OK]%v File '%v' was created!\n", green, noColor, name)

	//
	return file.Sync()
}
