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
	SendMsg(true, "WAIT", "Create project folder and config files:", "cyan", false)
	if err := MakeFolder(appPath, 0750); err != nil {
		return ThrowError(err.Error())
	}

	// Create config files for app
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
		SendMsg(true, "WAIT", "Create Ansible playbook and roles for deploy", "cyan", false)

		// Create playbook
		filesToMake := map[string][]byte{
			"deploy-playbook.yml": embed.Get("/deploy-playbook.yml"),
		}
		if err := MakeFiles(appPath, filesToMake); err != nil {
			return ThrowError(err.Error())
		}

		// Roles
		ErrChecker(
			Create(&Config{
				Name:   "roles",
				Match:  "^(roles)$",
				View:   "roles",
				Folder: appPath,
			},
				registry,
			),
		)
	}

	// Create backend files
	SendMsg(true, "WAIT", "Create project backend:", "cyan", false)
	ErrChecker(
		Create(&Config{
			Name:   strings.ToLower(appBackend),
			Match:  "^(net/http|fiber|echo)$",
			View:   "backend",
			Folder: appPath,
		},
			registry,
		),
	)

	// Create frontend files
	if appFrontend != "none" {
		SendMsg(true, "WAIT", "Create project frontend:", "cyan", false)
		ErrChecker(
			Create(&Config{
				Name:   strings.ToLower(appFrontend),
				Match:  "^(preact|react-js|react-ts)$",
				View:   "frontend",
				Folder: appPath,
			},
				registry,
			),
		)
	}

	// Docker containers
	if appWebServer != "none" || appDatabase != "none" {

		SendMsg(true, "NEXT", "Configuring Docker containers...", "yellow", false)

		// Create container files
		if appWebServer != "none" {
			SendMsg(true, "WAIT", "Create container with web/proxy server:", "cyan", false)
			ErrChecker(
				Create(&Config{
					Name:   strings.ToLower(appWebServer),
					Match:  "^(nginx)$",
					View:   "webserver",
					Folder: appPath,
				},
					registry,
				),
			)
		}

		// Create container files
		if appDatabase != "none" {
			SendMsg(true, "WAIT", "Create container with database:", "cyan", false)
			ErrChecker(
				Create(&Config{
					Name:   strings.ToLower(appDatabase),
					Match:  "^(postgres)$",
					View:   "database",
					Folder: appPath,
				},
					registry,
				),
			)
		}
	}

	// Stop timer
	stopTimer := time.Since(startTimer).String()

	// END message
	SendMsg(true, "FINISH", "Completed in "+stopTimer+".", "green", true)
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
	SendMsg(true, "FINISH", "Completed in "+stopTimer+".", "green", true)
	SendMsg(true, "DOCS", "A helpful documentation here → https://create-go.app", "yellow", false)
	SendMsg(false, "!", "Go to the `"+deployHost+"` to see your deployed project! :)", "yellow", true)

	return nil
}
