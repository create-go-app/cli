package app

import (
	"github.com/create-go-app/cli/v5/embed"
	"github.com/create-go-app/cli/v5/internal/config"
)

// App ...
type App struct {
	Config     *config.Config
	EmbedFiles *embed.Files
}

// New ...
func New(config *config.Config, embedFiles *embed.Files) (*App, error) {
	app := &App{
		Config:     config,
		EmbedFiles: embedFiles,
	}
	return app, nil
}
