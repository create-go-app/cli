package cgapp

import (
	"github.com/urfave/cli/v2"
)

// Options
var (
	appPath          string
	appBackend       string
	appFrontend      string
	appWebServer     string
	appDatabase      string
	appSilentRunning string
)

// New function for start new CLI
func New() (*cli.App, error) {
	// Init
	app := &cli.App{}

	// Configure
	app.Name = "cgapp"
	app.Usage = "set up a new Go (Golang) full stack app by running one command."
	app.Version = version
	app.EnableBashCompletion = true

	// CLI commands
	app.Commands = []*cli.Command{
		{
			Name:  "create",
			Usage: "create new Go app",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:        "path",
					Aliases:     []string{"p"},
					Value:       ".",
					Usage:       "path to create app, ex. ~/projects/my-app (default: \".\")",
					Required:    false,
					Destination: &appPath,
				},
				&cli.StringFlag{
					Name:        "backend",
					Aliases:     []string{"b"},
					Value:       "net/http",
					Usage:       "backend for your app, ex. Fiber, Echo (default: \"net/http\")",
					Required:    false,
					Destination: &appBackend,
				},
				&cli.StringFlag{
					Name:        "frontend",
					Aliases:     []string{"f"},
					Value:       "none",
					Usage:       "frontend for your app, ex. Preact, React.js, React.ts (default: \"none\")",
					Required:    false,
					Destination: &appFrontend,
				},
				&cli.StringFlag{
					Name:        "webserver",
					Aliases:     []string{"w"},
					Value:       "nginx",
					Usage:       "web/proxy server for your app (default: \"nginx\")",
					Required:    false,
					Destination: &appWebServer,
				},
				&cli.StringFlag{
					Name:        "database",
					Aliases:     []string{"d"},
					Value:       "none",
					Usage:       "database for your app, ex. Postgres (default: \"none\")",
					Required:    false,
					Destination: &appDatabase,
				},
				&cli.StringFlag{
					Name:        "silent",
					Value:       "none",
					Usage:       "silent running (default: \"none\")",
					Required:    false,
					Destination: &appSilentRunning,
				},
			},
			Action: CreateCLIAction,
		},
	}

	return app, nil
}
