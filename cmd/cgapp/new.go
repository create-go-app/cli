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

	// Configs folder
	configFolder string = "/configs"
)

var (
	// App options:
	appBackend  string
	appFrontend string
	appPath     string
)

// New function for start new CLI
func New(version string, registry map[string]string) {
	// Configure CLI app
	cgapp := &cli.App{
		Name:    "cgapp",
		Usage:   "set up a new Go (Golang) full stack app by running one command.",
		Version: version,
		Flags: []cli.Flag{
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
				Name:        "path",
				Aliases:     []string{"p"},
				Value:       ".",
				Usage:       "path to create app, ex. ~/projects/my-app",
				Required:    false,
				Destination: &appPath,
			},
		},
		Action: func(c *cli.Context) error {
			// Start message
			fmt.Printf("\n%v▶️ Start creating new app...%v\n", yellow, noColor)

			// Create app folder
			ErrChecker(os.Mkdir(appPath, 0755))
			fmt.Printf("\n%v[✔️]%v App main folder was created!\n", green, noColor)

			// Create configs files
			fmt.Printf("\n%v▼ Creating app config%v\n\n", cyan, noColor)
			ErrChecker(CreateConfig(&embedConfig{
				embedFolder: configFolder,
				appFolder:   appPath,
			}))
			fmt.Printf("\n%v[✔️]%v App config was created!\n", green, noColor)

			// Create backend files
			fmt.Printf("\n%v▼ Creating app backend%v\n\n", cyan, noColor)
			ErrChecker(
				CreateApp(&appConfig{
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
				fmt.Printf("\n%v▼ Creating app frontend%v\n\n", cyan, noColor)
				ErrChecker(
					CreateApp(&appConfig{
						name:   strings.ToLower(appFrontend),
						match:  "^(p?react|vue|svelte|angular)$",
						view:   "frontend",
						folder: appPath,
					},
						registry,
					),
				)

				// Install dependencies for frontend
				fmt.Printf(
					"\n%v▼ Installing frontend dependencies%v (may take some time!)\n\n",
					cyan, noColor,
				)

				// Create progress bar with 0%
				bar := progressbar.NewOptions(100, progressbar.OptionSetRenderBlankState(true))

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
				fmt.Printf("\n\n%v[✔️]%v Frontend dependencies was installed!\n", green, noColor)
			}

			// End message
			fmt.Printf(
				"\n%v👌 Done! Run `make` from '%v' folder...%v\n\n",
				green, appPath, noColor,
			)

			// Default return
			return nil
		},
	}

	// Run new CLI
	ErrChecker(cgapp.Run(os.Args))
}
