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
	// CLI version
	version = "1.3.0"

	// Regexp patterns
	regexpAnsiblePattern   = "^(deploy)$"
	regexpBackendPattern   = "^(net/http|fiber|echo)$"
	regexpWebServerPattern = "^(nginx)$"
	regexpDatabasePattern  = "^(postgres)$"
)

// Registry ...
type Registry struct {
	Repositories map[string]string
}

// Command ...
type Command struct {
	Runner string
	Create string
	Args   map[string]string
}

var (
	// Repositories collection
	repositories = map[string]*Registry{
		// Ansible roles
		"roles": {
			Repositories: map[string]string{
				"deploy": "github.com/create-go-app/ansible-roles-deploy",
			},
		},

		// Backend templates
		"backend": {
			Repositories: map[string]string{
				"net/http": "github.com/create-go-app/net_http-go-template",
				"fiber":    "github.com/create-go-app/fiber-go-template",
				"echo":     "github.com/create-go-app/echo-go-template",
			},
		},

		// Docker containers with web/proxy servers
		"webserver": {
			Repositories: map[string]string{
				"nginx": "github.com/create-go-app/nginx-docker",
			},
		},

		// Docker containers with databases
		"database": {
			Repositories: map[string]string{
				"postgres": "github.com/create-go-app/postgres-docker",
			},
		},
	}

	// Commands collection
	commands = map[string]*Command{
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
