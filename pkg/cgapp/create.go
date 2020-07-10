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
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/create-go-app/cli/pkg/registry"
)

// CreateProjectFromRegistry function for create a new project from repository.
func CreateProjectFromRegistry(p *registry.Project, r map[string]*registry.Repository) {
	// Define vars.
	var pattern string

	// Checking for nil.
	if p == nil || r == nil {
		SendMsg(true, "[ERROR]", "Project template or registry not found!", "red", true)
		os.Exit(1)
	}

	// Create path in project root folder.
	folder := filepath.Join(p.RootFolder, p.Type)

	// Switch project type.
	switch p.Type {
	case "roles":
		pattern = registry.RegexpAnsiblePattern
		folder = filepath.Join(p.RootFolder, p.Type, p.Name) // re-define folder
		break
	case "backend":
		pattern = registry.RegexpBackendPattern
		break
	case "webserver":
		pattern = registry.RegexpWebServerPattern
		break
	case "database":
		pattern = registry.RegexpDatabasePattern
		break
	}

	// Create match expration.
	match, err := regexp.MatchString(pattern, p.Name)
	if err != nil {
		SendMsg(true, "[ERROR]", err.Error(), "red", true)
		os.Exit(1)
	}

	// Check for regexp.
	if match {
		// Re-define vars.
		template := r[p.Type].List[p.Name]

		// If match, create from default template.
		if err := GitClone(folder, template); err != nil {
			SendMsg(true, "[ERROR]", err.Error(), "red", true)
			os.Exit(1)
		}

		// Show success report.
		SendMsg(false, "[OK]", strings.Title(p.Type)+": created with default template `"+template+"`!", "cyan", false)
	} else {
		// Else create from user template (from GitHub, etc).
		if err := GitClone(folder, p.Name); err != nil {
			SendMsg(true, "[ERROR]", err.Error(), "red", true)
			os.Exit(1)
		}

		// Show success report.
		SendMsg(false, "[OK]", strings.Title(p.Type)+": created with user template `"+p.Name+"`!", "cyan", false)
	}

	// Cleanup project.
	foldersToRemove := []string{".git", ".github"}
	RemoveFolders(folder, foldersToRemove)
}

// CreateProjectFromCMD function for create a new project from a comand line.
func CreateProjectFromCMD(p *registry.Project, c map[string]*registry.Command) {
	// Define vars.
	var options []string

	// Checking for nil.
	if p == nil || c == nil {
		SendMsg(true, "[ERROR]", "Project template or commands not found!", "red", true)
		os.Exit(1)
	}

	// Create path in project root folder.
	folder := filepath.Join(p.RootFolder, p.Type)

	// Split framework name and template.
	project, err := StringSplit(":", p.Name)
	if err != nil {
		SendMsg(true, "[ERROR]", err.Error(), "red", true)
		os.Exit(1)
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
		break
	case "preact":
		// preact create [template] [dest] [args...]
		options = []string{create, "default", p.Type, args["cwd"], p.RootFolder, args["name"], "cgapp"}
		if len(project) > 1 {
			options = []string{create, project[1], p.Type, args["cwd"], p.RootFolder, args["name"], "cgapp"}
		}
		break
	case "svelte":
		// npx degit [template] [dest]
		options = []string{create, args["template"], folder}
		break
	default:
		SendMsg(true, "[ERROR]", "Frontend template"+p.Name+" not found!", "red", true)
		os.Exit(1)
	}

	// Run execution command.
	if err := ExecCommand(runner, options); err != nil {
		SendMsg(true, "[ERROR]", err.Error(), "red", true)
		os.Exit(1)
	}

	// Cleanup project.
	folderToRemove := []string{".git", ".github"}
	RemoveFolders(folder, folderToRemove)

	// Show success report.
	SendMsg(false, "[OK]", "Frontend: created with template `"+p.Name+"`!", "cyan", false)
}
