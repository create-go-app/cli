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
	"fmt"
	"os"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/create-go-app/cli/pkg/registry"
	"github.com/create-go-app/cli/pkg/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	useConfigFile bool                   // indicate using config
	projectConfig map[string]interface{} // parse project config
	rolesConfig   []string               // parse Ansible roles config
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"new"},
	Short:   "Create a new project via interactive UI or configuration file",
	Long:    "\nCreate a new project via interactive UI or configuration file.",
	Run:     runCreateCommand,
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().BoolVarP(&useConfigFile, "use-config", "", false, "config file (default is $PWD/.cgapp.yml)")
}

// initConfig reads in config file, if set.
func initConfig() {
	if useConfigFile {
		// Get current directory
		currentDir, err := os.Getwd()
		if err != nil {
			utils.SendMsg(true, "[ERROR]", err.Error(), "red", true)
			os.Exit(1)
		}

		viper.AddConfigPath(currentDir) // add config path
		viper.SetConfigName(".cgapp")   // set config name

		// If a config file is found, read it in.
		if err = viper.ReadInConfig(); err != nil {
			utils.SendMsg(true, "[ERROR]", err.Error(), "red", true)
			os.Exit(1)
		}

		// Parse configs
		if err = viper.UnmarshalKey("project", &projectConfig); err != nil {
			utils.SendMsg(true, "[ERROR]", err.Error(), "red", true)
			os.Exit(1)
		}
		if err = viper.UnmarshalKey("roles", &rolesConfig); err != nil {
			utils.SendMsg(true, "[ERROR]", err.Error(), "red", true)
			os.Exit(1)
		}
	}
}

// runInitCommand ...
func runCreateCommand(cmd *cobra.Command, args []string) {
	// Init variables.
	var backend, frontend, webserver, database string
	var roles []string

	// Start message.
	utils.SendMsg(true, "* * *", "Create a new project via Create Go App CLI v"+registry.CLIVersion+"...", "yellow", true)

	// If config not set, go to survey.
	if !useConfigFile {

		// Define survey's questions.
		questions := []*survey.Question{
			{
				Name: "backend",
				Prompt: &survey.Select{
					Message: "Choose a backend framework:",
					Options: []string{"net/http", "Fiber", "Echo", "Gin"},
					Default: "Fiber",
				},
				Validate: survey.Required,
			},
			{
				Name: "frontend",
				Prompt: &survey.Select{
					Message: "Choose a frontend UI library:",
					Options: []string{"none", "React", "Preact", "Svelte"},
					Default: "none",
				},
			},
			{
				Name: "webserver",
				Prompt: &survey.Select{
					Message: "Choose a web/proxy server:",
					Options: []string{"none", "Nginx"},
					Default: "none",
				},
			},
			{
				Name: "database",
				Prompt: &survey.Select{
					Message: "Choose a database:",
					Options: []string{"none", "Postgres"},
					Default: "none",
				},
			},
			{
				Name: "roles",
				Prompt: &survey.MultiSelect{
					Message: "Choose an Ansible roles:",
					Options: []string{"deploy"},
					Default: []string{"deploy"},
					Help:    "These are fully configured Ansible roles, that will help you automate a deployment process.",
				},
			},
			{
				Name: "agree",
				Prompt: &survey.Confirm{
					Message: "If all is well, can I create this project?",
					Default: true,
				},
			},
		}

		// Define survey's answers.
		answers := struct {
			Backend   string
			Frontend  string
			Webserver string
			Database  string
			Roles     []string
			Agree     bool
		}{}

		// Start survey.
		if err := survey.Ask(
			questions,
			&answers,
			// See: https://github.com/mgutz/ansi#style-format
			survey.WithIcons(func(icons *survey.IconSet) {
				icons.Question.Format = "cyan"
				icons.Question.Text = "[?]"
				icons.Help.Format = "blue"
				icons.Help.Text = "Help ->"
			}),
		); err != nil {
			utils.SendMsg(true, "[ERROR]", err.Error(), "red", true)
			os.Exit(1)
		}

		// If something went wrong, cancel and exit.
		if !answers.Agree {
			utils.SendMsg(true, "[!]", "You're stopped a project creation.", "red", false)
			utils.SendMsg(false, "[!]", "Run `cgapp create` once again!", "red", true)
			os.Exit(1)
		}

		// Define variables for better display.
		backend = strings.ToLower(answers.Backend)
		frontend = strings.ToLower(answers.Frontend)
		webserver = strings.ToLower(answers.Webserver)
		database = strings.ToLower(answers.Database)
		roles = answers.Roles

	} else {

		// Exit, if config is invalid.
		if projectConfig == nil {
			utils.SendMsg(false, "[ERROR]", "Config file invalid or empty!", "red", true)
			os.Exit(1)
		}

		// Re-define variables, if using config file
		backend = strings.ToLower(projectConfig["backend"].(string))
		frontend = strings.ToLower(projectConfig["frontend"].(string))
		webserver = strings.ToLower(projectConfig["webserver"].(string))
		database = strings.ToLower(projectConfig["database"].(string))
		roles = rolesConfig

	}

	fmt.Println(backend, frontend, webserver, database, roles)

	// End message.
	utils.SendMsg(true, "(i)", "A helpful documentation and next steps -> https://create-go.app/", "green", false)
	utils.SendMsg(false, "(i)", "Run `cgapp deploy` to deploy your project to a remote server.", "green", true)
}
