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

	"github.com/AlecAivazis/survey/v2"
	"github.com/create-go-app/cli/pkg/registry"
	"github.com/create-go-app/cli/pkg/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	useConfigFile                          bool                   // indicate using config
	projectConfig                          map[string]interface{} // parse project config
	rolesConfig, roles                     []string               // parse Ansible roles config
	backend, frontend, webserver, database string                 // define project variables
	answers                                registry.Answers       // define answers variable
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"new"},
	Short:   "Create a new project via interactive UI or configuration file",
	Long:    "\nCreate a new project via interactive UI or configuration file.",
	Run: func(cmd *cobra.Command, args []string) {
		// Start message.
		utils.SendMsg(true, "* * *", "Create a new project via Create Go App CLI v"+registry.CLIVersion+"...", "yellow", true)

		// If config is set and correct, skip survey and use it.
		if useConfigFile && projectConfig != nil {
			// Re-define variables from config file ($PWD/.cgapp.yml).
			backend = strings.ToLower(projectConfig["backend"].(string))
			frontend = strings.ToLower(projectConfig["frontend"].(string))
			webserver = strings.ToLower(projectConfig["webserver"].(string))
			database = strings.ToLower(projectConfig["database"].(string))
			roles = rolesConfig
		} else {
			// Start survey.
			if err := survey.Ask(
				registry.Questions,
				&answers,
				// See: https://github.com/mgutz/ansi#style-format
				survey.WithIcons(func(icons *survey.IconSet) {
					icons.Question.Format = "cyan"
					icons.Question.Text = "[?]"
					icons.Help.Format = "blue"
					icons.Help.Text = "Help ->"
					icons.Error.Format = "yellow"
					icons.Error.Text = "Warning ->"
				}),
			); err != nil {
				utils.SendMsg(true, "[ERROR]", err.Error(), "red", true)
				os.Exit(1)
			}

			// If something went wrong, cancel and exit.
			if !answers.Agree {
				utils.SendMsg(true, "[!]", "You're stopped creation of a new project.", "red", false)
				utils.SendMsg(false, "[!]", "Run `cgapp create` once again!", "red", true)
				os.Exit(1)
			}

			// Define variables for better display.
			backend = strings.ToLower(answers.Backend)
			frontend = strings.ToLower(answers.Frontend)
			webserver = strings.ToLower(answers.Webserver)
			database = strings.ToLower(answers.Database)
			roles = answers.Roles
		}

		// End message.
		utils.SendMsg(true, "(i)", "A helpful documentation and next steps -> https://create-go.app/", "green", false)
		utils.SendMsg(false, "(i)", "Run `cgapp deploy` to deploy your project to a remote server.", "green", true)
	},
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().BoolVarP(
		&useConfigFile,
		"use-config", "", false,
		"use config file to create a new project (default is $PWD/.cgapp.yml)",
	)
}

// initConfig reads in config file, if set.
func initConfig() {
	if useConfigFile {
		// Get current directory.
		currentDir, _ := os.Getwd()

		viper.AddConfigPath(currentDir) // add config path
		viper.SetConfigName(".cgapp")   // set config name

		// If a config file is found, read it in.
		if err := viper.ReadInConfig(); err != nil {
			utils.SendMsg(true, "[ERROR]", err.Error(), "red", true)
			os.Exit(1)
		}

		// Parse configs
		_ = viper.UnmarshalKey("project", &projectConfig)
		_ = viper.UnmarshalKey("roles", &rolesConfig)
	}
}
