// Copyright 2022 Vic Shóstak and Create Go App Contributors. All rights reserved.
// Use of this source code is governed by Apache 2.0 license
// that can be found in the LICENSE file.

package cgapp

import (
	"fmt"
	"time"

	"github.com/mattn/go-colorable"
)

var (
	Stdout = colorable.NewColorableStdout() // add a colorable std out
	Stderr = colorable.NewColorableStderr() // add a colorable std err
)

// ShowMessage function for showing output messages.
func ShowMessage(level, text string, startWithNewLine, endWithNewLine bool) {
	// Define variables.
	var startLine, endLine string

	if startWithNewLine {
		startLine = "\n" // set a new line
	}

	if endWithNewLine {
		endLine = "\n" // set a new line
	}

	// Formatting message.
	message := fmt.Sprintf("%s %s %s %s", startLine, colorizeLevel(level), text, endLine)

	// Return output.
	_, err := fmt.Fprintln(Stdout, message)
	if err != nil {
		return
	}
}

// ShowError function for send error message to output.
func ShowError(text string) error {
	return fmt.Errorf("%s%s", colorizeLevel("error"), text)
}

// CalculateDurationTime func to calculate duration time.
func CalculateDurationTime(startTimer time.Time) string {
	return fmt.Sprintf("%.0f", time.Since(startTimer).Seconds())
}

// colorizeLevel function for send (colored or common) message to output.
func colorizeLevel(level string) string {
	// Define variables.
	var (
		red         = "\033[0;31m"
		green       = "\033[0;32m"
		yellow      = "\033[1;33m"
		noColor     = "\033[0m"
		color, icon string
	)

	// Switch color.
	switch level {
	case "success":
		color = green
		icon = "[OK]"
	case "error":
		color = red
		icon = "[ERROR]"
	case "info":
		color = yellow
		icon = "[INFO]"
	default:
		color = noColor
	}

	// Send common or colored caption.
	return fmt.Sprintf("%s%s%s", color, icon, noColor)
}
