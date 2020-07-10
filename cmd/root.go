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

	"github.com/AlecAivazis/survey/v2"
	"github.com/create-go-app/cli/pkg/cgapp"
	"github.com/create-go-app/cli/pkg/registry"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	useConfigFile                          bool                                           // indicate using config (from $PWD/.cgapp.yml)
	projectConfig                          map[string]interface{}                         // parse project config
	rolesConfig                            map[string]interface{}                         // parse Ansible roles config
	backend, frontend, webserver, database string                                         // define project variables
	installAnsibleRoles, askBecomePass     bool                                           // install Ansible roles, ask become pass
	username, host, network                string                                         // define deploy variables
	playbook                               string                 = "deploy-playbook.yml" // default Ansible playbook
	createAnswers                          registry.CreateAnswers                         // define answers variable for `create` command
	deployAnswers                          registry.DeployAnswers                         // define answers variable for `deploy` command

	// Config for survey icons and colors.
	// See: https://github.com/mgutz/ansi#style-format
	surveyIconsConfig = func(icons *survey.IconSet) {
		icons.Question.Format = "cyan"
		icons.Question.Text = "[?]"
		icons.Help.Format = "blue"
		icons.Help.Text = "Help ->"
		icons.Error.Format = "yellow"
		icons.Error.Text = "Note ->"
	}
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "cgapp",
	Version: registry.CLIVersion,
	Short:   "A powerful CLI for the Create Go App project",
	Long: `
A powerful CLI for the Create Go App project.

Create a new production-ready project with backend (Golang), 
frontend (JavaScript, TypeScript) and deploy automation 
(Ansible, Docker) by running one CLI command.

-> Focus on writing code and thinking of business logic!
<- The Create Go App CLI will take care of the rest.`,
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().BoolVarP(
		&useConfigFile,
		"use-config", "c", false,
		"use config file to create a new project or deploy to a remote server (default is $PWD/.cgapp.yml)",
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
			cgapp.SendMsg(true, "[ERROR]", err.Error(), "red", true)
			os.Exit(1)
		}

		// Parse configs
		_ = viper.UnmarshalKey("project", &projectConfig)
		_ = viper.UnmarshalKey("roles", &rolesConfig)
	}
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		cgapp.SendMsg(true, "[ERROR]", err.Error(), "red", true)
		os.Exit(1)
	}
}
