package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"time"

	"github.com/markbates/pkger"
	"github.com/schollz/progressbar/v2"
	"github.com/urfave/cli/v2"
	"gopkg.in/src-d/go-git.v4"
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

	// Templates registry
	registry = map[string]string{
		"net/http": "create-go-app/net_http-go-template",
		"echo":     "create-go-app/echo-go-template",
		// "gin":      "create-go-app/gin-go-template",
		// "iris":     "create-go-app/iris-go-template",
		"react":  "create-go-app/react-js-template",
		"preact": "create-go-app/preact-js-template",
		// "vue":      "create-go-app/vue-js-template",
		// "svelte":   "create-go-app/svelte-js-template",
	}
)

// embedConfig struct for embed configuration
type embedConfig struct {
	embedFolder string
	appFolder   string
}

// appConfig struct for app configuration
type appConfig struct {
	name   string
	match  string
	view   string
	folder string
}

func main() {
	// Embed ./configs folder
	pkger.Include("/configs")

	// Configure CLI app
	cgapp := &cli.App{
		Name:    "cgapp",
		Usage:   "set up a new Go (Golang) full stack app by running one command.",
		Version: "0.1.2",
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
			errChecker(os.Mkdir(appPath, 0755))
			fmt.Printf("\n%v[✔️]%v App main folder was created!\n", green, noColor)

			// Create configs files
			fmt.Printf("\n%v▼ Creating app config%v\n\n", cyan, noColor)
			errChecker(createConfig(&embedConfig{
				embedFolder: configFolder,
				appFolder:   appPath,
			}))
			fmt.Printf("\n%v[✔️]%v App config was created!\n", green, noColor)

			// Create backend files
			fmt.Printf("\n%v▼ Creating app backend%v\n\n", cyan, noColor)
			errChecker(createApp(&appConfig{
				name:   strings.ToLower(appBackend),
				match:  "^(net/http|echo|gin|iris)$",
				view:   "backend",
				folder: appPath,
			}))

			// If need to create frontend too
			if appFrontend != "none" {
				// Create frontend files
				fmt.Printf("\n%v▼ Creating app frontend%v\n\n", cyan, noColor)
				errChecker(createApp(&appConfig{
					name:   strings.ToLower(appFrontend),
					match:  "^(p?react|vue|svelte|angular)$",
					view:   "frontend",
					folder: appPath,
				}))

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
				errChecker(cmd.Run())

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
				"\n%v👌 Done! Run `make` from folder '%v'...%v\n\n",
				green, appPath, noColor,
			)

			// Default return
			return nil
		},
	}

	// Run CLI app
	errChecker(cgapp.Run(os.Args))
}

// createConfig function for create app's config files from embed folder
func createConfig(e *embedConfig) error {
	return pkger.Walk(e.embedFolder, func(path string, info os.FileInfo, err error) error {
		// Define files paths
		folder := e.appFolder + string(os.PathSeparator) + info.Name()

		// Create files
		if !info.IsDir() {
			// Open file from embed binary
			from, err := pkger.Open(path)
			errChecker(err)
			defer from.Close()

			// Create file
			to, err := os.Create(folder)
			errChecker(err)
			defer to.Close()

			// Copy data from embed binary to real file
			_, err = io.Copy(to, from)
			errChecker(err)

			// Show report for each file
			fmt.Printf("— File '%v' was created!\n", folder)
		}

		// Default return
		return nil
	})
}

// createApp function for create app
func createApp(c *appConfig) error {
	// Create path to backend|frontend folder
	folder := c.folder + string(os.PathSeparator) + c.view

	// Create match expration for backend frameworks
	match, _ := regexp.MatchString(c.match, c.name)
	if match {
		// If match, create from default template
		_, err := git.PlainClone(folder, false, &git.CloneOptions{
			URL:      "https://github.com/" + registry[c.name],
			Progress: os.Stdout,
		})
		errChecker(err)

		// Clean
		os.RemoveAll(folder + string(os.PathSeparator) + ".git")

		// Show success report
		fmt.Printf(
			"\n%v[✔️]%v %v (%v) was created with default template '%v'!\n",
			green, noColor,
			strings.Title(c.view),
			strings.Title(c.name),
			registry[c.name],
		)
	} else {
		// Else create from user template (from GitHub, etc)
		_, err := git.PlainClone(folder, false, &git.CloneOptions{
			URL:      "https://" + c.name,
			Progress: os.Stdout,
		})
		errChecker(err)

		// Clean
		os.RemoveAll(folder + string(os.PathSeparator) + ".git")

		// Show success report
		fmt.Printf(
			"\n%v[✔️]%v %v was created with user template '%v'!\n",
			green, noColor,
			strings.Title(c.view),
			c.name,
		)
	}

	// Default return
	return nil
}

// errChecker function for check error
func errChecker(err error) {
	if err != nil {
		// Show error report
		fmt.Printf("\n%v[✘] Error: %v%v\n\n", red, err, noColor)
		os.Exit(1)
	}
}
