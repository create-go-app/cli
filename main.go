package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"

	"github.com/markbates/pkger"
	"github.com/urfave/cli/v2"
	"gopkg.in/src-d/go-git.v4"
)

const (
	noColor = "\033[0m"
	red     = "\033[0;31m"
	green   = "\033[0;32m"
	cyan    = "\033[0;36m"
	yellow  = "\033[1;33m"
)

var (
	appName     string
	appBackend  string
	appFrontend string
	appPath     string
	registry    = map[string]string{
		"echo":   "create-go-app/echo-go-template",
		"preact": "create-go-app/preact-js-template",
	}
)

func main() {
	// Embed ./configs folder
	pkger.Include("/configs")
	// Configure CLI app
	cgapp := &cli.App{
		Name:    "cgapp",
		Usage:   "set up a new Go (Golang) full stack app by running one command.",
		Version: "0.1.0",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "name",
				Aliases:     []string{"n"},
				Usage:       "name of your go module (ex. github.com/user/my-app)",
				Required:    true,
				Destination: &appName,
			},
			&cli.StringFlag{
				Name:        "path",
				Aliases:     []string{"p"},
				Value:       ".",
				Usage:       "path to create app (ex. ~/projects/my-app)",
				Required:    false,
				Destination: &appPath,
			},
			&cli.StringFlag{
				Name:        "backend",
				Aliases:     []string{"b"},
				Usage:       "backend for your app (support: Echo)",
				Required:    true,
				Destination: &appBackend,
			},
			&cli.StringFlag{
				Name:        "frontend",
				Aliases:     []string{"f"},
				Value:       "none",
				Usage:       "frontend for your app (support: Preact)",
				Required:    false,
				Destination: &appFrontend,
			},
		},
		Action: func(c *cli.Context) error {
			// Start message
			fmt.Printf("\n%v▶️ Start creating new app '%v'...%v\n", yellow, appName, noColor)
			// Create configs files
			fmt.Printf("\n%v▼ Creating app config%v\n\n", cyan, noColor)
			errChecker(createConfig())
			fmt.Printf("\n%v[✔️]%v App config was created!\n", green, noColor)
			// Create backend files
			fmt.Printf("\n%v▼ Creating app backend%v\n\n", cyan, noColor)
			errChecker(createBackend())
			// If need to create frontend too
			if appFrontend != "none" {
				// Create frontend files
				fmt.Printf("\n%v▼ Creating app frontend%v\n\n", cyan, noColor)
				errChecker(createFrontend())
			}
			// Show report
			fmt.Printf("\n%v👌 Done! Run `docker-compose up`%v\n\n", green, noColor)
			// Default return
			return nil
		},
	}
	// Run CLI app
	errChecker(cgapp.Run(os.Args))
}

// createConfig: function for create app's config files
func createConfig() error {
	return pkger.Walk("/configs", func(path string, info os.FileInfo, err error) error {
		// Define files paths
		folder := appPath + string(os.PathSeparator) + info.Name()
		// Create files (exclude .DS_Store)
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
			// Show report
			fmt.Printf("— File '%v' was created!\n", folder)
		}
		return nil
	})
}

// createBackend: function for create app's backend
func createBackend() error {
	// Convert all args to lowercase
	appBackend = strings.ToLower(appBackend)
	//
	folder := appPath + string(os.PathSeparator) + "backend"
	// Create match expration for backend frameworks
	match, _ := regexp.MatchString("^(nethttp|echo|gin|iris)$", appBackend)
	if match {
		// If match, create from default template
		_, err := git.PlainClone(folder, false, &git.CloneOptions{
			URL:      "https://github.com/" + registry[appBackend],
			Progress: os.Stdout,
		})
		errChecker(err)
		// Show report
		fmt.Printf(
			"\n%v[✔️]%v Backend (%v) was created with default template '%v'!\n",
			green, noColor, strings.Title(appBackend), registry[appBackend],
		)
	} else {
		// Else create from user template (from GitHub, etc)
		_, err := git.PlainClone(folder, false, &git.CloneOptions{
			URL:      appBackend,
			Progress: os.Stdout,
		})
		errChecker(err)
		// Show report
		fmt.Printf(
			"\n%v[✔️]%v Backend was created with user template '%v'!\n",
			green, noColor, appBackend,
		)
	}
	return nil
}

// createFrontend: function for create app's frontend
func createFrontend() error {
	// Convert arg to lowercase
	appFrontend = strings.ToLower(appFrontend)
	//
	folder := appPath + string(os.PathSeparator) + "frontend"
	// Create match expration for front frameworks
	match, _ := regexp.MatchString("^(p?react|vue|svelte|angular)$", appFrontend)
	if match {
		// If match, create from default template
		_, err := git.PlainClone(folder, false, &git.CloneOptions{
			URL:      "https://github.com/" + registry[appFrontend],
			Progress: os.Stdout,
		})
		errChecker(err)
		// Show report
		fmt.Printf(
			"\n%v[✔️]%v Frontend (%v) was created with default template '%v'!\n",
			green, noColor, strings.Title(appFrontend), registry[appFrontend],
		)
	} else {
		// Else create from user template (from GitHub, etc)
		_, err := git.PlainClone(folder, false, &git.CloneOptions{
			URL:      appFrontend,
			Progress: os.Stdout,
		})
		errChecker(err)
		// Show report
		fmt.Printf(
			"\n%v[✔️]%v Frontend was created with user template '%v'!\n",
			green, noColor, appFrontend,
		)
	}
	return nil
}

// errChecker: function for check error
func errChecker(err error) {
	if err != nil {
		// Show report
		fmt.Printf("\n%v[✘] Error: %v%v\n\n", red, err, noColor)
		os.Exit(1)
	}
}
