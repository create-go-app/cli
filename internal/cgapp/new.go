package cgapp

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	box "github.com/create-go-app/cli/configs"
	"github.com/urfave/cli/v2"
)

const (
	// Colors:
	red    string = "\033[0;31m"
	green  string = "\033[0;32m"
	cyan   string = "\033[0;36m"
	yellow string = "\033[1;33m"

	// Clear color
	noColor string = "\033[0m"
)

var (
	// App options:
	appPath      string
	appBackend   string
	appFrontend  string
	appWebServer string
	appDatabase  string
)

// New function for start new CLI
func New(version string, registry map[string]string) {
	// Configure CLI app
	cgapp := &cli.App{
		Name:                 "cgapp",
		Usage:                "set up a new Go (Golang) full stack app by running one command.",
		Version:              version,
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			{
				Name:  "create",
				Usage: "create new Go app",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "path",
						Aliases:     []string{"p"},
						Value:       ".",
						Usage:       "path to create app, ex. ~/projects/my-app",
						Required:    false,
						Destination: &appPath,
					},
					&cli.StringFlag{
						Name:        "backend",
						Aliases:     []string{"b"},
						Value:       "net/http",
						Usage:       "backend for your app, ex. Echo, Gin, Iris, net/http",
						Required:    false,
						Destination: &appBackend,
					},
					&cli.StringFlag{
						Name:        "frontend",
						Aliases:     []string{"f"},
						Value:       "none",
						Usage:       "frontend for your app, ex. (P)React, Vue, Svelte",
						Required:    false,
						Destination: &appFrontend,
					},
					&cli.StringFlag{
						Name:        "webserver",
						Aliases:     []string{"w"},
						Value:       "nginx",
						Usage:       "web/proxy server for your app, ex. Nginx",
						Required:    false,
						Destination: &appWebServer,
					},
					&cli.StringFlag{
						Name:        "database",
						Aliases:     []string{"d"},
						Value:       "none",
						Usage:       "database for your app, ex. Postgres",
						Required:    false,
						Destination: &appDatabase,
					},
				},
				Action: func(c *cli.Context) error {
					/*
					 *	START message
					 */

					fmt.Printf("\n%v[START] Creating a new app...%v\n", green, noColor)

					/*
					 *	FOLDER for app
					 */

					fmt.Printf("\n%v> App folder and config files%v\n", cyan, noColor)

					// Create main folder for app
					ErrChecker(os.Mkdir(appPath, 0750))
					fmt.Printf("\n%v[OK]%v App folder was created!\n", green, noColor)

					/*
					 *	CONFIG files
					 */

					// Create config files for app
					ErrChecker(File(".editorconfig", box.Get("/dotfiles/.editorconfig")))
					ErrChecker(File(".gitignore", box.Get("/dotfiles/.gitignore")))
					ErrChecker(File("Makefile", box.Get("/dotfiles/Makefile")))

					/*
					 *	BACKEND files
					 */

					fmt.Printf("\n%v> App backend%v\n\n", cyan, noColor)

					// Create backend files
					ErrChecker(
						Create(&Config{
							name:   strings.ToLower(appBackend),
							match:  "^(net/http|echo|gin|iris)$",
							view:   "backend",
							folder: appPath,
						},
							registry,
						),
					)

					/*
					 *	FRONTEND files
					 */

					if appFrontend != "none" {
						fmt.Printf("\n%v> App frontend%v\n\n", cyan, noColor)

						// Create frontend files
						ErrChecker(
							Create(&Config{
								name:   strings.ToLower(appFrontend),
								match:  "^(p?react|vue|svelte)$",
								view:   "frontend",
								folder: appPath,
							},
								registry,
							),
						)

						/*
						 *	FRONTEND dependencies
						 */

						fmt.Printf("\n%v> Frontend dependencies%v\n", cyan, noColor)

						fmt.Printf(
							"\n%v[WAIT]%v Installing frontend dependencies (may take some time)!\n",
							yellow, noColor,
						)

						// Go to ./frontend folder and run npm install
						cmd := exec.Command("npm", "install")
						cmd.Dir = filepath.Join(appPath, "frontend")
						ErrChecker(cmd.Run())

						fmt.Printf(
							"%v[OK]%v Frontend dependencies was installed!\n",
							green, noColor,
						)
					}

					/*
					 *	DOCKER containers
					 */

					fmt.Printf(
						"\n%v[START] Configuring Docker containers...%v\n",
						green, noColor,
					)

					/*
					 *	WEB/PROXY SERVER container
					 */

					fmt.Printf("\n%v> Web/proxy server%v\n\n", cyan, noColor)

					// Create container files
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

					/*
					 *	DATABASE container
					 */

					if appDatabase != "none" {
						fmt.Printf("\n%v> Database%v\n\n", cyan, noColor)

						// Create database files
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

					/*
					 *	DOCKER-COMPOSE file
					 */

					fmt.Printf("\n%v> File docker-compose.yml%v\n\n", cyan, noColor)

					// Check ./frontend folder
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

					/*
					 *	END message
					 */

					fmt.Printf(
						"\n%v[DONE] Run %vmake%v %vfrom '%v' folder!%v\n\n",
						green, yellow, noColor, green, appPath, noColor,
					)

					return nil
				},
			},
		},
	}

	// Run new CLI
	ErrChecker(cgapp.Run(os.Args))
}
