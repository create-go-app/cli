package cgapp

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/posener/gitfs"
	"github.com/posener/gitfs/fsutil"
)

// CopyFolderFromGit function for get all files from git repository
func CopyFolderFromGit(repositoryName, folderName string) error {
	// Define repository
	repository := filepath.Join(repositoryName, folderName)

	// Define folder path
	folder := filepath.Join(appPath)

	// Define context
	ctx := context.Background()

	// Create filesystem from repository
	fs, err := gitfs.New(ctx, repository)
	ErrChecker(err)

	// Create walker for filesystem
	walker := fsutil.Walk(fs, "")

	// Walk for each file
	for walker.Step() {
		// Error report
		ErrChecker(walker.Err())

		// Re-define each file path
		folder = filepath.Join(appPath, walker.Path())

		// If not directory, create file
		if !walker.Stat().IsDir() {
			// Open file
			from, err := fs.Open(walker.Path())
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
	}

	// Default return
	return nil
}
