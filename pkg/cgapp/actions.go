package cgapp

import (
	"bufio"
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
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
	SendMessage("\n[*] Create Go App v"+version, "yellow")
	SendMessage("\n[START] Creating a new project in `"+appPath+"` folder...", "green")

	// Create main folder for app
	SendMessage("\n[PROCESS] Project folder and config files", "cyan")
	ErrChecker(os.Mkdir(appPath, 0750))
	SendMessage("[OK] Project folder was created!", "")

	// Create config files for app
	ErrChecker(File(".gitignore", embed.Get("/.gitignore")))
	ErrChecker(File(".editorconfig", embed.Get("/.editorconfig")))
	ErrChecker(File("Taskfile.yml", embed.Get("/Taskfile.yml")))

	// Create Ansible playbook and download roles, if not skipped
	if !c.Bool("skip-ansible-roles") {
		SendMessage("\n[PROCESS] Ansible playbook and roles for deploy", "cyan")

		// Playbook
		ErrChecker(File("deploy-playbook.yml", embed.Get("/deploy-playbook.yml")))

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
	SendMessage("\n[PROCESS] Project backend", "cyan")
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
		SendMessage("\n[PROCESS] Project frontend", "cyan")
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

		// Install dependencies
		SendMessage("\n[PROCESS] Frontend dependencies", "cyan")
		SendMessage("[WAIT] Installing frontend dependencies (may take some time)!", "yellow")

		// Go to ./frontend folder and run npm install
		cmd := exec.Command("npm", "install")
		cmd.Dir = filepath.Join(appPath, "frontend")
		ErrChecker(cmd.Run())

		SendMessage("[OK] Frontend dependencies was installed!", "green")
	}

	// Docker containers
	if appWebServer != "none" || appDatabase != "none" {

		SendMessage("\n[NEXT] Docker containers...", "green")

		// Create container files
		if appWebServer != "none" {
			SendMessage("\n[PROCESS] Web/proxy server container", "cyan")
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
			SendMessage("\n[PROCESS] Database container", "cyan")
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
	SendMessage("\n[DONE] Completed in "+stopTimer+".", "cyan")
	SendMessage("\n[!] Helpful instructions here → https://create-go.app/detailed-guides/", "yellow")
	SendMessage("[!] Go to the `"+appPath+"` folder and make something beautiful! :)\n", "yellow")

	return nil
}

// DeployCLIAction action for `deploy` CLI command
func DeployCLIAction(c *cli.Context) error {
	// Start timer
	startTimer := time.Now()

	// START message
	SendMessage("\n[*] Create Go App v"+version, "yellow")
	SendMessage("\n[START] Deploying project to the `"+deployHost+"`...", "green")

	// Create main folder for app
	SendMessage("\n[PROCESS] Run Ansible playbook `"+deployPlaybook+"`\n", "cyan")

	// Check, if need to ask password for user
	// See: https://docs.ansible.com/ansible/latest/user_guide/become.html#become-command-line-options
	askBecomePass := ""
	if c.Bool("ask-become-pass") {
		askBecomePass = "--ask-become-pass" // #nosec G101
	}

	// Create buffer for stderr
	stderr := &bytes.Buffer{}

	// Collect command line
	cmd := exec.Command(
		"ansible-playbook",
		deployPlaybook,
		"-u",
		deployUsername,
		"-e",
		"host="+deployHost+" network_name="+deployDockerNetwork,
		askBecomePass,
	)

	// Set buffer for stderr from cmd
	cmd.Stderr = stderr

	// Create a new reader
	cmdReader, err := cmd.StdoutPipe()
	ErrChecker(err)

	// Create a new scanner and run goroutine func with output
	scanner := bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan() {
			SendMessage(scanner.Text(), "")
		}
	}()

	// Run executing command
	if err := cmd.Run(); err != nil {
		SendMessage(stderr.String(), "red")
	}

	// Stop timer
	stopTimer := time.Since(startTimer).String()

	// END message
	SendMessage("[DONE] Completed in "+stopTimer+".", "cyan")
	SendMessage("\n[!] Helpful instructions here → https://create-go.app/detailed-guides/", "yellow")
	SendMessage("[!] Go to the `"+deployHost+"` to see your deployed project! :)\n", "yellow")

	return nil
}
