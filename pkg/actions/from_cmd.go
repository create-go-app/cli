package actions

import (
	"os"
	"path/filepath"

	"github.com/create-go-app/cli/pkg/registry"
	"github.com/create-go-app/cli/pkg/utils"
)

// CreateProjectFromCMD function for create a new project from a comand line.
func CreateProjectFromCMD(p *registry.Project, c map[string]*registry.Command) {
	// Define vars.
	var options []string

	// Checking for nil.
	if p == nil || c == nil {
		utils.SendMsg(true, "[ERROR]", "Project template or commands not found!", "red", true)
		os.Exit(1)
	}

	// Create path in project root folder.
	folder := filepath.Join(p.RootFolder, p.Type)

	// Split framework name and template.
	project, err := utils.StringSplit(":", p.Name)
	if err != nil {
		utils.SendMsg(true, "[ERROR]", err.Error(), "red", true)
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
		utils.SendMsg(true, "[ERROR]", "Frontend template"+p.Name+" not found!", "red", true)
		os.Exit(1)
	}

	// Run execution command.
	if err := utils.ExecCommand(runner, options); err != nil {
		utils.SendMsg(true, "[ERROR]", err.Error(), "red", true)
		os.Exit(1)
	}

	// Cleanup project.
	folderToRemove := []string{".git", ".github"}
	utils.RemoveFolders(folder, folderToRemove)

	// Show success report.
	utils.SendMsg(false, "OK", "Frontend: created with template `"+p.Name+"`!", "", false)
}
