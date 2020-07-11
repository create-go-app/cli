/*
Package cgapp includes a powerful CLI for the Create Go App project.

Create a new production-ready project with backend (Golang),
frontend (JavaScript, TypeScript) and deploy automation
(Ansible, Docker) by running one CLI command.

-> Focus on writing code and thinking of business logic!
<- The Create Go App CLI will take care of the rest.

A helpful documentation and next steps -> https://create-go.app/

Copyright © 2019-present Vic Shóstak <truewebartisans@gmail.com> (https://1wa.co)

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
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
		break
	case "green":
		textColor = green
		break
	case "yellow":
		textColor = yellow
		break
	case "red":
		textColor = red
		break
	case "cyan":
		textColor = cyan
		break
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
		return nil, throwError("Frontend template not set!")
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

// throwError function for throw an error.
func throwError(text string) error {
	return fmt.Errorf(BeautifyText(text, "red"))
}
