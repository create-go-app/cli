package cgapp

import "os"

// ErrChecker function for check error
func ErrChecker(err error) {
	if err != nil {
		// Show error report
		SendMessage("[ERROR] "+err.Error(), "red")
		os.Exit(1)
	}
}
