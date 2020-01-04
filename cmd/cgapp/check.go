package cgapp

import (
	"fmt"
	"os"
)

// ErrChecker function for check error
func ErrChecker(err error) {
	if err != nil {
		// Show error report
		fmt.Printf("\n%v[✘] Error: %v%v\n\n", red, err, noColor)
		os.Exit(1)
	}
}
