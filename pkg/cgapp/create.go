// Copyright 2019-present Vic Shóstak. All rights reserved.
// Use of this source code is governed by Apache 2.0 license
// that can be found in the LICENSE file.

package cgapp

import (
	"path/filepath"
	"regexp"
	"strings"

	"github.com/create-go-app/cli/pkg/registry"
)

// CreateProjectFromRegistry function for create a new project from repository.
func CreateProjectFromRegistry(p *registry.Project, r map[string]*registry.Repository, m string) error {
	// Define vars.
	var template string

	// Checking for nil.
	if p == nil || r == nil || m == "" {
		return throwError("Project template, registry or pattern not found!")
	}

	// Create path in project root folder.
	folder := filepath.Join(p.RootFolder, p.Type)

	// Create match expration.
	match, err := regexp.MatchString(m, p.Name)
	if err != nil {
		return throwError(err.Error())
	}

	// Check for regexp.
	if match {
		// Re-define vars.
		template = r[p.Type].List[p.Name]

		// If match, create from default template.
		if err := GitClone(folder, template); err != nil {
			return throwError(err.Error())
		}
	} else {
		// Re-define vars.
		template = p.Name

		// Else create from user template (from GitHub, etc).
		if err := GitClone(folder, template); err != nil {
			return throwError(err.Error())
		}
	}

	// Show success report.
	SendMsg(false, "[OK]", strings.Title(p.Type)+": created with the `"+template+"` template!", "cyan", false)

	// Cleanup project.
	foldersToRemove := []string{".git", ".github"}
	RemoveFolders(folder, foldersToRemove)

	return nil
}

// CreateProjectFromCmd function for create a new project from a comand line.
func CreateProjectFromCmd(p *registry.Project, c map[string]*registry.Command, m string) error {
	// Define vars.
	var options []string

	// Checking for nil.
	if p == nil || c == nil || m == "" {
		return throwError("Project template, commands or pattern not found!")
	}

	// Create path in project root folder.
	folder := filepath.Join(p.RootFolder, p.Type)

	// Create match expration for name.
	match, err := regexp.MatchString(m, p.Name)
	if err != nil {
		return throwError(err.Error())
	}

	if match {
		// Split frontend library/framework name and template.
		project, err := stringSplit(":", p.Name)
		if err != nil {
			return throwError(err.Error())
		}

		// Re-define vars for more beauty view.
		runner := c[project[0]].Runner
		create := c[project[0]].Create
		args := c[project[0]].Args

		// Collect project runner and options.
		switch project[0] {
		case "react":
			// npx create-react-app [template]
			options = []string{create, folder}
			if len(project) > 1 {
				options = []string{create, folder, args["template"], "cra-template-" + project[1]}
			}
		case "preact":
			// preact create [template] [dest] [args...]
			options = []string{create, "default", p.Type, args["cwd"], p.RootFolder, args["name"], "cgapp"}
			if len(project) > 1 {
				options = []string{create, project[1], p.Type, args["cwd"], p.RootFolder, args["name"], "cgapp"}
			}
		case "vue":
			// vue create [options] <app-name>
			options = []string{create, "--default", "--bare", p.Type}
			if len(project) == 2 {
				options = []string{create, "--preset", project[1], "--bare", p.Type}
			}
			if len(project) == 3 {
				options = []string{create, "--preset", project[1] + ":" + project[2], "--bare", "--clone", p.Type}
			}
		case "angular":
			// ng new <app-name> [options]
			options = []string{create, "cgapp", "--defaults", "--routing", "--directory", p.Type}
		case "svelte":
			// npx degit [template] [dest]
			options = []string{create, args["template"], folder}
		case "sapper":
			// npx degit [template] [dest]
			options = []string{create, args["template"] + "#rollup", folder}
			if len(project) > 1 {
				options = []string{create, args["template"] + "#" + project[1], folder}
			}
		}

		// Run execution command.
		if err := ExecCommand(runner, options); err != nil {
			return throwError(err.Error())
		}
	} else {
		// Create frontend from given repository (GitHub, etc).
		if err := GitClone(folder, p.Name); err != nil {
			return throwError(err.Error())
		}
	}

	// Cleanup project.
	folderToRemove := []string{".git", ".github"}
	RemoveFolders(folder, folderToRemove)

	// Show success report.
	SendMsg(true, "[OK]", "Frontend: created with template `"+p.Name+"`!", "cyan", false)

	return nil
}
