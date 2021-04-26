// Copyright 2019-present Vic Shóstak. All rights reserved.
// Use of this source code is governed by Apache 2.0 license
// that can be found in the LICENSE file.

package cgapp

import (
	"fmt"
	"regexp"
)

// ShowMessage function for send message to output.
func ShowMessage(level, text string, startWithNewLine, endWithNewLine bool) string {
	// Define variables.
	var startNewLine, endNewLine string

	if startWithNewLine {
		startNewLine = "\n" // set a new line
	}

	if endWithNewLine {
		endNewLine = "\n" // set a new line
	}

	return startNewLine + colorizeLevel(level) + text + endNewLine
}

// colorizeLevel function for send (colored or common) message to output.
func colorizeLevel(level string) string {
	// Define variables.
	var (
		red         string = "\033[0;31m"
		green       string = "\033[0;32m"
		cyan        string = "\033[0;36m"
		yellow      string = "\033[1;33m"
		noColor     string = "\033[0m"
		color, icon string
	)

	// Switch color.
	switch level {
	case "success":
		color = green
		icon = "[OK] "
	case "warning":
		color = yellow
		icon = "[WARNING] "
	case "error":
		color = red
		icon = "[ERROR] "
	case "info":
		color = cyan
		icon = "[INFO] "
	default:
		color = noColor
	}

	// Send common or colored caption.
	return color + icon + noColor
}

// stringSplit function for split string by pattern.
func stringSplit(pattern, match string) ([]string, error) {
	// Error, when empty or nil.
	if pattern == "" || match == "" {
		return nil, fmt.Errorf("Frontend template not set!")
	}

	// Define empty []string{} for splitted strings.
	splittedStrings := []string{}

	// Create regexp.
	re := regexp.MustCompile(pattern)

	// Split match string.
	split := re.Split(match, -1)
	for str := range split {
		// Append all matched strings to set.
		splittedStrings = append(splittedStrings, split[str])
	}

	return splittedStrings, nil
}
