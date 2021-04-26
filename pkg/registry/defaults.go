// Copyright 2019-present Vic Shóstak. All rights reserved.
// Use of this source code is governed by Apache 2.0 license
// that can be found in the LICENSE file.

package registry

import (
	"embed"

	"github.com/AlecAivazis/survey/v2"
)

const (
	// CLIVersion version of Create Go App CLI.
	CLIVersion = "2.0.0"
	// RegexpBackendPattern pattern for backend.
	RegexpBackendPattern = "^(net/http|fiber)$"
	// RegexpFrontendPattern pattern for backend.
	RegexpFrontendPattern = "^(p?react:?|vue(:?[\\w]+)?(:?[\\w-_0-9\\/]+)?|angular|svelte|sapper:?)"
)

// Project struct for describe project.
type Project struct {
	Type       string
	Name       string
	RootFolder string
}

// Repository struct for describe repositories collection.
type Repository struct {
	List map[string]string
}

// Command struct for describe commands collection.
type Command struct {
	Runner string
	Create string
	Args   map[string]string
}

// CreateAnswers struct for a survey's answers for `create` command.
type CreateAnswers struct {
	Backend       string
	Frontend      string
	Proxy         string
	AgreeCreation bool `survey:"agree"`
}

var (
	// EmbedMiscFiles misc files and configs.
	//go:embed misc/*
	EmbedMiscFiles embed.FS

	// EmbedRoles Ansible roles.
	//go:embed roles/*
	EmbedRoles embed.FS

	// EmbedTemplates template files.
	//go:embed templates/*
	EmbedTemplates embed.FS

	// Repositories collection.
	Repositories = map[string]*Repository{
		// Backend templates.
		"backend": {
			List: map[string]string{
				"net/http": "github.com/create-go-app/net_http-go-template",
				"fiber":    "github.com/create-go-app/fiber-go-template",
			},
		},
	}

	// Commands collection.
	Commands = map[string]*Command{
		"react": {
			Runner: "npx",
			Create: "create-react-app",
			Args: map[string]string{
				"template": "--template",
			},
		},
		"preact": {
			Runner: "preact",
			Create: "create",
			Args: map[string]string{
				"cwd":  "--cwd",
				"name": "--name",
			},
		},
		"vue": {
			Runner: "vue",
			Create: "create",
			Args:   map[string]string{},
		},
		"angular": {
			Runner: "ng",
			Create: "new",
			Args:   map[string]string{},
		},
		"svelte": {
			Runner: "npx",
			Create: "degit",
			Args: map[string]string{
				"template": "sveltejs/template",
			},
		},
		"sapper": {
			Runner: "npx",
			Create: "degit",
			Args: map[string]string{
				"template": "sveltejs/sapper-template",
			},
		},
	}

	// CreateQuestions survey's questions for `create` command.
	CreateQuestions = []*survey.Question{
		{
			Name: "backend",
			Prompt: &survey.Select{
				Message: "Choose a backend framework:",
				Options: []string{
					"net/http",
					"Fiber",
				},
				Default: "Fiber",
			},
			Validate: survey.Required,
		},
		{
			Name: "frontend",
			Prompt: &survey.Select{
				Message: "Choose a frontend UI library:",
				Options: []string{
					"none",
					"React",
					"Preact",
					"Vue",
					"Angular",
					"Svelte",
					"Sapper",
				},
				Default: "none",
			},
		},
		{
			Name: "proxy",
			Prompt: &survey.Select{
				Message: "Choose a proxy server:",
				Options: []string{
					"none",
					"Traefik",
					"Traefik (with DNS challenge)",
					"Nginx",
				},
				Default: "traefik",
			},
		},
		{
			Name: "agree",
			Prompt: &survey.Confirm{
				Message: "If everything is okay, can I create this project for you? ;)",
				Default: true,
			},
		},
	}
)
