/*
Package cmd includes all of the Create Go App CLI commands.

Create a new production-ready project with backend (Golang),
frontend (JavaScript, TypeScript) and deploy automation
(Ansible, Docker) by running one CLI command.

-> Focus on writing code and thinking of business logic!
<- The Create Go App CLI will take care of the rest.

Copyright © 2019-present Vic Shóstak <truewebartisans@gmail.com> (https://1wa.co)

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"os"
	"strings"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/create-go-app/cli/pkg/cgapp"
	"github.com/create-go-app/cli/pkg/embed"
	"github.com/create-go-app/cli/pkg/registry"
	"github.com/spf13/cobra"
)

// createCmd represents the `create` command.
var createCmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"new"},
	Short:   "Create a new project via interactive UI or configuration file",
	Long:    "\nCreate a new project via interactive UI or configuration file.",
	Run:     runCreateCmd,
}

// runCreateCmd represents runner for the `create` command.
var runCreateCmd = func(cmd *cobra.Command, args []string) {
	// Start message.
	cgapp.SendMsg(true, "* * *", "Create a new project via Create Go App CLI v"+registry.CLIVersion+"...", "yellow", true)

	// If config is set and correct, skip survey and use it.
	if useConfigFile && projectConfig != nil {
		// Re-define variables from config file (default is $PWD/.cgapp.yml).
		backend = strings.ToLower(projectConfig["backend"].(string))
		frontend = strings.ToLower(projectConfig["frontend"].(string))
		webserver = strings.ToLower(projectConfig["webserver"].(string))
		database = strings.ToLower(projectConfig["database"].(string))

		// Check, if config file contains `roles` section
		if rolesConfig != nil {
			installAnsibleRoles = true
		}
	} else {
		// Start survey.
		if err := survey.Ask(
			registry.CreateQuestions, &createAnswers, survey.WithIcons(surveyIconsConfig),
		); err != nil {
			cgapp.SendMsg(true, "[ERROR]", err.Error(), "red", true)
			os.Exit(1)
		}

		// If something went wrong, cancel and exit.
		if !createAnswers.AgreeCreation {
			cgapp.SendMsg(true, "[!]", "You're stopped creation of a new project.", "red", false)
			cgapp.SendMsg(false, "[!]", "Run `cgapp create` once again!", "red", true)
			os.Exit(1)
		}

		// Insert empty line.
		cgapp.SendMsg(false, "", "", "", false)

		// Define variables for better display.
		backend = strings.ToLower(createAnswers.Backend)
		frontend = strings.ToLower(createAnswers.Frontend)
		webserver = strings.ToLower(createAnswers.Webserver)
		database = strings.ToLower(createAnswers.Database)
		installAnsibleRoles = createAnswers.InstallAnsibleRoles
	}

	// Start timer.
	startTimer := time.Now()

	// Get current directory.
	currentDir, err := os.Getwd()
	if err != nil {
		cgapp.SendMsg(true, "[ERROR]", err.Error(), "red", true)
		os.Exit(1)
	}

	// Create config files for your project.
	cgapp.SendMsg(false, "*", "Create config files for your project...", "cyan", true)

	// Create configuration files.
	filesToMake := map[string][]byte{
		".gitignore":     embed.Get("/.gitignore"),
		".gitattributes": embed.Get("/.gitattributes"),
		".editorconfig":  embed.Get("/.editorconfig"),
		"Taskfile.yml":   embed.Get("/Taskfile.yml"),
	}
	if err := cgapp.MakeFiles(currentDir, filesToMake); err != nil {
		cgapp.SendMsg(true, "[ERROR]", err.Error(), "red", true)
		os.Exit(1)
	}

	// Create Ansible playbook and download roles, if not skipped.
	if installAnsibleRoles {
		cgapp.SendMsg(true, "*", "Create Ansible playbook and roles...", "cyan", true)

		// Create playbook.
		fileToMake := map[string][]byte{
			"deploy-playbook.yml": embed.Get("/deploy-playbook.yml"),
		}
		if err := cgapp.MakeFiles(currentDir, fileToMake); err != nil {
			cgapp.SendMsg(true, "[ERROR]", err.Error(), "red", true)
			os.Exit(1)
		}

		// Create Ansible roles.
		cgapp.CreateProjectFromRegistry(
			&registry.Project{
				Type:       "roles",
				Name:       "deploy",
				RootFolder: currentDir,
			},
			registry.Repositories,
		)
	}

	// Create backend files.
	cgapp.SendMsg(true, "*", "Create project backend...", "cyan", true)
	cgapp.CreateProjectFromRegistry(
		&registry.Project{
			Type:       "backend",
			Name:       backend,
			RootFolder: currentDir,
		},
		registry.Repositories,
	)

	if frontend != "none" {
		// Create frontend files.
		cgapp.SendMsg(true, "*", "Create project frontend...", "cyan", true)
		cgapp.CreateProjectFromCmd(
			&registry.Project{
				Type:       "frontend",
				Name:       frontend,
				RootFolder: currentDir,
			},
			registry.Commands,
		)
	}

	// Docker containers.
	if webserver != "none" || database != "none" {

		cgapp.SendMsg(true, "* * *", "Configuring Docker containers...", "yellow", false)

		if webserver != "none" {
			// Create container with a web/proxy server.
			cgapp.SendMsg(true, "*", "Create container with web/proxy server...", "cyan", true)
			cgapp.CreateProjectFromRegistry(
				&registry.Project{
					Type:       "webserver",
					Name:       webserver,
					RootFolder: currentDir,
				},
				registry.Repositories,
			)
		}

		if database != "none" {
			// Create container with a database.
			cgapp.SendMsg(true, "*", "Create container with database...", "cyan", true)
			cgapp.CreateProjectFromRegistry(
				&registry.Project{
					Type:       "database",
					Name:       database,
					RootFolder: currentDir,
				},
				registry.Repositories,
			)
		}
	}

	// Stop timer
	stopTimer := time.Since(startTimer).String()

	// End message.
	cgapp.SendMsg(true, "* * *", "Completed in "+stopTimer+"!", "yellow", true)
	cgapp.SendMsg(false, "(i)", "A helpful documentation and next steps -> https://create-go.app/", "green", false)
	cgapp.SendMsg(false, "(i)", "Run `cgapp deploy` to deploy your project to a remote server.", "green", true)
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.PersistentFlags().BoolVarP(
		&useConfigFile,
		"use-config", "c", false,
		"use config file to create a new project or deploy to a remote server (default is $PWD/.cgapp.yml)",
	)
}
