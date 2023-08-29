package helpers

import (
	"errors"
	"fmt"
	"os"
)

// CheckProjectStructure function for checking required project structure by names.
func CheckProjectStructure(elements []string) error {
	// Create a new slice for join errors.
	errs := make([]error, 0)

	// Loop for given elements.
	for _, element := range elements {
		// Start checking element.
		if _, err := os.Stat(element); err != nil {
			errs = append(errs, fmt.Errorf("'%s' folder/file is required, but not founded on the current dir", element))
		}
	}

	return errors.Join(errs...)
}
