/*
Package cmd includes all of the Create Go App CLI commands.

Copyright © 2020 Vic Shóstak <truewebartisans@gmail.com> (https://1wa.co)

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"os"

	"github.com/create-go-app/cli/pkg/embed"
	"github.com/create-go-app/cli/pkg/utils"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:     "init",
	Aliases: []string{"i"},
	Short:   "Init a configuration file for the Create Go App project",
	Long:    "\nInit a configuration file for the Create Go App project.",
	Run: func(cmd *cobra.Command, args []string) {
		// Get current directory.
		currentDir, _ := os.Getwd()

		// Start message.
		utils.SendMsg(true, "* * *", "Init a configuration file in `"+currentDir+"` folder...", "yellow", true)

		// Create configuration file.
		fileToMake := map[string][]byte{
			".cgapp.yml": embed.Get("/.cgapp.yml"),
		}
		if err := utils.MakeFiles(currentDir, fileToMake); err != nil {
			utils.SendMsg(true, "[ERROR]", err.Error(), "red", true)
			os.Exit(1)
		}

		// End message.
		utils.SendMsg(true, "(i)", "A helpful documentation and next steps -> https://create-go.app/", "green", false)
		utils.SendMsg(false, "(i)", "Run `cgapp create --use-config` to create a new project by this configuration file.", "green", true)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
