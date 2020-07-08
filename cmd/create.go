/*
Package cmd includes all of the Create Go App CLI commands.

Copyright © 2020 Vic Shóstak <truewebartisans@gmail.com> (https://1wa.co)

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
	"github.com/create-go-app/cli/pkg/embed"
	"github.com/create-go-app/cli/pkg/registry"
	"github.com/create-go-app/cli/pkg/utils"
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
	utils.SendMsg(true, "* * *", "Create a new project via Create Go App CLI v"+registry.CLIVersion+"...", "yellow", true)

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
			utils.SendMsg(true, "[ERROR]", err.Error(), "red", true)
			os.Exit(1)
		}

		// If something went wrong, cancel and exit.
		if !createAnswers.AgreeCreation {
			utils.SendMsg(true, "[!]", "You're stopped creation of a new project.", "red", false)
			utils.SendMsg(false, "[!]", "Run `cgapp create` once again!", "red", true)
			os.Exit(1)
		}

		// Insert empty line.
		utils.SendMsg(false, "", "", "", false)

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
	currentDir, _ := os.Getwd()

	// Create config files for your project.
	utils.SendMsg(false, "*", "Create config files for your project...", "cyan", true)

	// Create configuration files.
	filesToMake := map[string][]byte{
		".gitignore":     embed.Get("/.gitignore"),
		".gitattributes": embed.Get("/.gitattributes"),
		".editorconfig":  embed.Get("/.editorconfig"),
		"Taskfile.yml":   embed.Get("/Taskfile.yml"),
	}
	if err := utils.MakeFiles(currentDir, filesToMake); err != nil {
		utils.SendMsg(true, "[ERROR]", err.Error(), "red", true)
		os.Exit(1)
	}

	// Create Ansible playbook and download roles, if not skipped.
	if installAnsibleRoles {
		utils.SendMsg(true, "*", "Create Ansible playbook and roles...", "cyan", true)

		// Create playbook.
		fileToMake := map[string][]byte{
			"deploy-playbook.yml": embed.Get("/deploy-playbook.yml"),
		}
		if err := utils.MakeFiles(currentDir, fileToMake); err != nil {
			utils.SendMsg(true, "[ERROR]", err.Error(), "red", true)
			os.Exit(1)
		}

		// Create Ansible roles.
		// if err := CreateProjectFromRegistry(
		// 	&Project{Type: "roles", Name: "deploy", RootFolder: appPath},
		// 	registry.Repositories,
		// ); err != nil {
		// 	return ThrowError(err.Error())
		// }
	}

	// Create backend files.
	utils.SendMsg(true, "*", "Create project backend...", "cyan", false)
	// if err := CreateProjectFromRegistry(
	// 	&Project{Type: "backend", Name: strings.ToLower(appBackend), RootFolder: appPath},
	// 	registry,
	// ); err != nil {
	// 	utils.SendMsg(true, "[ERROR]", err.Error(), "red", true)
	// 	os.Exit(1)
	// }

	if frontend != "none" {
		// Create frontend files.
		utils.SendMsg(true, "*", "Create project frontend...", "cyan", false)
		// if err := CreateProjectFromCMD(
		// 	&Project{Type: "frontend", Name: answers.Frontend, RootFolder: currentDir},
		// 	cmds,
		// ); err != nil {
		// 	utils.SendMsg(true, "[ERROR]", err.Error(), "red", true)
		// 	os.Exit(1)
		// }
	}

	// Docker containers.
	if webserver != "none" || database != "none" {

		utils.SendMsg(true, "* * *", "Configuring Docker containers...", "yellow", false)

		if webserver != "none" {
			// Create container with a web/proxy server.
			utils.SendMsg(true, "*", "Create container with web/proxy server...", "cyan", false)
			// if err := CreateProjectFromRegistry(
			// 	&Project{Type: "webserver", Name: answers.Webserver, RootFolder: currentDir},
			// 	registry,
			// ); err != nil {
			// 	utils.SendMsg(true, "[ERROR]", err.Error(), "red", true)
			// 	os.Exit(1)
			// }
		}

		if database != "none" {
			// Create container with a database.
			utils.SendMsg(true, "*", "Create container with database...", "cyan", false)
			// if err := CreateProjectFromRegistry(
			// 	&Project{Type: "database", Name: answers.Database, RootFolder: currentDir},
			// 	registry,
			// ); err != nil {
			// 	utils.SendMsg(true, "[ERROR]", err.Error(), "red", true)
			// 	os.Exit(1)
			// }
		}
	}

	// Stop timer
	stopTimer := time.Since(startTimer).String()

	// End message.
	utils.SendMsg(true, "* * *", "Completed in "+stopTimer+"!", "yellow", true)
	utils.SendMsg(false, "(i)", "A helpful documentation and next steps -> https://create-go.app/", "green", false)
	utils.SendMsg(false, "(i)", "Run `cgapp deploy` to deploy your project to a remote server.", "green", true)
}

func init() {
	rootCmd.AddCommand(createCmd)
}
