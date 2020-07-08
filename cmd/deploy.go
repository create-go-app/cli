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
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/create-go-app/cli/pkg/registry"
	"github.com/create-go-app/cli/pkg/utils"
	"github.com/spf13/cobra"
)

// deployCmd represents the `deploy` command
var deployCmd = &cobra.Command{
	Use:     "deploy",
	Aliases: []string{"serve"},
	Short:   "Deploy your project to a remote server via interactive UI or configuration file",
	Long:    "\nDeploy your project to a remote server via interactive UI or configuration file.",
	Run:     runDeployCmd,
}

// runDeployCmd represents runner for the `deploy` command
var runDeployCmd = func(cmd *cobra.Command, args []string) {
	// Start message.
	utils.SendMsg(true, "* * *", "Deploying project via Create Go App CLI v"+registry.CLIVersion+"...", "yellow", true)

	// If config is set and correct, skip survey and use it.
	if useConfigFile && rolesConfig != nil {
		// Re-define variables from config file (default is $PWD/.cgapp.yml).
		username = strings.ToLower(rolesConfig["username"].(string))
		host = strings.ToLower(rolesConfig["host"].(string))
		network = strings.ToLower(rolesConfig["network"].(string))
	} else {
		// Start survey.
		if err := survey.Ask(
			registry.DeployQuestions, &deployAnswers, survey.WithIcons(surveyIconsConfig),
		); err != nil {
			utils.SendMsg(true, "[ERROR]", err.Error(), "red", true)
			os.Exit(1)
		}

		// If something went wrong, cancel and exit.
		if !deployAnswers.AgreeDeployment {
			utils.SendMsg(true, "[!]", "You're stopped deployment process of your project.", "red", false)
			utils.SendMsg(false, "[!]", "Run `cgapp deploy` once again!", "red", true)
			os.Exit(1)
		}

		// Insert empty line.
		utils.SendMsg(false, "", "", "", false)

		// Define variables for better display.
		username = deployAnswers.Username
		host = deployAnswers.Host
		network = deployAnswers.Network
	}

	// Start timer.
	startTimer := time.Now()

	// Create config files for your project.
	utils.SendMsg(false, "*", "Run Ansible playbook `"+playbook+"`...", "cyan", true)

	fmt.Println(username, host, network, playbook)

	// Stop timer
	stopTimer := time.Since(startTimer).String()

	// End message.
	utils.SendMsg(true, "* * *", "Completed in "+stopTimer+"!", "yellow", true)
	utils.SendMsg(false, "(i)", "A helpful documentation and next steps -> https://create-go.app/", "green", false)
	utils.SendMsg(false, "(i)", "Go to the `"+host+"` to see your deployed project! :)", "green", true)
}

func init() {
	rootCmd.AddCommand(deployCmd)
}
