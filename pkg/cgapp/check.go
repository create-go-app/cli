package cgapp

// ErrChecker function for easily checking error,
// instead of err != nil {...}
func ErrChecker(err error) {
	if err != nil {
		// Show error report
		SendMessage("[ERROR] "+err.Error(), "red")
		return
	}
}
