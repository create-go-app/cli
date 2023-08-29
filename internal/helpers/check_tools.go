package helpers

import (
	"errors"
	"fmt"
	"os/exec"
)

// CheckCLITools function for checking required tools by names.
func CheckCLITools(commands []string) error {
	// Create a new slice for join errors.
	errs := make([]error, 0)

	// Loop for given commands.
	for _, command := range commands {
		// Start execution command.
		if err := exec.Command(command, "-v").Start(); err != nil {
			errs = append(errs, fmt.Errorf("'%s' is required, but not installed on your system", command))
		}
	}

	return errors.Join(errs...)
}
