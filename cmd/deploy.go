// Copyright 2022 Vic Shóstak and Create Go App Contributors. All rights reserved.
// Use of this source code is governed by Apache 2.0 license
// that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/create-go-app/cli/v3/pkg/cgapp"
	"github.com/create-go-app/cli/v3/pkg/registry"
)

func init() {
	rootCmd.AddCommand(deployCmd)
	deployCmd.Flags().BoolVarP(
		&askBecomePass,
		"ask-become-pass", "k", false,
		"prompt you to provide the remote user sudo password (standard Ansible `--ask-become-pass` option)",
	)
}

// deployCmd represents the `deploy` command.
var deployCmd = &cobra.Command{
	Use:     "deploy",
	Aliases: []string{"push"},
	Short:   "Deploy your project to the remote server via Ansible",
	Long:    "\nDeploy your project to the remote server by Ansible playbooks and roles.",
	RunE:    runDeployCmd,
}

// runDeployCmd represents runner for the `deploy` command.
func runDeployCmd(cmd *cobra.Command, args []string) error {
	// Start message.
	cgapp.ShowMessage(
		"",
		fmt.Sprintf("Deploying project via Create Go App CLI v%v...", registry.CLIVersion),
		true, true,
	)

	// Set Ansible playbook and inventory files.
	if askBecomePass {
		// With entering password.
		options = []string{"playbook.yml", "-i", "hosts.ini", "-K"}
	} else {
		// Without entering password.
		options = []string{"playbook.yml", "-i", "hosts.ini"}
	}

	// Create config files for your project.
	cgapp.ShowMessage(
		"info",
		"Ansible playbook for deploy your project is running. Please wait for completion!",
		false, false,
	)

	// Start timer.
	startTimer := time.Now()

	// Run execution for Ansible playbook.
	if err := cgapp.ExecCommand("ansible-playbook", options, false); err != nil {
		return cgapp.ShowError(err.Error())
	}

	// Stop timer.
	stopTimer := cgapp.CalculateDurationTime(startTimer)
	cgapp.ShowMessage(
		"info",
		fmt.Sprintf("Completed in %v seconds!", stopTimer),
		false, true,
	)

	// Ending message.
	cgapp.ShowMessage(
		"",
		"Have a great project launch! :)",
		false, true,
	)

	return nil
}
