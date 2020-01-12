package cgapp

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/schollz/progressbar/v2"
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

	// Configs
	dotfiles string = "configs/dotfiles"
	docker   string = "configs/docker"
)

var (
	// App options:
	appBackend  string
	appFrontend string
	appPath     string
	appStack    string
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
				Name:  "init",
				Usage: "init new app",
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
				},
				Action: func(c *cli.Context) error {
					// Start message
					fmt.Printf("\n%v> Start creating new app...%v\n", yellow, noColor)

					// Create app folder
					ErrChecker(os.Mkdir(appPath, 0755))
					fmt.Printf("\n%v[OK]%v Main app folder was created!\n", green, noColor)

					// Copy configs files
					fmt.Printf("\n%v> Copy app config files%v\n\n", cyan, noColor)
					ErrChecker(CopyFolder(dotfiles))
					fmt.Printf("\n%v[OK]%v Config files was copied!\n", green, noColor)

					// Create backend files
					fmt.Printf("\n%v> Creating app backend%v\n\n", cyan, noColor)
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
						fmt.Printf("\n%v> Creating app frontend%v\n\n", cyan, noColor)
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

						// Install dependencies for frontend
						fmt.Printf(
							"\n%v> Installing frontend dependencies%v (may take some time!)\n\n",
							cyan, noColor,
						)

						// Create progress bar with 0%
						bar := progressbar.NewOptions(
							100,
							progressbar.OptionSetRenderBlankState(true),
						)

						// Go to ./frontend folder and run npm install
						cmd := exec.Command("npm", "install")
						cmd.Dir = appPath + string(os.PathSeparator) + "frontend"
						ErrChecker(cmd.Run())

						// Run progress bar from 0% to 100%
						for i := 0; i < 100; i++ {
							bar.Add(1)
							time.Sleep(10 * time.Millisecond)
						}

						// Show success report
						fmt.Printf(
							"\n\n%v[OK]%v Frontend dependencies was installed!\n",
							green, noColor,
						)
					}

					// End message
					fmt.Printf(
						"\n%v[DONE] Run `make` from '%v' folder...%v\n\n",
						green, appPath, noColor,
					)

					// Default return
					return nil
				},
			},
			{
				Name:  "docker",
				Usage: "create configured Docker containers",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "path",
						Aliases:     []string{"p"},
						Value:       ".",
						Usage:       "path to create containers, ex. ~/projects/my-app",
						Required:    false,
						Destination: &appPath,
					},
				},
				Subcommands: []*cli.Command{
					{
						Name:     "nginx",
						Usage:    "container with Nginx (alpine:latest) and Certbot",
						Category: "Configured Docker containers",
						Action: func(c *cli.Context) error {
							// Start message
							fmt.Printf(
								"\n%v> Start configuring Docker containers...%v\n",
								yellow, noColor,
							)

							// Create app folder, if not exist
							if _, err := os.Stat(appPath); !os.IsNotExist(err) {
								os.Mkdir(appPath, 0755)
							}

							// Check frontend folder
							_, err := os.Stat(appPath + string(os.PathSeparator) + "frontend")
							if !os.IsNotExist(err) {
								// If exists, copy fullstack app docker-compose file
								appStack = docker + string(os.PathSeparator) + "nginx-fullstack"
							} else {
								// Else, copy only backend docker-compose file
								appStack = docker + string(os.PathSeparator) + "nginx-backend-only"
							}

							// Copy docker-compose configs files
							fmt.Printf("\n%v> Copy docker-compose.yml file%v\n\n", cyan, noColor)
							ErrChecker(CopyFile(appStack, "docker-compose.yml"))
							fmt.Printf(
								"\n%v[OK]%v docker-compose.yml file was copied!\n",
								green, noColor,
							)

							// Create container files
							fmt.Printf(
								"\n%v> Creating container with Nginx and Certbot%v\n\n",
								cyan, noColor,
							)
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

							// End message
							fmt.Printf(
								"\n%v[DONE] Run `docker-compose up` from '%v' folder...%v\n\n",
								green, appPath, noColor,
							)

							// Default return
							return nil
						},
					},
				},
			},
		},
	}

	// Run new CLI
	ErrChecker(cgapp.Run(os.Args))
}
