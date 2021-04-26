// Copyright 2019-present Vic Shóstak. All rights reserved.
// Use of this source code is governed by Apache 2.0 license
// that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"log"
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
	_ = cgapp.ShowMessage("", "Create a new project via Create Go App CLI v"+registry.CLIVersion+"...", true, true)

	// Start timer.
	startTimer := time.Now()

	// Start survey.
	if err := survey.Ask(
		registry.CreateQuestions, &createAnswers, survey.WithIcons(surveyIconsConfig),
	); err != nil {
		log.Fatal(cgapp.ShowMessage("error", err.Error(), true, true))
	}

	// If something went wrong, cancel and exit.
	if !createAnswers.AgreeCreation {
		log.Fatal(
			cgapp.ShowMessage("error", "Creation of a new project was stopped. Run `cgapp create` once again!", true, true),
		)
	}

	// Insert empty line.
	_ = cgapp.ShowMessage("", "", true, false)

	// Define variables for better display.
	backend = strings.ToLower(createAnswers.Backend)
	frontend = strings.ToLower(createAnswers.Frontend)
	proxy = strings.ToLower(createAnswers.Proxy)

	// Get current directory.
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(cgapp.ShowMessage("error", err.Error(), true, true))
	}

	// Create backend files.
	_ = cgapp.ShowMessage("warning", "Create project backend...", true, true)
	if err := cgapp.CreateProjectFromRegistry(
		&registry.Project{
			Type:       "backend",
			Name:       backend,
			RootFolder: currentDir,
		},
		registry.Repositories,
		registry.RegexpBackendPattern,
	); err != nil {
		log.Fatal(cgapp.ShowMessage("error", err.Error(), true, true))
	}

	if frontend != "none" {
		// Create frontend files.
		_ = cgapp.ShowMessage("warning", "Create project frontend...", true, true)
		if err := cgapp.CreateProjectFromCmd(
			&registry.Project{
				Type:       "frontend",
				Name:       frontend,
				RootFolder: currentDir,
			},
			registry.Commands,
			registry.RegexpFrontendPattern,
		); err != nil {
			log.Fatal(cgapp.ShowMessage("error", err.Error(), true, true))
		}
	}

	if proxy != "none" {
		//
		switch proxy {
		case "traefik":
			inventoryVariables = map[string]interface{}{}
			playbookVariables = map[string]interface{}{}
		case "traefik (with dns challenge)":
			inventoryVariables = map[string]interface{}{}
			playbookVariables = map[string]interface{}{}
		default:
			log.Fatal(cgapp.ShowMessage("error", "The proxy server has not been set!", true, true))
		}

		// Copy Ansible playbook, inventory and roles from embedded file system, if not skipped.
		_ = cgapp.ShowMessage("warning", "Create Ansible playbook and inventory...", true, true)
		if err := cgapp.CopyFromEmbeddedFS(
			&cgapp.EmbeddedFileSystem{
				Name:       registry.EmbedTemplates,
				RootFolder: "templates",
				SkipDir:    true,
			},
		); err != nil {
			log.Fatal(cgapp.ShowMessage("error", err.Error(), true, true))
		}

		//
		if err := cgapp.GenerateFileFromTemplate("hosts.ini.tmpl", inventoryVariables); err != nil {
			log.Fatal(cgapp.ShowMessage("error", err.Error(), true, true))
		}

		//
		if err := cgapp.GenerateFileFromTemplate("playbook.yml.tmpl", playbookVariables); err != nil {
			log.Fatal(cgapp.ShowMessage("error", err.Error(), true, true))
		}

		// Copy Ansible playbooks and roles from embedded file system, if not skipped.
		_ = cgapp.ShowMessage("warning", "Create Ansible roles for deploy...", true, true)
		if err := cgapp.CopyFromEmbeddedFS(
			&cgapp.EmbeddedFileSystem{
				Name:       registry.EmbedRoles,
				RootFolder: "roles",
				SkipDir:    false,
			},
		); err != nil {
			log.Fatal(cgapp.ShowMessage("error", err.Error(), true, true))
		}

		//
		_ = cgapp.ShowMessage("info", "Please, fill out Ansible inventory file (`hosts.ini`) before deploy!", false, true)
	}

	// Copy misc files from embedded file system.
	_ = cgapp.ShowMessage("warning", "Create misc files for your project...", true, true)
	if err := cgapp.CopyFromEmbeddedFS(
		&cgapp.EmbeddedFileSystem{
			Name:       registry.EmbedMiscFiles,
			RootFolder: "misc",
			SkipDir:    true,
		},
	); err != nil {
		log.Fatal(cgapp.ShowMessage("error", err.Error(), true, true))
	}

	// Stop timer.
	stopTimer := fmt.Sprintf("%.0f", time.Since(startTimer).Seconds())

	// End message.
	_ = cgapp.ShowMessage("success", "Completed in "+stopTimer+" seconds!", true, true)
	_ = cgapp.ShowMessage("", "A helpful documentation and next steps -> https://create-go.app/", false, true)
}

func init() {
	rootCmd.AddCommand(createCmd)
}
