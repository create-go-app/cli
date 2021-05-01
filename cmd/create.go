// Copyright 2019-present Vic Shóstak. All rights reserved.
// Use of this source code is governed by Apache 2.0 license
// that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"log"
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
	Short:   "Create a new project via interactive UI",
	Long:    "\nCreate a new project via interactive UI.",
	Run:     runCreateCmd,
}

// runCreateCmd represents runner for the `create` command.
var runCreateCmd = func(cmd *cobra.Command, args []string) {
	// Start message.
	cgapp.ShowMessage(
		"",
		fmt.Sprintf("Create a new project via Create Go App CLI v%v...", registry.CLIVersion),
		true, true,
	)

	// Start survey.
	if err := survey.Ask(
		registry.CreateQuestions, &createAnswers, survey.WithIcons(surveyIconsConfig),
	); err != nil {
		log.Fatal(cgapp.ShowError(err.Error()))
	}

	// Define variables for better display.
	backend = strings.Replace(createAnswers.Backend, "/", "_", -1)
	frontend = createAnswers.Frontend
	proxy = createAnswers.Proxy

	// If something went wrong, cancel and exit.
	if !createAnswers.AgreeCreation {
		log.Fatal(
			cgapp.ShowError("Creation of a new project was stopped. Run `cgapp create` again!"),
		)
	}

	// Start timer.
	startTimer := time.Now()

	/*
		The project's backend part creation.
	*/

	// Clone backend files from git repository.
	if err := cgapp.GitClone(
		"backend",
		fmt.Sprintf("github.com/create-go-app/%v-go-template", backend),
	); err != nil {
		log.Fatal(cgapp.ShowError(err.Error()))
	}

	// Cleanup project.
	cgapp.RemoveFolders("backend", []string{".git", ".github"})

	// Show success report.
	cgapp.ShowMessage(
		"success",
		fmt.Sprintf("Backend was created with template `%v`!", backend),
		true, false,
	)

	/*
		The project's frontend part creation.
	*/

	if frontend != "none" {
		// Create frontend files.
		if err := cgapp.ExecCommand(
			"npm",
			[]string{"init", "@vitejs/app", "frontend", "--", "--template", frontend},
			true,
		); err != nil {
			log.Fatal(cgapp.ShowError(err.Error()))
		}

		// Cleanup project.
		cgapp.RemoveFolders("frontend", []string{".git", ".github"})

		// Show success report.
		cgapp.ShowMessage(
			"success",
			fmt.Sprintf("Frontend was created with template `%v`!", frontend),
			false, false,
		)
	}

	/*
		The project's webserver part creation.
	*/

	if proxy != "none" {
		// Copy Ansible roles from embedded file system.
		if err := cgapp.CopyFromEmbeddedFS(
			&cgapp.EmbeddedFileSystem{
				Name:       registry.EmbedRoles,
				RootFolder: "roles",
				SkipDir:    false,
			},
		); err != nil {
			log.Fatal(cgapp.ShowError(err.Error()))
		}

		// Copy Ansible playbook, inventory and roles from embedded file system.
		if err := cgapp.CopyFromEmbeddedFS(
			&cgapp.EmbeddedFileSystem{
				Name:       registry.EmbedTemplates,
				RootFolder: "templates",
				SkipDir:    true,
			},
		); err != nil {
			log.Fatal(cgapp.ShowError(err.Error()))
		}

		// Set template variables for Ansible playbook and inventory files.
		inventory = registry.AnsibleInventoryVariables[proxy].List
		playbook = registry.AnsiblePlaybookVariables[proxy].List

		// Generate Ansible inventory file.
		if err := cgapp.GenerateFileFromTemplate("hosts.ini.tmpl", inventory); err != nil {
			log.Fatal(cgapp.ShowError(err.Error()))
		}

		// Generate Ansible playbook file.
		if err := cgapp.GenerateFileFromTemplate("playbook.yml.tmpl", playbook); err != nil {
			log.Fatal(cgapp.ShowError(err.Error()))
		}

		// Set unused proxy roles.
		if proxy == "traefik" || proxy == "traefik-acme-dns" {
			proxyList = []string{"nginx"}
		} else if proxy == "nginx" {
			proxyList = []string{"traefik"}
		}

		// Set unused proxy role.
		if frontend == "none" {
			proxyList = append(proxyList, "frontend")
		}

		// Delete unused proxy and/or frontend roles.
		cgapp.RemoveFolders("roles", proxyList)

		// Success messages.
		cgapp.ShowMessage(
			"success",
			fmt.Sprintf("Web/Proxy server configuration for `%v` was created!", proxy),
			false, false,
		)
		cgapp.ShowMessage(
			"success",
			"Ansible inventory, playbook and roles for deploying was created!",
			false, false,
		)
	}

	/*
		The project's misc files part creation.
	*/

	// Copy from embedded file system.
	if err := cgapp.CopyFromEmbeddedFS(
		&cgapp.EmbeddedFileSystem{
			Name:       registry.EmbedMiscFiles,
			RootFolder: "misc",
			SkipDir:    true,
		},
	); err != nil {
		log.Fatal(cgapp.ShowError(err.Error()))
	}

	// Stop timer.
	stopTimer := fmt.Sprintf("%.0f", time.Since(startTimer).Seconds())
	cgapp.ShowMessage(
		"info",
		fmt.Sprintf("Completed in %v seconds!", stopTimer),
		true, true,
	)

	// Ending messages.
	if proxy != "none" {
		cgapp.ShowMessage(
			"",
			"* Please put credentials into the Ansible inventory file (`hosts.ini`) before you start deploying a project!",
			false, false,
		)
	}
	if frontend != "none" {
		cgapp.ShowMessage(
			"",
			fmt.Sprintf("* Visit https://vitejs.dev/guide/ for more information about using the `%v` frontend template!", frontend),
			false, false,
		)
	}
	cgapp.ShowMessage(
		"",
		"* A helpful documentation and next steps -> https://create-go.app/",
		false, true,
	)
	cgapp.ShowMessage(
		"",
		"Have a happy new project! :)",
		false, true,
	)
}

func init() {
	rootCmd.AddCommand(createCmd)
}
