// Copyright 2019-present Vic Shóstak. All rights reserved.
// Use of this source code is governed by Apache 2.0 license
// that can be found in the LICENSE file.

package cmd

import (
	"os"

	"github.com/create-go-app/cli/pkg/cgapp"
	"github.com/create-go-app/cli/pkg/registry"
	"github.com/spf13/cobra"
)

// initCmd represents the init command.
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Init a configuration file for the Create Go App project",
	Long:  "\nInit a configuration file for the Create Go App project.",
	Run:   runInitCmd,
}

// runInitCmd represents runner for the `init` command.
var runInitCmd = func(cmd *cobra.Command, args []string) {
	// Get current directory.
	currentDir, _ := os.Getwd()

	// Start message.
	cgapp.SendMsg(true, "* * *", "Init a configuration file in `"+currentDir+"` folder...", "yellow", true)

	// Create configuration file.
	fileToMake := map[string][]byte{
		".cgapp.yml": registry.EmbedCGAPPConfig,
	}
	if err := cgapp.MakeFiles(currentDir, fileToMake); err != nil {
		cgapp.SendMsg(true, "[ERROR]", err.Error(), "red", true)
		os.Exit(1)
	}

	// End message.
	cgapp.SendMsg(true, "(i)", "A helpful documentation and next steps -> https://create-go.app/", "green", false)
	cgapp.SendMsg(false, "(i)", "Run `cgapp create --use-config` to create a new project by this configuration file.", "green", true)
}

func init() {
	rootCmd.AddCommand(initCmd)
}
