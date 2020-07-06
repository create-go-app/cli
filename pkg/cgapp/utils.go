package cgapp

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"

	"github.com/go-git/go-git/v5"
)

// BeautifyText function for send (colored or common) message to output
func BeautifyText(text, color string) string {
	// Define variables
	var (
		red       string = "\033[0;31m"
		green     string = "\033[0;32m"
		cyan      string = "\033[0;36m"
		yellow    string = "\033[1;33m"
		noColor   string = "\033[0m"
		textColor string
	)

	// Switch color
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

	// Send common or colored text
	return textColor + text + noColor
}

// SendMsg ...
func SendMsg(startWithNewLine bool, caption, text, color string, endWithNewLine bool) {
	var startNewLine, endNewLine string

	if startWithNewLine {
		startNewLine = "\n"
	}

	if endWithNewLine {
		endNewLine = "\n"
	}

	fmt.Println(BeautifyText(startNewLine+"["+caption+"] "+text+endNewLine, color))
}

// ThrowError ...
func ThrowError(text string) error {
	return fmt.Errorf(BeautifyText(text, "red"))
}

// ExecCommand ...
func ExecCommand(command string, options []string) error {
	//
	if command == "" {
		return ThrowError("No command to execute!")
	}

	// Create buffer for stderr
	stderr := &bytes.Buffer{}

	// Collect command line
	cmd := exec.Command(command, options...) // #nosec G204

	// Set buffer for stderr from cmd
	cmd.Stderr = stderr

	// Create a new reader
	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		return ThrowError(err.Error())
	}

	// Start executing command
	if err := cmd.Start(); err != nil {
		return ThrowError(stderr.String())
	}

	// Create a new scanner and run goroutine func with output
	scanner := bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan() {
			SendMsg(false, "~", scanner.Text(), "", false)
		}
	}()

	// Wait for executing command
	if err := cmd.Wait(); err != nil {
		return ThrowError(stderr.String())
	}

	return nil
}

// StringSplit ...
func StringSplit(pattern, match string) ([]string, error) {
	// Error, when empty or nil
	if pattern == "" || match == "" {
		return nil, ThrowError("Frontend template not set!")
	}

	// Define empty []string{} for splitted strings
	splittedStrings := []string{}

	// Create regexp
	re := regexp.MustCompile(pattern)

	// Split match string
	split := re.Split(match, -1)
	for str := range split {
		// Append all matched strings to set
		splittedStrings = append(splittedStrings, split[str])
	}

	return splittedStrings, nil
}

// MakeFiles function for massively create folders
func MakeFiles(rootFolder string, filesToMake map[string][]byte) error {
	for file, data := range filesToMake {
		folder := filepath.Join(rootFolder, file)

		// Write to created file
		if err := ioutil.WriteFile(folder, data, 0755); err != nil {
			return ThrowError("File `" + file + "` was not created!")
		}

		// Show report for file
		SendMsg(false, "OK", "File `"+file+"` was created!", "", false)
	}

	return nil
}

// MakeFolder function for create folder
func MakeFolder(folderName string, chmod os.FileMode) error {
	// Check if folder exists, fail if it does
	if _, err := os.Stat(folderName); !os.IsNotExist(err) {
		return ThrowError("Folder `" + folderName + "` exists!")
	}

	// Create folder
	if err := os.Mkdir(folderName, chmod); err != nil {
		return ThrowError("Folder `" + folderName + "` was not created!")
	}

	// Show report for folder
	SendMsg(false, "OK", "Folder `"+folderName+"` was created!", "", false)

	return nil
}

// RemoveFolders ...
func RemoveFolders(rootFolder string, foldersToRemove []string) error {
	for _, folder := range foldersToRemove {
		if err := os.RemoveAll(filepath.Join(rootFolder, folder)); err != nil {
			return ThrowError(err.Error())
		}
	}

	return nil
}

// GitClone function for `git clone` defined project template
func GitClone(rootFolder, templateName string) error {
	// Clone project template
	_, err := git.PlainClone(rootFolder, false, &git.CloneOptions{
		URL: "https://" + templateName,
	})
	if err != nil {
		return ThrowError("Repository was not cloned!")
	}

	return nil
}
