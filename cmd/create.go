// Copyright 2022 Vic Shóstak and Create Go App Contributors. All rights reserved.
// Use of this source code is governed by Apache 2.0 license
// that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"

	"github.com/create-go-app/cli/v3/pkg/cgapp"
	"github.com/create-go-app/cli/v3/pkg/registry"
)

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().BoolVarP(
		&useCustomTemplate,
		"template", "t", false,
		"enables to use custom backend and frontend templates",
	)
}

// createCmd represents the `create` command.
var createCmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"new"},
	Short:   "Create a new project via interactive UI",
	Long:    "\nCreate a new project via interactive UI.",
	RunE:    runCreateCmd,
}

// runCreateCmd represents runner for the `create` command.
func runCreateCmd(cmd *cobra.Command, args []string) error {
	// Start message.
	cgapp.ShowMessage(
		"",
		fmt.Sprintf(
			"Create a new project via Create Go App CLI v%v...",
			registry.CLIVersion,
		),
		true, true,
	)

	// Start survey.
	if useCustomTemplate {
		// Custom survey.
		if err := survey.Ask(
			registry.CustomCreateQuestions,
			&customCreateAnswers,
			survey.WithIcons(surveyIconsConfig),
		); err != nil {
			return cgapp.ShowError(err.Error())
		}

		// Define variables for better display.
		backend = customCreateAnswers.Backend
		frontend = customCreateAnswers.Frontend
		proxy = customCreateAnswers.Proxy
	} else {
		// Default survey.
		if err := survey.Ask(
			registry.CreateQuestions,
			&createAnswers,
			survey.WithIcons(surveyIconsConfig),
		); err != nil {
			return cgapp.ShowError(err.Error())
		}

		// Define variables for better display.
		backend = fmt.Sprintf(
			"github.com/create-go-app/%v-go-template",
			strings.ReplaceAll(createAnswers.Backend, "/", "_"),
		)
		frontend = createAnswers.Frontend
		proxy = createAnswers.Proxy
	}

	// Catch the cancel action (hit "n" in the last question).
	if (!createAnswers.AgreeCreation && !useCustomTemplate) || (!customCreateAnswers.AgreeCreation && useCustomTemplate) {
		cgapp.ShowMessage(
			"",
			"Oh no! You said \"no\", so I won't create anything. Hope to see you soon!",
			true, true,
		)
		return nil
	}

	// Start timer.
	startTimer := time.Now()

	/*
		The project's backend part creation.
	*/

	// Clone backend files from git repository.
	if err := cgapp.GitClone("backend", backend); err != nil {
		return cgapp.ShowError(err.Error())
	}

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
		// Checking, if use custom templates.
		if useCustomTemplate {
			// Clone frontend files from git repository.
			if err := cgapp.GitClone("frontend", frontend); err != nil {
				return cgapp.ShowError(err.Error())
			}
		} else {
			switch {
			case frontend == "next" || frontend == "next-ts":
				var isTypeScript string
				if frontend == "next-ts" {
					isTypeScript = "--typescript"
				}

				// Create a default frontend template with Next.js (React).
				if err := cgapp.ExecCommand(
					"npx", []string{"create-next-app@latest", "frontend", isTypeScript}, true,
				); err != nil {
					return err
				}
			case frontend == "nuxt3":
				// Create a default frontend template with Nuxt 3 (Vue.js 3, TypeScript).
				if err := cgapp.ExecCommand(
					"npx", []string{"nuxi", "init", "frontend"}, true,
				); err != nil {
					return err
				}
			default:
				// Create a default frontend template from Vite (Pure JS/TS, React, Preact, Vue, Svelte, Lit).
				if err := cgapp.ExecCommand(
					"npm", []string{"init", "vite@latest", "frontend", "--", "--template", frontend}, true,
				); err != nil {
					return err
				}
			}
		}

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

	// Copy Ansible playbook, inventory and roles from embedded file system.
	if err := cgapp.CopyFromEmbeddedFS(
		&cgapp.EmbeddedFileSystem{
			Name:       registry.EmbedTemplates,
			RootFolder: "templates",
			SkipDir:    true,
		},
	); err != nil {
		return cgapp.ShowError(err.Error())
	}

	// Set template variables for Ansible playbook and inventory files.
	inventory = registry.AnsibleInventoryVariables[proxy].List
	playbook = registry.AnsiblePlaybookVariables[proxy].List

	// Generate Ansible inventory file.
	if err := cgapp.GenerateFileFromTemplate("hosts.ini.tmpl", inventory); err != nil {
		return cgapp.ShowError(err.Error())
	}

	// Generate Ansible playbook file.
	if err := cgapp.GenerateFileFromTemplate("playbook.yml.tmpl", playbook); err != nil {
		return cgapp.ShowError(err.Error())
	}

	// Show success report.
	if proxy != "none" {
		cgapp.ShowMessage(
			"success",
			fmt.Sprintf("Web/Proxy server configuration for `%v` was created!", proxy),
			false, false,
		)
	}

	/*
		The project's Ansible roles part creation.
	*/

	// Copy Ansible roles from embedded file system.
	if err := cgapp.CopyFromEmbeddedFS(
		&cgapp.EmbeddedFileSystem{
			Name:       registry.EmbedRoles,
			RootFolder: "roles",
			SkipDir:    false,
		},
	); err != nil {
		return cgapp.ShowError(err.Error())
	}

	// Show success report.
	cgapp.ShowMessage(
		"success",
		"Ansible inventory, playbook and roles for deploying your project was created!",
		false, false,
	)

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
		return cgapp.ShowError(err.Error())
	}

	/*
		Cleanup project.
	*/

	// Set unused proxy roles.
	switch proxy {
	case "traefik", "traefik-acme-dns":
		proxyList = []string{"nginx"}
	case "nginx":
		proxyList = []string{"traefik"}
	default:
		proxyList = []string{"traefik", "nginx"}
	}

	// Delete unused roles, backend and frontend files.
	cgapp.RemoveFolders("roles", proxyList)
	cgapp.RemoveFolders("backend", []string{".git", ".github"})
	cgapp.RemoveFolders("frontend", []string{".git", ".github"})

	// Stop timer.
	stopTimer := cgapp.CalculateDurationTime(startTimer)
	cgapp.ShowMessage(
		"info",
		fmt.Sprintf("Completed in %v seconds!", stopTimer),
		true, true,
	)

	// Ending messages.
	cgapp.ShowMessage(
		"",
		"* Please put credentials into the Ansible inventory file (`hosts.ini`) before you start deploying a project!",
		false, false,
	)
	if !useCustomTemplate && frontend != "none" {
		cgapp.ShowMessage(
			"",
			fmt.Sprintf("* Visit https://vitejs.dev/guide/ for more info about using the `%v` frontend template!", frontend),
			false, false,
		)
	}
	cgapp.ShowMessage(
		"",
		"* A helpful documentation and next steps with your project is here https://create-go.app/wiki",
		false, true,
	)
	cgapp.ShowMessage(
		"",
		"Have a happy new project! :)",
		false, true,
	)

	return nil
}
