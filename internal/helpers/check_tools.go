package helpers

import (
	"errors"
	"fmt"
	"os/exec"
)

// CheckCLITools ...
func CheckCLITools(commands []string) error {
	// Create a new slice for join errors.
	errs := make([]error, 0)

	//
	for _, command := range commands {
		// Start execution command.
		cmd := exec.Command(command, "-v")

		//
		_, err := cmd.Output()
		if err != nil {
			errs = append(errs, fmt.Errorf("'%s' is required, but not installed on your system", command))
		}
	}

	return errors.Join(errs...)
}
