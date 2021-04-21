// Copyright 2019-present Vic Shóstak. All rights reserved.
// Use of this source code is governed by Apache 2.0 license
// that can be found in the LICENSE file.

package registry

import (
	_ "embed" // for embed config files

	"github.com/AlecAivazis/survey/v2"
)

const (
	// CLIVersion version of Create Go App CLI.
	CLIVersion = "1.8.0"
	// RegexpAnsiblePattern pattern for Ansible roles.
	RegexpAnsiblePattern = "^(deploy)$"
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
	Webserver           string
	InstallAnsibleRoles bool `survey:"roles"`
	AgreeCreation       bool `survey:"agree"`
}

// DeployAnswers struct for a survey's answers for `deploy` command.
type DeployAnswers struct {
	Username        string
	Host            string
	Network         string
	BackendPort     string
	AskBecomePass   bool `survey:"become"`
	AgreeDeployment bool `survey:"agree"`
}

var (
	// EmbedCGAPPConfig main config file for Create Go App.
	//go:embed configs/.cgapp.yml
	EmbedCGAPPConfig []byte

	// EmbedEditorConfig config for auto configuration IDE.
	//go:embed configs/.editorconfig
	EmbedEditorConfig []byte

	// EmbedGitAttributes config for git attributes.
	//go:embed configs/.gitattributes
	EmbedGitAttributes []byte

	// EmbedGitIgnore config for git ignore files.
	//go:embed configs/.gitignore
	EmbedGitIgnore []byte

	// EmbedMakefile file for rapid manipulation with a new app.
	//go:embed configs/Makefile
	EmbedMakefile []byte

	// EmbedDeployPlaybook Ansible playbook for deployment.
	//go:embed configs/deploy-playbook.yml
	EmbedDeployPlaybook []byte

	// Repositories collection.
	Repositories = map[string]*Repository{
		// Backend templates.
		"backend": {
			List: map[string]string{
				"net/http": "github.com/create-go-app/net_http-go-template",
				"fiber":    "github.com/create-go-app/fiber-go-template",
			},
		},

		// Docker containers with web/proxy servers.
		"webserver": {
			List: map[string]string{
				"nginx":   "github.com/create-go-app/nginx-docker",
				"traefik": "github.com/create-go-app/traefik-docker",
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
			Name: "webserver",
			Prompt: &survey.Select{
				Message: "Choose a web/proxy server:",
				Options: []string{"none", "Nginx", "Traefik"},
				Default: "none",
			},
		},
		{
			Name: "roles",
			Prompt: &survey.Confirm{
				Message: "Do you want to create Ansible playbook for deploy your project?",
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
			Name: "username",
			Prompt: &survey.Input{
				Message: "Enter username:",
				Default: "root",
			},
			Validate: survey.Required,
		},
		{
			Name: "become",
			Prompt: &survey.Confirm{
				Message: "Do you need to enter password for this username?",
				Default: true,
			},
		},
		{
			Name: "host",
			Prompt: &survey.Input{
				Message: "Enter host name to deploy:",
				Default: "localhost",
			},
			Validate: survey.Required,
		},
		{
			Name: "network",
			Prompt: &survey.Input{
				Message: "Enter name of Docker network:",
				Default: "cgapp_network",
			},
			Validate: survey.Required,
		},
		{
			Name: "port",
			Prompt: &survey.Input{
				Message: "Enter port of backend Docker container:",
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
