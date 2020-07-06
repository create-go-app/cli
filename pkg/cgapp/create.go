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
		SendMsg(false, "OK", strings.Title(p.Type)+": created with default template `"+registry[p.Name]+"`!", "", false)
	} else {
		// Else create from user template (from GitHub, etc)
		if err := GitClone(folder, p.Name); err != nil {
			return ThrowError(err.Error())
		}

		// Show success report
		SendMsg(false, "OK", strings.Title(p.Type)+": created with user template `"+p.Name+"`!", "", false)
	}

	return nil
}

// CreateProjectFromCMD ...
func CreateProjectFromCMD(p *Project, cmd map[string]map[string]string) error {
	// Define vars
	var options []string

	// Create path in project root folder
	folder := filepath.Join(p.RootFolder, p.Type)

	// Split framework name and template
	project := StringSplit(":", p.Name)

	// Error, when empty
	if len(project) == 0 {
		return ThrowError("Frontend template not set!")
	}

	// Collect project runner and options
	switch project[0] {
	case "react":
		options = []string{cmd["react"]["create"], folder}
		if len(project) > 1 {
			options = []string{cmd["react"]["create"], folder, cmd["react"]["template"], "cra-template-" + project[1]}
		}
		break
	default:
		return ThrowError("Frontend template" + p.Name + " not found!")
	}

	//
	if err := ExecCommand(cmd[project[0]]["runner"], options); err != nil {
		return ThrowError(err.Error())
	}

	return nil
}
