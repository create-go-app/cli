package cgapp

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/go-git/go-git/v5"
)

// Config struct for project configuration
type Config struct {
	Name   string
	Match  string
	View   string
	Folder string
}

// Create function for create app
func Create(c *Config, registry map[string]string) error {
	// Create path to folder
	folder := filepath.Join(c.Folder, c.View)

	// Create match expration for frameworks/containers
	match, err := regexp.MatchString(c.Match, c.Name)
	ErrChecker(err)

	// Check for regexp
	if match {
		// If match, create from default template
		_, err := git.PlainClone(folder, false, &git.CloneOptions{
			URL:      "https://github.com/" + registry[c.Name],
			Progress: os.Stdout,
		})
		ErrChecker(err)

		// Show success report
		SendMessage(
			"[OK] "+strings.Title(c.View)+" was created with default template `"+registry[c.Name]+"`!",
			"green",
		)
	} else {
		// Else create from user template (from GitHub, etc)
		_, err := git.PlainClone(folder, false, &git.CloneOptions{
			URL:      "https://" + c.Name,
			Progress: os.Stdout,
		})
		ErrChecker(err)

		// Show success report
		SendMessage(
			"[OK] "+strings.Title(c.View)+" was created with user template `"+c.Name+"`!",
			"green",
		)
	}

	// Clean
	ErrChecker(os.RemoveAll(filepath.Join(folder, ".git")))
	ErrChecker(os.RemoveAll(filepath.Join(folder, ".github")))

	return nil
}
