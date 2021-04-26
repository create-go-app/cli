// Copyright 2019-present Vic Shóstak. All rights reserved.
// Use of this source code is governed by Apache 2.0 license
// that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/create-go-app/cli/pkg/cgapp"
	"github.com/create-go-app/cli/pkg/registry"
	"github.com/spf13/cobra"
)

// generateCmd represents the `generate` command.
var generateCmd = &cobra.Command{
	Use:     "generate",
	Aliases: []string{"gen", "init"},
	Short:   "Generate a new Ansible inventory file for deploy the Create Go App project",
	Long:    "\nGenerate a new Ansible inventory file for deploy the Create Go App project.",
	Run:     runGenerateCmd,
}

// runGenerateCmd represents runner for the `generate` command.
var runGenerateCmd = func(cmd *cobra.Command, args []string) {
	// Start message.
	cgapp.SendMsg(true, "* * *", "Generate a new Ansible inventory file in current folder...", "yellow", true)

	// Start timer.
	startTimer := time.Now()

	switch generateDeployConfig {
	case "traefik":
		//

	case "traefik:wildcards":
		//

	case "nginx":
		//

	case "haproxy":
		//

	}

	// Copy Create Go App CLI configuration file from embedded file system.
	if err := cgapp.CopyFromEmbeddedFS(
		&cgapp.EmbeddedFileSystem{
			Name:       registry.EmbedConfigs,
			RootFolder: "config",
			SkipDir:    true,
		},
	); err != nil {
		cgapp.SendMsg(true, "[ERROR]", err.Error(), "red", true)
		os.Exit(1)
	}

	// Stop timer.
	stopTimer := fmt.Sprintf("%.0f", time.Since(startTimer).Seconds())

	// End message.
	cgapp.SendMsg(true, "* * *", "Completed in "+stopTimer+" seconds!", "yellow", true)
	cgapp.SendMsg(false, "(i)", "A helpful documentation and next steps -> https://create-go.app/", "green", false)
	cgapp.SendMsg(false, "(i)", "Please, fill out this Ansible inventory file (`hosts.ini`) before deploy!", "green", false)
	cgapp.SendMsg(false, "(i)", "Then run `cgapp deploy` to deploy your project.", "green", true)
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.PersistentFlags().StringVarP(
		&generateDeployConfig,
		"proxy", "p", "",
		"generate a new Ansible inventory file (`$PWD/hosts.ini`) for specified proxy",
	)
	if err := rootCmd.MarkPersistentFlagRequired("proxy"); err != nil {
		log.Fatal(cgapp.BeautifyText(err.Error(), "red"))
	}
}
