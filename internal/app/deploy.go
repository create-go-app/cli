package app

import (
	"github.com/create-go-app/cli/v5/internal/executor"
	"github.com/create-go-app/cli/v5/internal/helpers"
)

// Deploy ...
func (app *App) Deploy() error {
	// Check, if a config file has 'ansible' block.
	if app.Config.Deploy.Ansible != nil {
		// Create hosts.ini file for Ansible.
		if err := helpers.GenerateFileFromTemplate(
			app.EmbedFiles.Templates,
			"templates/hosts.ini.tmpl",
			"hosts.ini",
			app.Config,
		); err != nil {
			return err
		}

		// Create playbook.yml file for Ansible.
		if err := helpers.GenerateFileFromTemplate(
			app.EmbedFiles.Templates,
			"templates/playbook.yml.tmpl",
			"playbook.yml",
			app.Config,
		); err != nil {
			return err
		}

		// Check, if an 'ansible' block has a 'become_sudo_user' option.
		if app.Config.Deploy.Ansible.BecomeSudoUser {
			// Start Ansible deploying process with sudo.
			if err := executor.Execute(
				"ansible-playbook", "playbook.yml", "-i", "hosts.ini", "-K",
			); err != nil {
				return err
			}
		} else {
			// Start Ansible deploying process without sudo.
			if err := executor.Execute(
				"ansible-playbook", "playbook.yml", "-i", "hosts.ini",
			); err != nil {
				return err
			}
		}
	}

	return nil
}
