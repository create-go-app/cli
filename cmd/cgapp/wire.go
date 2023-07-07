//go:build wireinject

package main

import (
	"github.com/google/wire"

	"github.com/create-go-app/cli/v4/embed"
	"github.com/create-go-app/cli/v4/internal/app"
	"github.com/create-go-app/cli/v4/internal/config"
)

// initialize provides dependency injection process by the "google/wire" package.
func initialize(path string) (*app.App, error) {
	wire.Build(config.New, embed.New, app.New)
	return &app.App{}, nil
}
