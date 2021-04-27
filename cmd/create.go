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
	Short:   "Create a new project via interactive UI or configuration file",
	Long:    "\nCreate a new project via interactive UI or configuration file.",
	Run:     runCreateCmd,
}

// runCreateCmd represents runner for the `create` command.
var runCreateCmd = func(cmd *cobra.Command, args []string) {
	// Start message.
	_ = cgapp.ShowMessage(
		"",
		fmt.Sprintf("Create a new project via Create Go App CLI v%v...", registry.CLIVersion),
		true, true,
	)

	// Start timer.
	startTimer := time.Now()

	// Start survey.
	if err := survey.Ask(
		registry.CreateQuestions, &createAnswers, survey.WithIcons(surveyIconsConfig),
	); err != nil {
		log.Fatal(cgapp.ShowMessage("error", err.Error(), true, true))
	}

	// Define variables for better display.
	backend = strings.Replace(createAnswers.Backend, "/", "_", -1)
	frontend = createAnswers.Frontend
	proxy = createAnswers.Proxy

	// If something went wrong, cancel and exit.
	if !createAnswers.AgreeCreation {
		log.Fatal(
			cgapp.ShowMessage(
				"error",
				"Creation of a new project was stopped. Run `cgapp create` once again!",
				true, true,
			),
		)
	}

	// Create backend files.
	_ = cgapp.ShowMessage("warning", "Create backend for your project...", true, true)
	if err := cgapp.GitClone(
		"backend",
		fmt.Sprintf("github.com/create-go-app/%v-go-template", backend),
	); err != nil {
		log.Fatal(cgapp.ShowMessage("error", err.Error(), true, true))
	}

	// Cleanup project.
	cgapp.RemoveFolders("backend", []string{".git", ".github"})

	// Show success report.
	_ = cgapp.ShowMessage(
		"success",
		fmt.Sprintf("Backend was created with template `%v`!", backend),
		true, false,
	)

	if frontend != "none" {
		// Create frontend files.
		_ = cgapp.ShowMessage("warning", "Create frontend for your project...", true, true)
		if err := cgapp.ExecCommand(
			"npm",
			[]string{"init", "@vitejs/app", "frontend", "--", "--template", frontend},
		); err != nil {
			log.Fatal(cgapp.ShowMessage("error", err.Error(), true, true))
		}

		// Cleanup project.
		cgapp.RemoveFolders("frontend", []string{".git", ".github"})

		// Show success report.
		_ = cgapp.ShowMessage(
			"success",
			fmt.Sprintf("Frontend was created with template `%v`!", frontend),
			true, false,
		)
	}

	if proxy != "none" {
		// Copy Ansible playbooks and roles from embedded file system.
		_ = cgapp.ShowMessage("warning", "Create Ansible roles...", true, true)
		if err := cgapp.CopyFromEmbeddedFS(
			&cgapp.EmbeddedFileSystem{
				Name:       registry.EmbedRoles,
				RootFolder: "roles",
				SkipDir:    false,
			},
		); err != nil {
			log.Fatal(cgapp.ShowMessage("error", err.Error(), true, true))
		}

		// Copy Ansible playbook, inventory and roles from embedded file system.
		_ = cgapp.ShowMessage(
			"warning",
			"Create Ansible inventory and playbook files...",
			true, true,
		)
		if err := cgapp.CopyFromEmbeddedFS(
			&cgapp.EmbeddedFileSystem{
				Name:       registry.EmbedTemplates,
				RootFolder: "templates",
				SkipDir:    true,
			},
		); err != nil {
			log.Fatal(cgapp.ShowMessage("error", err.Error(), true, true))
		}

		// Set template variables for Ansible playbook and inventory files.
		switch proxy {
		case "traefik-acme-ca":
			// Traefik with simple ACME challenge via Let's Encrypt CA server.
			// See: https://doc.traefik.io/traefik/https/acme/#caserver
			inventory = map[string]interface{}{}
			playbook = map[string]interface{}{}
		case "traefik-acme-dns":
			// Traefik with more complex ACME challenge via your DNS provider.
			// See: https://doc.traefik.io/traefik/https/acme/#dnschallenge
			inventory = map[string]interface{}{}
			playbook = map[string]interface{}{}
		case "nginx":
			// Nginx.
			// See: https://nginx.org/en/docs/http/configuring_https_servers.html
			inventory = map[string]interface{}{}
			playbook = map[string]interface{}{}
		case "haproxy":
			// HAProxy.
			// See: http://cbonte.github.io/haproxy-dconv/2.4/intro.html#3.3.2
			inventory = map[string]interface{}{}
			playbook = map[string]interface{}{}
		default:
			log.Fatal(cgapp.ShowMessage("error", "The proxy server has not been set!", true, true))
		}

		// Generate Ansible inventory file.
		if err := cgapp.GenerateFileFromTemplate("hosts.ini.tmpl", inventory); err != nil {
			log.Fatal(cgapp.ShowMessage("error", err.Error(), true, true))
		}

		// Generate Ansible playbook file.
		if err := cgapp.GenerateFileFromTemplate("playbook.yml.tmpl", playbook); err != nil {
			log.Fatal(cgapp.ShowMessage("error", err.Error(), true, true))
		}

		//
		_ = cgapp.ShowMessage(
			"success",
			fmt.Sprintf("Ansible inventory, playbook and roles for `%v` was created!", proxy),
			false, true,
		)
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
	_ = cgapp.ShowMessage(
		"success",
		fmt.Sprintf("Completed in %v seconds!", stopTimer),
		true, true,
	)

	// Ending message.
	if proxy != "none" {
		_ = cgapp.ShowMessage(
			"",
			"Please, fill out Ansible inventory file (`$PWD/hosts.ini`) before deploy!",
			true, true,
		)
	}
	_ = cgapp.ShowMessage(
		"",
		"A helpful documentation and next steps -> https://create-go.app/",
		false, true,
	)
}

func init() {
	rootCmd.AddCommand(createCmd)
}
