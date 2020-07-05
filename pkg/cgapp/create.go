package cgapp

import (
	"path/filepath"
	"regexp"
	"strings"
)

// Project struct for describe project
type Project struct {
	Name       string
	Type       string
	RootFolder string
}

// CreateProjectFromRegistry function for create a new project from repository
func CreateProjectFromRegistry(p *Project, registry map[string]string) error {
	// Define vars
	var pattern string

	// Create path in project root folder
	folder := filepath.Join(p.RootFolder, p.Type)

	// Switch project type
	switch p.Type {
	case "roles":
		pattern = regexpAnsiblePattern
		break
	case "backend":
		pattern = regexpBackendPattern
		break
	case "webserver":
		pattern = regexpWebServerPattern
		break
	case "database":
		pattern = regexpDatabasePattern
		break
	}

	// Create match expration
	match, err := regexp.MatchString(pattern, p.Name)
	if err != nil {
		return ThrowError(err.Error())
	}

	// Check for regexp
	if match {
		// If match, create from default template
		if err := GitClone(folder, registry[p.Name]); err != nil {
			return ThrowError(err.Error())
		}

		// Show success report
		SendMsg(false, "OK", strings.Title(p.Type)+" was created with default template `"+registry[p.Name]+"`!", "", false)
	} else {
		// Else create from user template (from GitHub, etc)
		if err := GitClone(folder, p.Name); err != nil {
			return ThrowError(err.Error())
		}

		// Show success report
		SendMsg(false, "OK", strings.Title(p.Type)+" was created with user template `"+p.Name+"`!", "", false)
	}

	// Clean
	if err := RemoveFolders(folder, []string{".git", ".github"}); err != nil {
		return ThrowError(err.Error())
	}

	return nil
}
