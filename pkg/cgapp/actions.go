package cgapp

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/create-go-app/cli/internal/embed"
	"github.com/urfave/cli/v2"
)

// CreateCLIAction actions for `create` CLI command
func CreateCLIAction(c *cli.Context) error {
	// Start timer
	startTimer := time.Now()

	// START message
	SendMessage("\n[*] Create Go App v"+version, "yellow")
	SendMessage("\n[START] Creating a new app in `"+appPath+"` folder...", "green")

	// Create main folder for app
	SendMessage("\n[PROCESS] App folder and config files", "cyan")
	ErrChecker(os.Mkdir(appPath, 0750))
	SendMessage("[OK] App folder was created!", "")

	// Create config files for app
	ErrChecker(File(".gitignore", embed.Get("/.gitignore")))
	ErrChecker(File(".editorconfig", embed.Get("/.editorconfig")))
	ErrChecker(File("deploy-playbook.yml", embed.Get("/deploy-playbook.yml")))

	// Create backend files
	SendMessage("\n[PROCESS] App backend", "cyan")
	ErrChecker(
		Create(&Config{
			name:   strings.ToLower(appBackend),
			match:  "^(net/http|fiber|echo)$",
			view:   "backend",
			folder: appPath,
		},
			registry,
		),
	)

	// Create frontend files
	if appFrontend != "none" {
		SendMessage("\n[PROCESS] App frontend", "cyan")
		ErrChecker(
			Create(&Config{
				name:   strings.ToLower(appFrontend),
				match:  "^(preact|react-js|react-ts)$",
				view:   "frontend",
				folder: appPath,
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
					name:   strings.ToLower(appWebServer),
					match:  "^(nginx)$",
					view:   "webserver",
					folder: appPath,
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
					name:   strings.ToLower(appDatabase),
					match:  "^(postgres)$",
					view:   "database",
					folder: appPath,
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
	SendMessage("\n[!] Next steps & helpful instructions here → https://shrts.website/cgapp/faq", "yellow")
	SendMessage("[!] Go to the `"+appPath+"` folder and make something beautiful! :)\n", "yellow")

	return nil
}

// DeployCLIAction actions for `deploy` CLI command
func DeployCLIAction(c *cli.Context) error {
	// Start timer
	startTimer := time.Now()

	// START message
	SendMessage("\n[*] Create Go App v"+version, "yellow")
	SendMessage("\n[START] Deploying project to the `"+deployHost+"`...", "green")

	// Create main folder for app
	SendMessage("\n[PROCESS] Run Ansible playbook `"+deployPlaybook+"`", "cyan")

	// Check, if need to ask password for user
	// See: https://docs.ansible.com/ansible/latest/user_guide/become.html#become-command-line-options
	askBecomePass := ""
	if c.Bool("ask-become-pass") {
		askBecomePass = "--ask-become-pass" // #nosec G101
	}

	// Execute command
	cmd := exec.Command(
		"ansible-playbook",
		deployPlaybook,
		"-u "+deployUsername,
		"--extra-vars 'host="+deployHost+" network_name="+deployDockerNetwork+"'",
		askBecomePass,
	)
	ErrChecker(cmd.Run())

	// Stop timer
	stopTimer := time.Since(startTimer).String()

	// END message
	SendMessage("\n[DONE] Completed in "+stopTimer+".", "cyan")
	SendMessage("\n[!] Next steps & helpful instructions here → https://shrts.website/cgapp/faq", "yellow")
	SendMessage("[!] Go to the `"+deployHost+"` to see your deployed project! :)\n", "yellow")

	return nil
}
