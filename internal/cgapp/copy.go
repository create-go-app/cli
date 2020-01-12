package cgapp

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// CopyFolder function for copy all files from app config
func CopyFolder(folderName string) error {
	return filepath.Walk(folderName, func(path string, file os.FileInfo, err error) error {
		// Define files paths
		folder := appPath + string(os.PathSeparator) + file.Name()

		// Create files (skip directories)
		if !file.IsDir() {
			// Open file
			from, err := os.Open(path)
			ErrChecker(err)
			defer from.Close()

			// Create file
			to, err := os.Create(folder)
			ErrChecker(err)
			defer to.Close()

			// Copy file
			_, err = io.Copy(to, from)
			ErrChecker(err)

			// Show report for each file
			fmt.Printf("— File '%v' was copied!\n", folder)
		}

		// Default return
		return nil
	})
}

// CopyFile function for copy only single file from app config
func CopyFile(folderName, fileName string) error {
	return filepath.Walk(folderName, func(path string, file os.FileInfo, err error) error {
		// Define files paths
		folder := appPath + string(os.PathSeparator) + file.Name()

		// Create files (skip directories)
		if !file.IsDir() && fileName == file.Name() {
			// Open file
			from, err := os.Open(path)
			ErrChecker(err)
			defer from.Close()

			// Create file
			to, err := os.Create(folder)
			ErrChecker(err)
			defer to.Close()

			// Copy file
			_, err = io.Copy(to, from)
			ErrChecker(err)

			// Show report for each file
			fmt.Printf("\n— File '%v' was created!\n", folder)
		}

		// Default return
		return nil
	})
}
