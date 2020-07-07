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
	"reflect"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	useConfigFile bool                   // indicate using config
	projectConfig map[string]interface{} // parse project config
	rolesConfig   map[string]interface{} // parse Ansible roles config
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new project via interactive UI or configuration file",
	Run:   runCreateCommand,
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().BoolVarP(&useConfigFile, "config", "c", true, "config file (default is $PWD/.cgapp.yml)")
}

// initConfig reads in config file, if set.
func initConfig() {
	if useConfigFile {
		// Get current directory
		currentDir, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
		}

		viper.AddConfigPath(currentDir) // add config path
		viper.SetConfigName(".cgapp")   // set config name

		// If a config file is found, read it in.
		if err := viper.ReadInConfig(); err != nil {
			fmt.Println(err)
		}

		// Parse configs
		_ = viper.UnmarshalKey("project", &projectConfig)
		_ = viper.UnmarshalKey("roles", &rolesConfig)
	}
}

// runInitCommand ...
func runCreateCommand(cmd *cobra.Command, args []string) {

	//
	fmt.Println(projectConfig["backend"], reflect.TypeOf(rolesConfig["deploy"]))
}
