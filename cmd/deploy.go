// Copyright 2019-present Vic Shóstak. All rights reserved.
// Use of this source code is governed by Apache 2.0 license
// that can be found in the LICENSE file.

package cmd

import (
	"os"
	"strings"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/create-go-app/cli/pkg/cgapp"
	"github.com/create-go-app/cli/pkg/registry"
	"github.com/spf13/cobra"
)

// deployCmd represents the `deploy` command.
var deployCmd = &cobra.Command{
	Use:     "deploy",
	Aliases: []string{"serve"},
	Short:   "Deploy your project to a remote server via interactive UI or configuration file",
	Long:    "\nDeploy your project to a remote server via interactive UI or configuration file.",
	Run:     runDeployCmd,
}

// runDeployCmd represents runner for the `deploy` command.
var runDeployCmd = func(cmd *cobra.Command, args []string) {
	// Start message.
	cgapp.SendMsg(true, "* * *", "Deploying project via Create Go App CLI v"+registry.CLIVersion+"...", "yellow", true)

	// If config is set and correct, skip survey and use it.
	if useConfigFile && rolesConfig != nil {
		// Re-define variables from config file (default is $PWD/.cgapp.yml).
		username = strings.ToLower(rolesConfig["username"].(string))
		host = strings.ToLower(rolesConfig["host"].(string))
		network = strings.ToLower(rolesConfig["network"].(string))
		port = strings.ToLower(rolesConfig["port"].(string))
		askBecomePass = rolesConfig["become"].(bool)
	} else {
		// Start survey.
		if err := survey.Ask(
			registry.DeployQuestions, &deployAnswers, survey.WithIcons(surveyIconsConfig),
		); err != nil {
			cgapp.SendMsg(true, "[ERROR]", err.Error(), "red", true)
			os.Exit(1)
		}

		// If something went wrong, cancel and exit.
		if !deployAnswers.AgreeDeployment {
			cgapp.SendMsg(true, "[!]", "You're stopped deployment process of your project.", "red", false)
			cgapp.SendMsg(false, "[!]", "Run `cgapp deploy` once again!", "red", true)
			os.Exit(1)
		}

		// Insert empty line.
		cgapp.SendMsg(false, "", "", "", false)

		// Define variables for better display.
		username = deployAnswers.Username
		host = deployAnswers.Host
		network = deployAnswers.Network
		port = deployAnswers.BackendPort
		askBecomePass = deployAnswers.AskBecomePass
	}

	// Start timer.
	startTimer := time.Now()

	// Create config files for your project.
	cgapp.SendMsg(false, "*", "Run Ansible playbook `"+playbook+"`...", "cyan", true)

	// Define Ansible options.
	options := []string{
		playbook,
		"-u", username,
		"-e", "host=" + host + " network_name=" + network + " backend_port=" + port,
	}

	// Check, if need to ask password for username.
	// See: https://docs.ansible.com/ansible/latest/user_guide/become.html#become-command-line-options
	if askBecomePass {
		options = []string{
			playbook,
			"-u", username,
			"-e", "host=" + host + " network_name=" + network + " backend_port=" + port,
			"--ask-become-pass",
		}
	}

	// Run execution for Ansible playbook.
	if err := cgapp.ExecCommand("ansible-playbook", options); err != nil {
		cgapp.SendMsg(true, "[ERROR]", err.Error(), "red", true)
		os.Exit(1)
	}

	// Stop timer.
	stopTimer := time.Since(startTimer).String()

	// End message.
	cgapp.SendMsg(true, "* * *", "Completed in "+stopTimer+"!", "yellow", true)
	cgapp.SendMsg(false, "(i)", "A helpful documentation and next steps -> https://create-go.app/", "green", false)
	cgapp.SendMsg(false, "(i)", "Go to the `"+host+"` to see your deployed project! :)", "green", true)
}

func init() {
	rootCmd.AddCommand(deployCmd)
	deployCmd.PersistentFlags().BoolVarP(
		&useConfigFile,
		"use-config", "c", false,
		"use config file to create a new project or deploy to a remote server (default is $PWD/.cgapp.yml)",
	)
}
