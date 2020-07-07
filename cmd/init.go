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
	"fmt"
	"os"

	"github.com/create-go-app/cli/internal/embed"
	"github.com/create-go-app/cli/pkg/utils"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Init a configuration file for the Create Go App project",
	Run:   runInitCommand,
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func runInitCommand(cmd *cobra.Command, args []string) {
	// Get current directory.
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	// Start message.
	utils.SendMsg(true, "[~]", "Init a configuration file in `"+currentDir+"` folder...", "yellow", true)

	// Create configuration file.
	fileToMake := map[string][]byte{
		".cgapp.yml": embed.Get("/.cgapp.yml"),
	}
	utils.MakeFiles(currentDir, fileToMake)

	// End message.
	utils.SendMsg(true, "[?]", "A helpful documentation and next steps -> https://create-go.app/", "cyan", false)
	utils.SendMsg(false, "[!]", "Run `cgapp create -c` to create a new project by this configuration file.", "cyan", true)
}
