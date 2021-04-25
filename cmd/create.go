// Copyright 2019-present Vic Shóstak. All rights reserved.
// Use of this source code is governed by Apache 2.0 license
// that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/create-go-app/cli/pkg/cgapp"
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

	// Start timer.
	startTimer := time.Now()

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
	installAnsibleRoles = createAnswers.InstallAnsibleRoles

	// Get current directory.
	currentDir, err := os.Getwd()
	if err != nil {
		cgapp.SendMsg(true, "[ERROR]", err.Error(), "red", true)
		os.Exit(1)
	}

	// Copy misc files from embedded file system.
	cgapp.SendMsg(false, "*", "Create misc files for your project...", "cyan", true)
	if err := cgapp.CopyFromEmbeddedFS(
		&cgapp.EmbeddedFileSystem{
			Name:       registry.EmbedMiscFiles,
			RootFolder: "misc",
			SkipDir:    true,
		},
	); err != nil {
		cgapp.SendMsg(true, "[ERROR]", err.Error(), "red", true)
		os.Exit(1)
	}

	if installAnsibleRoles {
		// Copy Ansible playbooks and roles from embedded file system, if not skipped.
		cgapp.SendMsg(true, "*", "Create Ansible playbooks and roles...", "cyan", true)
		if err := cgapp.CopyFromEmbeddedFS(
			&cgapp.EmbeddedFileSystem{
				Name:       registry.EmbedRoles,
				RootFolder: "roles",
				SkipDir:    false,
			},
		); err != nil {
			cgapp.SendMsg(true, "[ERROR]", err.Error(), "red", true)
			os.Exit(1)
		}
	}

	// Create backend files.
	cgapp.SendMsg(true, "*", "Create project backend...", "cyan", true)
	if err := cgapp.CreateProjectFromRegistry(
		&registry.Project{
			Type:       "backend",
			Name:       backend,
			RootFolder: currentDir,
		},
		registry.Repositories,
		registry.RegexpBackendPattern,
	); err != nil {
		cgapp.SendMsg(true, "[ERROR]", err.Error(), "red", true)
		os.Exit(1)
	}

	if frontend != "none" {
		// Create frontend files.
		cgapp.SendMsg(true, "*", "Create project frontend...", "cyan", false)
		if err := cgapp.CreateProjectFromCmd(
			&registry.Project{
				Type:       "frontend",
				Name:       frontend,
				RootFolder: currentDir,
			},
			registry.Commands,
			registry.RegexpFrontendPattern,
		); err != nil {
			cgapp.SendMsg(true, "[ERROR]", err.Error(), "red", true)
			os.Exit(1)
		}
	}

	// Stop timer.
	stopTimer := fmt.Sprintf("%.0f", time.Since(startTimer).Seconds())

	// End message.
	cgapp.SendMsg(true, "* * *", "Completed in "+stopTimer+" seconds!", "yellow", true)
	cgapp.SendMsg(false, "(i)", "A helpful documentation and next steps -> https://create-go.app/", "green", false)
	cgapp.SendMsg(false, "(i)", "Run `cgapp deploy [mode]` to deploy your project to a remote server or run on localhost.", "green", true)
}

func init() {
	rootCmd.AddCommand(createCmd)
}
