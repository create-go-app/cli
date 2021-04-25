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
	// RegexpWebServerPattern pattern for web/proxy servers.
	RegexpWebServerPattern = "^(nginx|traefik)$"
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
	Backend             string
	Frontend            string
	InstallAnsibleRoles bool `survey:"roles"`
	AgreeCreation       bool `survey:"agree"`
}

// DeployAnswers struct for a survey's answers for `deploy` command.
type DeployAnswers struct {
	Username        string
	Host            string
	Network         string
	BackendPort     string
	AskBecomePass   bool `survey:"become_pass"`
	AgreeDeployment bool `survey:"agree"`
}

var (
	// EmbedConfigs configs for Create Go App CLI.
	//go:embed configs/*
	EmbedConfigs embed.FS

	// EmbedMiscFiles misc files and configs.
	//go:embed misc/*
	EmbedMiscFiles embed.FS

	// EmbedRoles Ansible roles.
	//go:embed roles/*
	EmbedRoles embed.FS

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
				Options: []string{"net/http", "Fiber"},
				Default: "Fiber",
			},
			Validate: survey.Required,
		},
		{
			Name: "frontend",
			Prompt: &survey.Select{
				Message: "Choose a frontend UI library:",
				Options: []string{"none", "React", "Preact", "Vue", "Angular", "Svelte", "Sapper"},
				Default: "none",
			},
		},
		{
			Name: "roles",
			Prompt: &survey.Confirm{
				Message: "Do you want to create Ansible playbook and roles for deploy your project?",
				Default: true,
			},
		},
		{
			Name: "agree",
			Prompt: &survey.Confirm{
				Message: "If everything is okay, can I create this project? ;)",
				Default: true,
			},
		},
	}

	// DeployQuestions survey's questions for `deploy` command.
	DeployQuestions = []*survey.Question{
		{
			Name: "project_domain_url",
			Prompt: &survey.Input{
				Message: "Enter domain for this project (e.g. example.com):",
			},
			Validate: survey.Required,
		},
		{
			Name: "letsencript_email",
			Prompt: &survey.Input{
				Message: "Enter your Email address for generating Let's Encrypt SSL certificate (e.g. mail@example.com):",
			},
			Validate: survey.Required,
		},
		{
			Name: "system_user_name",
			Prompt: &survey.Input{
				Message: "Enter system username:",
				Default: "root",
			},
			Validate: survey.Required,
		},
		{
			Name: "become_pass",
			Prompt: &survey.Confirm{
				Message: "Do you need to enter password for this username?",
				Default: true,
			},
		},
		{
			Name: "host_name",
			Prompt: &survey.Input{
				Message: "Enter host name to deploy (from Ansible inventory):",
				Default: "localhost",
			},
			Validate: survey.Required,
		},
		{
			Name: "docker_network",
			Prompt: &survey.Input{
				Message: "Enter name of the Docker network:",
				Default: "cgapp_network",
			},
			Validate: survey.Required,
		},
		{
			Name: "traefik_dashboard_password",
			Prompt: &survey.Password{
				Message: "Create a new password for Traefik dashboard:",
			},
			Validate: survey.Required,
		},
		{
			Name: "backend_port",
			Prompt: &survey.Input{
				Message: "Enter port number, using for backend Docker container:",
				Default: "5000",
			},
			Validate: survey.Required,
		},
		{
			Name: "agree",
			Prompt: &survey.Confirm{
				Message: "If everything is okay, can I deploy this project? ;)",
				Default: true,
			},
		},
	}
)
