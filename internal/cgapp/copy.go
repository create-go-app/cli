package cgapp

import (
	"fmt"
	"io"
	"os"

	"github.com/markbates/pkger"
)

// CopyFolder function for copy all files from app config
func CopyFolder(folderName string) error {
	return pkger.Walk(folderName, func(path string, file os.FileInfo, err error) error {
		// Define files paths
		folder := appPath + string(os.PathSeparator) + file.Name()

		// Create files (skip directories)
		if !file.IsDir() {
			// Open file
			from, err := pkger.Open(path)
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
