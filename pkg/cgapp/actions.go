package cgapp

import (
	"strings"
	"time"

	"github.com/create-go-app/cli/internal/embed"
	"github.com/urfave/cli/v2"
)

// CreateCLIAction action for `create` CLI command
func CreateCLIAction(c *cli.Context) error {
	// Start timer
	startTimer := time.Now()

	// START message
	SendMsg(true, "*", "Create Go App v"+version, "yellow", false)
	SendMsg(true, "START", "Creating a new project in `"+appPath+"` folder...", "green", false)

	// Create main folder for app
	SendMsg(true, "WAIT", "Create project folder and config files", "cyan", false)
	if err := MakeFolder(appPath, 0750); err != nil {
		return ThrowError(err.Error())
	}

	// Create configuration files
	filesToMake := map[string][]byte{
		".gitignore":    embed.Get("/.gitignore"),
		".editorconfig": embed.Get("/.editorconfig"),
		"Taskfile.yml":  embed.Get("/Taskfile.yml"),
	}
	if err := MakeFiles(appPath, filesToMake); err != nil {
		return ThrowError(err.Error())
	}

	// Create Ansible playbook and download roles, if not skipped
	if !c.Bool("skip-ansible-roles") {
		SendMsg(true, "WAIT", "Create Ansible playbook and roles", "cyan", false)

		// Create playbook
		fileToMake := map[string][]byte{
			"deploy-playbook.yml": embed.Get("/deploy-playbook.yml"),
		}
		if err := MakeFiles(appPath, fileToMake); err != nil {
			return ThrowError(err.Error())
		}

		// Create Ansible roles
		if err := CreateProjectFromRegistry(
			&Project{Type: "roles", Name: "deploy", RootFolder: appPath},
			registry,
		); err != nil {
			return ThrowError(err.Error())
		}
	}

	// Create backend files
	SendMsg(true, "WAIT", "Create project backend", "cyan", false)
	if err := CreateProjectFromRegistry(
		&Project{Type: "backend", Name: strings.ToLower(appBackend), RootFolder: appPath},
		registry,
	); err != nil {
		return ThrowError(err.Error())
	}

	if appFrontend != "none" {
		// Create frontend files
		SendMsg(true, "WAIT", "Create project frontend", "cyan", false)
		if err := CreateProjectFromCMD(
			&Project{Type: "frontend", Name: strings.ToLower(appFrontend), RootFolder: appPath},
			cmds,
		); err != nil {
			return ThrowError(err.Error())
		}
	}

	// Docker containers
	if appWebServer != "none" || appDatabase != "none" {

		SendMsg(true, "NEXT", "Configuring Docker containers...", "yellow", false)

		if appWebServer != "none" {
			// Create container with web/proxy server
			SendMsg(true, "WAIT", "Create container with web/proxy server", "cyan", false)
			if err := CreateProjectFromRegistry(
				&Project{Type: "webserver", Name: strings.ToLower(appWebServer), RootFolder: appPath},
				registry,
			); err != nil {
				return ThrowError(err.Error())
			}
		}

		if appDatabase != "none" {
			// Create container with database
			SendMsg(true, "WAIT", "Create container with database", "cyan", false)
			if err := CreateProjectFromRegistry(
				&Project{Type: "database", Name: strings.ToLower(appDatabase), RootFolder: appPath},
				registry,
			); err != nil {
				return ThrowError(err.Error())
			}
		}
	}

	// Stop timer
	stopTimer := time.Since(startTimer).String()

	// END message
	SendMsg(true, "FINISH", "Completed in "+stopTimer+"!", "green", false)
	SendMsg(true, "DOCS", "A helpful documentation here → https://create-go.app", "yellow", false)
	SendMsg(false, "!", "Go to the `"+appPath+"` folder and make something beautiful! :)", "yellow", true)

	return nil
}

// DeployCLIAction action for `deploy` CLI command
func DeployCLIAction(c *cli.Context) error {
	// Start timer
	startTimer := time.Now()

	// START message
	SendMsg(true, "*", "Create Go App v"+version, "yellow", false)
	SendMsg(true, "START", "Deploying project to the `"+deployHost+"`...", "green", false)

	// Create main folder for app
	SendMsg(true, "WAIT", "Run Ansible playbook `"+deployPlaybook+"`", "cyan", true)

	// Collect options
	options := []string{
		deployPlaybook,
		"-u", deployUsername,
		"-e", "host=" + deployHost + " network_name=" + deployDockerNetwork,
	}

	// Check, if need to ask password for user
	// See: https://docs.ansible.com/ansible/latest/user_guide/become.html#become-command-line-options
	if c.Bool("ask-become-pass") {
		options = []string{
			deployPlaybook,
			"-u", deployUsername,
			"-e", "host=" + deployHost + " network_name=" + deployDockerNetwork,
			"--ask-become-pass",
		}
	}

	// Collect command line
	if err := ExecCommand("ansible-playbook", options); err != nil {
		return ThrowError(err.Error())
	}

	// Stop timer
	stopTimer := time.Since(startTimer).String()

	// END message
	SendMsg(true, "FINISH", "Completed in "+stopTimer+".", "green", false)
	SendMsg(true, "DOCS", "A helpful documentation here → https://create-go.app", "yellow", false)
	SendMsg(false, "!", "Go to the `"+deployHost+"` to see your deployed project! :)", "yellow", true)

	return nil
}
