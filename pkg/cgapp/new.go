package cgapp

import (
	"github.com/urfave/cli/v2"
)

var (
	// For `create` command
	appPath      string
	appBackend   string
	appFrontend  string
	appWebServer string
	appDatabase  string

	// For `deploy` command
	deployPlaybook      string
	deployUsername      string
	deployHost          string
	deployDockerNetwork string
	deployAskBecomePass bool
)

// New function for starting a new CLI
func New() (*cli.App, error) {
	// Init CLI
	app := &cli.App{}

	// Configure CLI
	app.Name = "cgapp"
	app.Usage = "create and deploy a new Go (Golang) app by running one command."
	app.Version = version
	app.EnableBashCompletion = true

	// Setting CLI commands
	app.Commands = []*cli.Command{
		{
			Name:  "create",
			Usage: "create a new project with the selected configuration",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:        "path",
					Aliases:     []string{"p"},
					Value:       ".",
					Usage:       "path to create project, ex. ~/projects/my-app",
					Required:    false,
					Destination: &appPath,
				},
				&cli.StringFlag{
					Name:        "backend",
					Aliases:     []string{"b"},
					Value:       "net/http",
					Usage:       "backend for your project, ex. Fiber, Echo",
					Required:    false,
					Destination: &appBackend,
				},
				&cli.StringFlag{
					Name:        "frontend",
					Aliases:     []string{"f"},
					Value:       "none",
					Usage:       "frontend for your project, ex. Preact, React.js, React.ts",
					Required:    false,
					Destination: &appFrontend,
				},
				&cli.StringFlag{
					Name:        "webserver",
					Aliases:     []string{"w"},
					Value:       "none",
					Usage:       "web/proxy server for your project, ex. Nginx",
					Required:    false,
					Destination: &appWebServer,
				},
				&cli.StringFlag{
					Name:        "database",
					Aliases:     []string{"d"},
					Value:       "none",
					Usage:       "database for your project, ex. Postgres",
					Required:    false,
					Destination: &appDatabase,
				},
			},
			Action: CreateCLIAction,
		},
		{
			Name:  "deploy",
			Usage: "deploy Docker containers with your project to a remote server or run on your local machine",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:        "playbook",
					Aliases:     []string{"p"},
					Value:       "deploy-playbook.yml",
					Usage:       "name of Ansible playbook, ex. my-play.yml",
					Destination: &deployPlaybook,
				},
				&cli.StringFlag{
					Name:        "username",
					Aliases:     []string{"u"},
					Usage:       "username of remote's server user or your local machine, ex. root",
					Required:    true,
					Destination: &deployUsername,
				},
				&cli.StringFlag{
					Name:        "host",
					Aliases:     []string{"s"},
					Value:       "localhost",
					Usage:       "host name of remote's server or local machine (from Ansible inventory), ex. do_server_1",
					Destination: &deployHost,
				},
				&cli.StringFlag{
					Name:        "network",
					Aliases:     []string{"n"},
					Value:       "cgapp_network",
					Usage:       "network for Docker containers, ex. my_net",
					Destination: &deployDockerNetwork,
				},
				&cli.BoolFlag{
					Name:        "ask-become-pass",
					Usage:       "asking you to enter become user's password at start",
					Destination: &deployAskBecomePass,
				},
			},
			Action: DeployCLIAction,
		},
	}

	return app, nil
}
