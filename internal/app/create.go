package app

import (
	"github.com/create-go-app/cli/v5/internal/helpers"
)

// Create ...
func (app *App) Create() error {
	// TODO: implement app

	//
	if err := helpers.GenerateFileFromTemplate(
		app.EmbedFiles.Templates,
		"templates/hosts.ini.tmpl",
		"hosts.ini",
		app.Config,
	); err != nil {
		return err
	}

	//
	if err := helpers.GenerateFileFromTemplate(
		app.EmbedFiles.Templates,
		"templates/playbook.yml.tmpl",
		"playbook.yml",
		app.Config,
	); err != nil {
		return err
	}

	return nil
}
