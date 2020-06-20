package cgapp

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/create-go-app/cli/cmd/box"
	"github.com/urfave/cli/v2"
)

// CreateCLIAction actions for `create` CLI command
func CreateCLIAction(c *cli.Context) error {
	// START message
	SendMessage("[*] Create Go App v"+version, "yellow")
	SendMessage("\n[START] Creating a new app...", "green")

	// Create main folder for app
	SendMessage("\n[PROCESS] App folder and config files", "cyan")
	ErrChecker(os.Mkdir(appPath, 0750))
	SendMessage("[OK] App folder was created!", "")

	// Create config files for app
	ErrChecker(File(".editorconfig", box.Get("/dotfiles/.editorconfig")))
	ErrChecker(File(".gitignore", box.Get("/dotfiles/.gitignore")))
	ErrChecker(File("Makefile", box.Get("/dotfiles/Makefile")))

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
	SendMessage("\n[START] Configuring Docker containers...", "green")

	// Check ./frontend folder
	SendMessage("\n[PROCESS] File docker-compose.yml", "cyan")
	_, err := os.Stat(filepath.Join(appPath, "frontend"))
	if !os.IsNotExist(err) {
		// If exists, create fullstack app docker-compose override file
		ErrChecker(
			File(
				"docker-compose.yml",
				box.Get("/docker/docker-compose.fullstack.yml"),
			),
		)
	} else {
		// Default docker-compose.yml
		ErrChecker(
			File(
				"docker-compose.yml",
				box.Get("/docker/docker-compose.backend.yml"),
			),
		)
	}

	// Production settings docker-compose.prod.yml
	ErrChecker(
		File(
			"docker-compose.prod.yml",
			box.Get("/docker/docker-compose.fullstack.yml"),
		),
	)

	// Create container files
	SendMessage("\n[PROCESS] Web/proxy server", "cyan")
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

	// Create database files
	if appDatabase != "none" {
		SendMessage("\n[PROCESS] Database", "cyan")
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

	// END message
	SendMessage("\n[DONE] Run `make` from `"+appPath+"` folder!", "yellow")

	return nil
}
