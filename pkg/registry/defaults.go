/*
Package registry includes available repositories & commands for the Create Go App CLI.

Create a new production-ready project with backend (Golang),
frontend (JavaScript, TypeScript) and deploy automation
(Ansible, Docker) by running one CLI command.

-> Focus on writing code and thinking of business logic!
<- The Create Go App CLI will take care of the rest.

A helpful documentation and next steps -> https://create-go.app/

Copyright © 2019-present Vic Shóstak <truewebartisans@gmail.com> (https://1wa.co)

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
package registry

import "github.com/AlecAivazis/survey/v2"

const (
	// CLIVersion version of Create Go App CLI.
	CLIVersion = "1.5.1"
	// RegexpAnsiblePattern pattern for Ansible roles.
	RegexpAnsiblePattern = "^(deploy)$"
	// RegexpBackendPattern pattern for backend.
	RegexpBackendPattern = "^(net/http|fiber|echo)$"
	// RegexpFrontendPattern pattern for backend.
	RegexpFrontendPattern = "^(p?react:?|vue(:?[\\w]+)?(:?[\\w-_0-9\\/]+)?|angular|svelte|sapper:?)"
	// RegexpWebServerPattern pattern for web/proxy servers.
	RegexpWebServerPattern = "^(nginx)$"
	// RegexpDatabasePattern pattern for databases.
	RegexpDatabasePattern = "^(postgres)$"
)

// Project struct for describe project.
type Project struct {
	Type       string
	Name       string
	RootFolder string
}

// Repository ...
type Repository struct {
	List map[string]string
}

// Command ...
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
	Database            string
	InstallAnsibleRoles bool `survey:"roles"`
	AgreeCreation       bool `survey:"agree"`
}

// DeployAnswers struct for a survey's answers for `deploy` command.
type DeployAnswers struct {
	Username        string
	Host            string
	Network         string
	AskBecomePass   bool `survey:"become"`
	AgreeDeployment bool `survey:"agree"`
}

var (
	// Repositories collection.
	Repositories = map[string]*Repository{
		// Ansible roles.
		"roles": {
			List: map[string]string{
				"deploy": "github.com/create-go-app/ansible-roles-deploy",
			},
		},

		// Backend templates.
		"backend": {
			List: map[string]string{
				"net/http": "github.com/create-go-app/net_http-go-template",
				"fiber":    "github.com/create-go-app/fiber-go-template",
				"echo":     "github.com/create-go-app/echo-go-template",
			},
		},

		// Docker containers with web/proxy servers.
		"webserver": {
			List: map[string]string{
				"nginx": "github.com/create-go-app/nginx-docker",
			},
		},

		// Docker containers with databases.
		"database": {
			List: map[string]string{
				"postgres": "github.com/create-go-app/postgres-docker",
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
				Options: []string{"net/http", "Fiber", "Echo", "Gin"},
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
				Options: []string{"none", "Nginx"},
				Default: "none",
			},
		},
		{
			Name: "database",
			Prompt: &survey.Select{
				Message: "Choose a database:",
				Options: []string{"none", "Postgres"},
				Default: "none",
			},
		},
		{
			Name: "roles",
			Prompt: &survey.Confirm{
				Message: "Do you want to install Ansible roles for deploy your project?",
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
			Name: "agree",
			Prompt: &survey.Confirm{
				Message: "If everything is okay, can I deploy this project? ;)",
				Default: true,
			},
		},
	}
)
