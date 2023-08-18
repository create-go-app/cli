// Package main includes call for the Create Go App CLI.
//
// Create a new production-ready project with backend (Go),
// frontend (JavaScript, TypeScript) and deploy automation
// (Ansible, Terraform, Docker) by running one CLI command.
//
// Focus on writing your code and thinking of the business logic!
// The Create Go App CLI will take care of the rest.
//
// A helpful documentation and next steps:
// https://github.com/create-go-app/cli
//
// # Copyright 2023 Vic Shóstak and Create Go App Contributors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/create-go-app/cli/v5/embed"
	"github.com/create-go-app/cli/v5/internal/helpers"
)

func main() {
	// Define and parse app flags.
	initDefaultConfig := flag.Bool("init", false, "generate an example config file ('.cgapp.yml') file in the current dir")
	createProject := flag.Bool("create", false, "create a new project in the current dir")
	deployProject := flag.Bool("deploy", false, "deploy your project to the remote host")
	configPath := flag.String("p", "", "set a path (or URL) to the config file")
	flag.Parse()

	// Start timer.
	start := time.Now()

	helpers.PrintStyled("👋 Hello and welcome to the Create Go App CLI (v5)!", "", "margin-top")

	// Check, if '-init' flag is true.
	if *initDefaultConfig {
		// Check, if the default config file ('.cgapp.yml') is existing in the current dir.
		_, err := os.Stat(".cgapp.yml")
		if err == nil {
			// If exists, skip a generation process.
			helpers.PrintStyled("Please, fix error(s):", "error", "margin-top-bottom")
			helpers.PrintStyled("the configuration file '.cgapp.yml' is found in the current dir, cannot be overwritten (data erasure protection)", "", "margin-left")
			helpers.PrintStyled("For more information, see https://github.com/create-go-app/cli/wiki", "warning", "margin-top-bottom")
			os.Exit(1)
		}

		// If not exists, get data from embed files.
		embedDefaultConfig, err := embed.ConfigsFiles.ReadFile("configs/default.yml")
		if err != nil {
			helpers.PrintStyled("Please, fix error(s):", "error", "margin-top-bottom")
			helpers.PrintStyled(err.Error(), "", "margin-left")
			helpers.PrintStyled("For more information, see https://github.com/create-go-app/cli/wiki", "warning", "margin-top-bottom")
			os.Exit(1)
		}

		// Create a new config file.
		if err = helpers.MakeFile(".cgapp.yml", embedDefaultConfig); err != nil {
			helpers.PrintStyled("Please, fix error(s):", "error", "margin-top-bottom")
			helpers.PrintStyled(err.Error(), "", "margin-left")
			helpers.PrintStyled("For more information, see https://github.com/create-go-app/cli/wiki", "warning", "margin-top-bottom")
			os.Exit(1)
		}

		helpers.PrintStyled(fmt.Sprintf("Successfully generated '.cgapp.yml' config file in the current dir! Time elapsed: %.2fs", time.Since(start).Seconds()), "success", "margin-top-bottom")
		helpers.PrintStyled("Next steps:", "info", "margin-left")
		helpers.PrintStyled("Edit config file with your options and parameters", "info", "margin-left-2")
		helpers.PrintStyled("Make awesome backend, frontend and setting up proxy", "info", "margin-left-2")
		helpers.PrintStyled("Run deploy process for your project", "info", "margin-left-2")
		helpers.PrintStyled("For more information, see https://github.com/create-go-app/cli/wiki", "warning", "margin-top-bottom")
		os.Exit(0)
	}

	// Check, if required tools (git, npm, docker) was installed on the local system.
	if err := helpers.CheckCLITools([]string{"git", "npm", "docker"}); err != nil {
		// If not installed, skip a generation process.
		helpers.PrintStyled("Please, fix error(s):", "error", "margin-top-bottom")
		helpers.PrintStyled(err.Error(), "", "margin-left")
		helpers.PrintStyled("For more information, see https://github.com/create-go-app/cli/wiki", "warning", "margin-top-bottom")
		os.Exit(1)
	}

	helpers.PrintStyled("Successfully checked required tools (git, npm, docker) on your local system!", "success", "margin-top")

	// Check, if '-p' flag has path (or URL).
	if *configPath == "" {
		helpers.PrintStyled("Config path (or URL) is not set, try to found '.cgapp.yml' file in the current dir...", "info", "margin-top")

		// Check, if the default config file ('.cgapp.yml') is existing in the current dir.
		_, err := os.Stat(".cgapp.yml")
		if err != nil {
			// If not exists, skip an initialization process.
			helpers.PrintStyled("Please, fix error(s):", "error", "margin-top-bottom")
			helpers.PrintStyled("the configuration file '.cgapp.yml' is not found", "", "margin-left")
			helpers.PrintStyled("For more information, see https://github.com/create-go-app/cli/wiki", "warning", "margin-top-bottom")
			os.Exit(1)
		}

		// If exists, set '.cgapp.yml' (in the current dir) to the config path.
		*configPath = ".cgapp.yml"

		helpers.PrintStyled("Config file '.cgapp.yml' was found, continue process...", "info", "margin-top-bottom")
	}

	helpers.PrintStyled(fmt.Sprintf("Analyzing the given configuration file ('%s')...", *configPath), "info", "margin-top")

	// Initialize app with config path.
	app, err := initialize(*configPath)
	if err != nil {
		helpers.PrintStyled("Please, fix error(s):", "error", "margin-top-bottom")
		helpers.PrintStyled(err.Error(), "", "margin-left")
		helpers.PrintStyled("For more information, see https://github.com/create-go-app/cli/wiki", "warning", "margin-top-bottom")
		os.Exit(1)
	}

	helpers.PrintStyled(fmt.Sprintf("Successfully initialized '%s' project with the given configuration file...", app.Config.Project.Name), "success", "margin-top")

	// Check, if '-create' flag is true.
	if *createProject {
		helpers.PrintStyled("Start creating project... please, wait!", "info", "margin-top")

		// Create a new project.
		if err = app.Create(); err != nil {
			helpers.PrintStyled("Please, fix error(s):", "error", "margin-top-bottom")
			helpers.PrintStyled(err.Error(), "", "margin-left")
			helpers.PrintStyled("For more information, see https://github.com/create-go-app/cli/wiki", "warning", "margin-top-bottom")
			os.Exit(1)
		}

		helpers.PrintStyled(fmt.Sprintf("Successfully created project in the current dir! Time elapsed: %.2fs", time.Since(start).Seconds()), "success", "margin-top-bottom")
		helpers.PrintStyled("Next steps:", "info", "margin-left")
		helpers.PrintStyled("Make awesome backend, frontend and setting up proxy", "info", "margin-left-2")
		helpers.PrintStyled("Run deploy process for your project", "info", "margin-left-2")
		helpers.PrintStyled("For more information, see https://github.com/create-go-app/cli/wiki", "warning", "margin-top-bottom")
		os.Exit(0)
	}

	// Check, if '-deploy' flag is true.
	if *deployProject {
		helpers.PrintStyled("Start deploying project... please, wait!", "info", "margin-top")

		// Deploy project to your remote host.
		if err = app.Deploy(); err != nil {
			helpers.PrintStyled("Please, fix error(s):", "error", "margin-top-bottom")
			helpers.PrintStyled(err.Error(), "", "margin-left")
			helpers.PrintStyled("For more information, see https://github.com/create-go-app/cli/wiki", "warning", "margin-top-bottom")
			os.Exit(1)
		}

		helpers.PrintStyled(fmt.Sprintf("Successfully deployed project to the remote server! Time elapsed: %.2fs", time.Since(start).Seconds()), "success", "margin-top-bottom")
		os.Exit(0)
	}
}
