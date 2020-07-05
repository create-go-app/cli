package cgapp

import (
	"path/filepath"
	"regexp"
	"strings"
)

// Project ...
type Project struct {
	Name       string
	Type       string
	RootFolder string
}

// CreateProjectFromRegistry function for create a new project from repository
func CreateProjectFromRegistry(p *Project, registry map[string]string) error {
	// Define vars
	var name, pattern string

	// Create path in project root folder
	folder := filepath.Join(p.RootFolder, p.Type)

	// Switch project type
	switch p.Type {
	case "roles":
		name = "roles"
		pattern = regexpAnsiblePattern
		break
	case "backend":
		name = p.Name
		pattern = regexpBackendPattern
		break
	case "webserver":
		name = p.Name
		pattern = regexpWebServerPattern
		break
	case "database":
		name = p.Name
		pattern = regexpDatabasePattern
		break
	}

	// Create match expration
	match, err := regexp.MatchString(pattern, name)
	if err != nil {
		return ThrowError(err.Error())
	}

	// Check for regexp
	if match {
		// If match, create from default template
		if err := GitClone(folder, registry[name]); err != nil {
			return ThrowError(err.Error())
		}

		// Show success report
		SendMsg(false, "OK", strings.Title(p.Type)+" was created with default template `"+registry[name]+"`!", "", false)
	} else {
		// Else create from user template (from GitHub, etc)
		if err := GitClone(folder, registry[name]); err != nil {
			return ThrowError(err.Error())
		}

		// Show success report
		SendMsg(false, "OK", strings.Title(p.Type)+" was created with user template `"+name+"`!", "", false)
	}

	// Clean
	if err := RemoveFolders(folder, []string{".git", ".github"}); err != nil {
		return ThrowError(err.Error())
	}

	return nil
}
