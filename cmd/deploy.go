// Copyright 2019-present Vic Shóstak. All rights reserved.
// Use of this source code is governed by Apache 2.0 license
// that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"log"
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
	_ = cgapp.ShowMessage("warning", "Deploying project via Create Go App CLI v"+registry.CLIVersion+"...", true, true)

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
	_ = cgapp.ShowMessage("warning", "Run Ansible playbook for deploy your project...", true, true)

	// Run execution for Ansible playbook.
	if err := cgapp.ExecCommand("ansible-playbook", options); err != nil {
		log.Fatal(cgapp.ShowMessage("error", err.Error(), true, true))
	}

	// Stop timer.
	stopTimer := fmt.Sprintf("%.0f", time.Since(startTimer).Seconds())

	// End message.
	_ = cgapp.ShowMessage("success", "Completed in "+stopTimer+" seconds!", true, true)
	_ = cgapp.ShowMessage("", "A helpful documentation and next steps -> https://create-go.app/", false, true)
}

func init() {
	rootCmd.AddCommand(deployCmd)
	deployCmd.PersistentFlags().BoolVarP(
		&askBecomePass,
		"ask-become-pass", "K", false,
		"prompt you to provide the remote user sudo password (standard Ansible `--ask-become-pass` option)",
	)
}
