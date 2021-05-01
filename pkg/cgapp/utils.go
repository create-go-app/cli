// Copyright 2019-present Vic Shóstak. All rights reserved.
// Use of this source code is governed by Apache 2.0 license
// that can be found in the LICENSE file.

package cgapp

import (
	"fmt"
	"time"
)

// ShowMessage function for send message to output.
func ShowMessage(level, text string, startWithNewLine, endWithNewLine bool) {
	// Define variables.
	var startLine, endLine string

	if startWithNewLine {
		startLine = "\n" // set a new line
	}

	if endWithNewLine {
		endLine = "\n" // set a new line
	}

	fmt.Println(startLine + colorizeLevel(level) + text + endLine)
}

// ShowError function for send error message to output.
func ShowError(text string) error {
	return fmt.Errorf(colorizeLevel("error") + text)
}

// CalculateDurationTime func to calculate duration time.
func CalculateDurationTime(startTimer time.Time) string {
	return fmt.Sprintf("%.0f", time.Since(startTimer).Seconds())
}

// colorizeLevel function for send (colored or common) message to output.
func colorizeLevel(level string) string {
	// Define variables.
	var (
		red         string = "\033[0;31m"
		green       string = "\033[0;32m"
		yellow      string = "\033[1;33m"
		noColor     string = "\033[0m"
		color, icon string
	)

	// Switch color.
	switch level {
	case "success":
		color = green
		icon = "[OK] "
	case "error":
		color = red
		icon = "[ERROR] "
	case "info":
		color = yellow
		icon = "[INFO] "
	default:
		color = noColor
	}

	// Send common or colored caption.
	return color + icon + noColor
}
