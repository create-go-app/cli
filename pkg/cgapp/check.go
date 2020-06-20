package cgapp

// ErrChecker function for check error
func ErrChecker(err error) {
	if err != nil {
		// Show error report
		SendMessage("[ERROR] "+err.Error(), "red")
		return
	}
}
