// Copyright 2019-present Vic Shóstak. All rights reserved.
// Use of this source code is governed by Apache 2.0 license
// that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/create-go-app/cli/pkg/cgapp"
	"github.com/create-go-app/cli/pkg/registry"
	"github.com/spf13/cobra"
)

// deployCmd represents the `deploy` command.
var deployCmd = &cobra.Command{
	Use:     "deploy",
	Aliases: []string{"push"},
	Short:   "Deploy your project to the remote server via Ansible",
	Long:    "\nDeploy your project to the remote server by Ansible playbooks and roles.",
	Run:     runDeployCmd,
}

// runDeployCmd represents runner for the `deploy` command.
var runDeployCmd = func(cmd *cobra.Command, args []string) {
	// Start message.
	cgapp.SendMsg(true, "* * *", "Deploying project via Create Go App CLI v"+registry.CLIVersion+"...", "yellow", true)

	// Start timer.
	startTimer := time.Now()

	// Define Ansible options.
	if askBecomePass {
		//
		options = []string{"playbook.yml", "-i", "hosts.ini", "-K"}
	} else {
		//
		options = []string{"playbook.yml", "-i", "hosts.ini"}
	}

	// Create config files for your project.
	cgapp.SendMsg(false, "*", "Run Ansible playbook for deploy your project...", "cyan", true)

	// Run execution for Ansible playbook.
	if err := cgapp.ExecCommand("ansible-playbook", options); err != nil {
		cgapp.SendMsg(true, "[ERROR]", err.Error(), "red", true)
		os.Exit(1)
	}

	// Stop timer.
	stopTimer := fmt.Sprintf("%.0f", time.Since(startTimer).Seconds())

	// End message.
	cgapp.SendMsg(true, "* * *", "Completed in "+stopTimer+" seconds!", "yellow", true)
	cgapp.SendMsg(false, "(i)", "A helpful documentation and next steps -> https://create-go.app/", "green", false)
	cgapp.SendMsg(false, "(i)", "Go to the project domain to see your deployed project! :)", "green", true)
}

func init() {
	rootCmd.AddCommand(deployCmd)
	deployCmd.PersistentFlags().BoolVarP(
		&askBecomePass,
		"ask-become-pass", "K", false,
		"prompt you to provide the remote user sudo password (standard Ansible `--ask-become-pass` option)",
	)
}
