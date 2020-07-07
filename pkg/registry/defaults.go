/*
Package registry includes available repositories & commands for the Create Go App CLI.

Copyright © 2020 Vic Shóstak <truewebartisans@gmail.com> (https://1wa.co)

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

const (
	// CLIVersion version of Create Go App CLI.
	CLIVersion = "1.3.0"
	// RegexpAnsiblePattern pattern for Ansible roles.
	RegexpAnsiblePattern = "^(deploy)$"
	// RegexpBackendPattern pattern for backend.
	RegexpBackendPattern = "^(net/http|fiber|echo)$"
	// RegexpWebServerPattern pattern for web/proxy servers.
	RegexpWebServerPattern = "^(nginx)$"
	// RegexpDatabasePattern pattern for databases.
	RegexpDatabasePattern = "^(postgres)$"
)

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
		"svelte": {
			Runner: "npx",
			Create: "degit",
			Args: map[string]string{
				"template": "sveltejs/template",
			},
		},
	}
)
