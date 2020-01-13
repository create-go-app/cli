package cgapp

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

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
				Name:  "start",
				Usage: "start new app",
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
						Value:       "none",
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
					// Start message
					fmt.Printf("\n%v[START] Creating a new app...%v\n", green, noColor)

					// Create app folder and config files
					fmt.Printf("\n%v> App folder and config files%v\n", cyan, noColor)
					ErrChecker(os.Mkdir(appPath, 0755))
					fmt.Printf("\n%v[OK]%v App folder was created!\n", green, noColor)
					ErrChecker(File("Makefile", ""))

					// Create backend files
					fmt.Printf("\n%v> App backend%v\n\n", cyan, noColor)
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

					// If need to create frontend too
					if appFrontend != "none" {
						// Create frontend files
						fmt.Printf("\n%v> App frontend%v\n\n", cyan, noColor)
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

						// Install frontend dependencies frontend
						fmt.Printf("\n%v> Frontend dependencies%v\n", cyan, noColor)

						// Go to ./frontend folder and run npm install
						fmt.Printf(
							"\n%v[WAIT]%v Installing frontend dependencies (may take some time)!\n",
							yellow, noColor,
						)
						cmd := exec.Command("npm", "install")
						cmd.Dir = filepath.Join(appPath, "frontend")
						ErrChecker(cmd.Run())

						// Show success report
						fmt.Printf(
							"%v[OK]%v Frontend dependencies was installed!\n",
							green, noColor,
						)
					}

					// Create Docker containers
					if appWebServer != "none" || appDatabase != "none" {
						// Start message
						fmt.Printf(
							"\n%v[START] Configuring Docker containers...%v\n",
							green, noColor,
						)

						// If need to create web/proxy server too
						if appWebServer != "none" {
							// Create container files
							fmt.Printf("\n%v> Web/proxy server%v\n\n", cyan, noColor)
							ErrChecker(
								Create(&Config{
									name:   "nginx",
									match:  "^(nginx)$",
									view:   "nginx",
									folder: appPath,
								},
									registry,
								),
							)
						}

						// If need to create database too
						if appDatabase != "none" {
							// Create database files
							fmt.Printf("\n%v> Database%v\n\n", cyan, noColor)
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

						// Create docker-compose configs file
						fmt.Printf("\n%v> File docker-compose.yml%v\n\n", cyan, noColor)

						// Check ./frontend folder
						_, err := os.Stat(filepath.Join(appPath, "frontend"))
						if !os.IsNotExist(err) {
							// If exists, create fullstack app docker-compose file
							ErrChecker(File("docker-compose.yml", ""))
						} else {
							// Else, create only backend docker-compose file
							ErrChecker(File("docker-compose.yml", ""))
						}
					}

					// End message
					fmt.Printf(
						"\n%v[DONE] Run %vdocker-compose up --build%v %vfrom '%v' folder!%v\n\n",
						green, yellow, noColor, green, appPath, noColor,
					)

					// Default return
					return nil
				},
			},
		},
	}

	// Run new CLI
	ErrChecker(cgapp.Run(os.Args))
}
