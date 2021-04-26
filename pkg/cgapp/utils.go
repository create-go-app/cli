// Copyright 2019-present Vic Shóstak. All rights reserved.
// Use of this source code is governed by Apache 2.0 license
// that can be found in the LICENSE file.

package cgapp

import (
	"fmt"
	"regexp"
)

// BeautifyText function for send (colored or common) message to output.
func BeautifyText(text, color string) string {
	// Define variables.
	var (
		red       string = "\033[0;31m"
		green     string = "\033[0;32m"
		cyan      string = "\033[0;36m"
		yellow    string = "\033[1;33m"
		noColor   string = "\033[0m"
		textColor string
	)

	// Switch color.
	switch color {
	case "":
		textColor = noColor
	case "green":
		textColor = green
	case "yellow":
		textColor = yellow
	case "red":
		textColor = red
	case "cyan":
		textColor = cyan
	}

	// Send common or colored text.
	return textColor + text + noColor
}

// SendMsg function forsend message to output.
func SendMsg(startWithNewLine bool, caption, text, color string, endWithNewLine bool) {
	// Define variables.
	var startNewLine, endNewLine string

	if startWithNewLine {
		startNewLine = "\n" // set new line
	}

	if endWithNewLine {
		endNewLine = "\n" // set new line
	}

	if caption == "" {
		fmt.Println(startNewLine + text + endNewLine) // common text
	} else {
		fmt.Println(startNewLine + BeautifyText(caption, color) + " " + text + endNewLine) // colorized text
	}
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
