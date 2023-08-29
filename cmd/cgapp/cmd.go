package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/create-go-app/cli/v5/embed"
	"github.com/create-go-app/cli/v5/internal/helpers"
)

// runCmd ...
func runCmd() {
	// Start timer.
	start := time.Now()

	// Parse flags.
	flag.Parse()

	// Set 'create' and 'deploy' config paths from option '-c'.
	createConfigPath, deployConfigPath := flag.Arg(2), flag.Arg(2)

	helpers.PrintStyled("👋 Hello and welcome to the Create Go App CLI (v5)!", "", "margin-top-bottom")

	// Switch over flags arguments.
	switch flag.Arg(0) {
	case "init":

		// Check, if the default config file ('.cgapp.yml') is existing in the current dir.
		_, err := os.Stat(".cgapp.yml")
		if err == nil {
			// If exists, skip a generation process.
			helpers.PrintStyled("Please, fix error(s):", "error", "margin-bottom")
			helpers.PrintStyled("the configuration file '.cgapp.yml' is found in the current dir, cannot be overwritten (data erasure protection)", "", "margin-left")
			helpers.PrintStyled("For more information, see https://github.com/create-go-app/cli/wiki", "warning", "margin-top-bottom")
			os.Exit(1)
		}

		// If not exists, get data from embed files.
		embedDefaultConfig, err := embed.ConfigsFiles.ReadFile("configs/default.yml")
		if err != nil {
			helpers.PrintStyled("Please, fix error(s):", "error", "margin-bottom")
			helpers.PrintStyled(err.Error(), "", "margin-left")
			helpers.PrintStyled("For more information, see https://github.com/create-go-app/cli/wiki", "warning", "margin-top-bottom")
			os.Exit(1)
		}

		// Create a new config file.
		if err = helpers.MakeFile(".cgapp.yml", embedDefaultConfig); err != nil {
			helpers.PrintStyled("Please, fix error(s):", "error", "margin-bottom")
			helpers.PrintStyled(err.Error(), "", "margin-left")
			helpers.PrintStyled("For more information, see https://github.com/create-go-app/cli/wiki", "warning", "margin-top-bottom")
			os.Exit(1)
		}

		helpers.PrintStyled(fmt.Sprintf("Successfully generated '.cgapp.yml' config file in the current dir! Time elapsed: %.2fs", time.Since(start).Seconds()), "success", "margin-bottom")
		helpers.PrintStyled("Next steps:", "info", "margin-left")
		helpers.PrintStyled("Edit config file with your options and parameters", "info", "margin-left-2")
		helpers.PrintStyled("Make awesome backend, frontend and setting up proxy", "info", "margin-left-2")
		helpers.PrintStyled("Run deploy process for your project", "info", "margin-left-2")
		helpers.PrintStyled("For more information, see https://github.com/create-go-app/cli/wiki", "warning", "margin-top-bottom")
		os.Exit(0)

	case "create":

		// Check, if required tools (git, npm, docker) was installed on the local system.
		if err := helpers.CheckCLITools([]string{"git", "npm", "docker"}); err != nil {
			// If not installed, skip a generation process.
			helpers.PrintStyled("Please, fix error(s):", "error", "margin-bottom")
			helpers.PrintStyled(err.Error(), "", "margin-left")
			helpers.PrintStyled("For more information, see https://github.com/create-go-app/cli/wiki", "warning", "margin-top-bottom")
			os.Exit(1)
		}

		helpers.PrintStyled("Successfully checked required tools (git, npm, docker) on your local system!", "success", "")

		// Check, if '--config' flag has a valid path (or URL).
		if createConfigPath == "" {
			helpers.PrintStyled("Config path (or URL) is not set, try to found '.cgapp.yml' file in the current dir...", "info", "")

			// Check, if the default config file ('.cgapp.yml') is existing in the current dir.
			if err := helpers.CheckProjectStructure([]string{".cgapp.yml"}); err != nil {
				// If not exists, skip an initialization process.
				helpers.PrintStyled("Please, fix error(s):", "error", "margin-top-bottom")
				helpers.PrintStyled(err.Error(), "", "margin-left")
				helpers.PrintStyled("For more information, see https://github.com/create-go-app/cli/wiki", "warning", "margin-top-bottom")
				os.Exit(1)
			}

			// If exists, set '.cgapp.yml' (in the current dir) to the config path.
			createConfigPath = ".cgapp.yml"

			helpers.PrintStyled("Config file '.cgapp.yml' was found, continue process...", "info", "")
		}

		helpers.PrintStyled(fmt.Sprintf("Analyzing the given configuration file ('%s')...", createConfigPath), "info", "")

		// Initialize app with config path.
		app, err := initialize(createConfigPath)
		if err != nil {
			helpers.PrintStyled("Please, fix error(s):", "error", "margin-top-bottom")
			helpers.PrintStyled(err.Error(), "", "margin-left")
			helpers.PrintStyled("For more information, see https://github.com/create-go-app/cli/wiki", "warning", "margin-top-bottom")
			os.Exit(1)
		}

		helpers.PrintStyled(fmt.Sprintf("Successfully initialized '%s' project with the given configuration file!", app.Config.Project.Name), "success", "")
		helpers.PrintStyled("Start creating project... please, wait!", "info", "")

		// Create a new project.
		if err = app.Create(); err != nil {
			helpers.PrintStyled("Please, fix error(s):", "error", "margin-top-bottom")
			helpers.PrintStyled(err.Error(), "", "margin-left")
			helpers.PrintStyled("For more information, see https://github.com/create-go-app/cli/wiki", "warning", "margin-top-bottom")
			os.Exit(1)
		}

		helpers.PrintStyled(fmt.Sprintf("Successfully created project in the current dir! Time elapsed: %.2fs", time.Since(start).Seconds()), "success", "margin-bottom")
		helpers.PrintStyled("Next steps:", "info", "")
		helpers.PrintStyled("Make awesome backend, frontend and setting up proxy", "info", "margin-left")
		helpers.PrintStyled("Run deploy process for your project", "info", "margin-left")
		helpers.PrintStyled("For more information, see https://github.com/create-go-app/cli/wiki", "warning", "margin-top-bottom")
		os.Exit(0)

	case "deploy":

		// Check, if required tools (ansible, ansible-playbook) was installed on the local system.
		if err := helpers.CheckCLITools([]string{"ansible", "ansible-playbook"}); err != nil {
			// If not installed, skip a generation process.
			helpers.PrintStyled("Please, fix error(s):", "error", "margin-bottom")
			helpers.PrintStyled(err.Error(), "", "margin-left")
			helpers.PrintStyled("For more information, see https://github.com/create-go-app/cli/wiki", "warning", "margin-top-bottom")
			os.Exit(1)
		}

		helpers.PrintStyled("Successfully checked required project structure to deploy on the current dir!", "success", "")

		// Check, if '-c' flag has a valid path (or URL).
		if deployConfigPath == "" {
			helpers.PrintStyled("Config path (or URL) is not set, try to found '.cgapp.yml' file in the current dir...", "info", "")

			// Check, if the default config file ('.cgapp.yml') is existing in the current dir.
			if err := helpers.CheckProjectStructure([]string{".cgapp.yml"}); err != nil {
				// If not exists, skip an initialization process.
				helpers.PrintStyled("Please, fix error(s):", "error", "margin-top-bottom")
				helpers.PrintStyled(err.Error(), "", "margin-left")
				helpers.PrintStyled("For more information, see https://github.com/create-go-app/cli/wiki", "warning", "margin-top-bottom")
				os.Exit(1)
			}

			// If exists, set '.cgapp.yml' (in the current dir) to the config path.
			deployConfigPath = ".cgapp.yml"

			helpers.PrintStyled("Config file '.cgapp.yml' was found, continue process...", "info", "")
		}

		helpers.PrintStyled(fmt.Sprintf("Analyzing the given configuration file ('%s')...", deployConfigPath), "info", "")

		// Initialize app with a config path.
		app, err := initialize(deployConfigPath)
		if err != nil {
			helpers.PrintStyled("Please, fix error(s):", "error", "margin-top-bottom")
			helpers.PrintStyled(err.Error(), "", "margin-left")
			helpers.PrintStyled("For more information, see https://github.com/create-go-app/cli/wiki", "warning", "margin-top-bottom")
			os.Exit(1)
		}

		helpers.PrintStyled(fmt.Sprintf("Successfully initialized '%s' project with the given configuration file!", app.Config.Project.Name), "success", "")

		// Check, if required project structure (ansible, ansible-playbook) was created on the current dir.
		if err = helpers.CheckProjectStructure([]string{"backend"}); err != nil {
			// If not installed, skip a generation process.
			helpers.PrintStyled("Please, fix error(s):", "error", "margin-top-bottom")
			helpers.PrintStyled(err.Error(), "", "margin-left")
			helpers.PrintStyled("For more information, see https://github.com/create-go-app/cli/wiki", "warning", "margin-top-bottom")
			os.Exit(1)
		}

		helpers.PrintStyled("Start deploying project... please, wait!", "info", "")

		// Deploy project to your remote server.
		if err = app.Deploy(); err != nil {
			helpers.PrintStyled("Please, fix error(s):", "error", "margin-top-bottom")
			helpers.PrintStyled(err.Error(), "", "margin-left")
			helpers.PrintStyled("For more information, see https://github.com/create-go-app/cli/wiki", "warning", "margin-top-bottom")
			os.Exit(1)
		}

		helpers.PrintStyled(fmt.Sprintf("Successfully deployed project to the remote server! Time elapsed: %.2fs", time.Since(start).Seconds()), "success", "margin-top-bottom")
		os.Exit(0)

	default:

		helpers.PrintStyled("cgapp [COMMAND] [OPTIONS]", "", "margin-bottom")
		helpers.PrintStyled("Available commands and options:", "success", "margin-bottom")
		helpers.PrintStyled("init      generate an example config file ('.cgapp.yml') file in the current dir", "", "margin-left")
		helpers.PrintStyled("create    create a new project in the current dir by the given config", "", "margin-left")
		helpers.PrintStyled("-c       path (or URL) to the config file (default: '.cgapp.yml')", "", "margin-left-2")
		helpers.PrintStyled("deploy    deploy your project to the remote host by the given config", "", "margin-left")
		helpers.PrintStyled("-c       path (or URL) to the config file (default: '.cgapp.yml')", "", "margin-left-2")
		helpers.PrintStyled("For more information, see https://github.com/create-go-app/cli/wiki", "warning", "margin-top-bottom")
		os.Exit(0)

	}
}
