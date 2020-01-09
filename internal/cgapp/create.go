package cgapp

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"

	"github.com/markbates/pkger"
	"gopkg.in/src-d/go-git.v4"
)

// embedConfig struct for embed configuration
type embedConfig struct {
	embedFolder string
	appFolder   string
}

// appConfig struct for app configuration
type appConfig struct {
	name   string
	match  string
	view   string
	folder string
}

// CreateConfig function for create app's config files from embed folder
func CreateConfig(e *embedConfig) error {
	return pkger.Walk(e.embedFolder, func(path string, info os.FileInfo, err error) error {
		// Define files paths
		folder := e.appFolder + string(os.PathSeparator) + info.Name()

		// Create files
		if !info.IsDir() {
			// Open file from embed binary
			from, err := pkger.Open(path)
			ErrChecker(err)
			defer from.Close()

			// Create file
			to, err := os.Create(folder)
			ErrChecker(err)
			defer to.Close()

			// Copy data from embed binary to real file
			_, err = io.Copy(to, from)
			ErrChecker(err)

			// Show report for each file
			fmt.Printf("— File '%v' was created!\n", folder)
		}

		// Default return
		return nil
	})
}

// CreateApp function for create app
func CreateApp(c *appConfig, registry map[string]string) error {
	// Create path to folder
	folder := c.folder + string(os.PathSeparator) + c.view

	// Create match expration for frameworks/containers
	match, _ := regexp.MatchString(c.match, c.name)
	if match {
		// If match, create from default template
		_, err := git.PlainClone(folder, false, &git.CloneOptions{
			URL:      "https://github.com/" + registry[c.name],
			Progress: os.Stdout,
		})
		ErrChecker(err)

		// Clean
		os.RemoveAll(folder + string(os.PathSeparator) + ".git")

		// Show success report
		fmt.Printf(
			"\n%v[✔️]%v %v was created with default template '%v'!\n",
			green, noColor,
			strings.Title(c.view),
			registry[c.name],
		)
	} else {
		// Else create from user template (from GitHub, etc)
		_, err := git.PlainClone(folder, false, &git.CloneOptions{
			URL:      "https://" + c.name,
			Progress: os.Stdout,
		})
		ErrChecker(err)

		// Clean
		os.RemoveAll(folder + string(os.PathSeparator) + ".git")

		// Show success report
		fmt.Printf(
			"\n%v[✔️]%v %v was created with user template '%v'!\n",
			green, noColor,
			strings.Title(c.view),
			c.name,
		)
	}

	// Default return
	return nil
}
