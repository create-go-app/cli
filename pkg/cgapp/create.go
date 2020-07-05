package cgapp

import (
	"path/filepath"
	"regexp"
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
	if err != nil {
		return ThrowError(err.Error())
	}

	// Check for regexp
	if match {
		// If match, create from default template
		if err := GitClone(folder, "https://github.com/", registry[c.Name]); err != nil {
			return ThrowError(err.Error())
		}

		// Show success report
		SendMsg(false, "OK", "Backend was created with default template `"+registry[c.Name]+"`!", "", false)
	} else {
		// Else create from user template (from GitHub, etc)
		if err := GitClone(folder, "https://", registry[c.Name]); err != nil {
			return ThrowError(err.Error())
		}

		// Show success report
		SendMsg(false, "OK", "Backend was created with user template `"+registry[c.Name]+"`!", "", false)
	}

	// Clean
	if err := RemoveFolders(folder, []string{".git", ".github"}); err != nil {
		return ThrowError(err.Error())
	}

	return nil
}
