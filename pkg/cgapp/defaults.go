package cgapp

const (
	// Regexp patterns
	regexpAnsiblePattern   = "^(roles)$"
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
	// Registry
	registry = map[string]*Registry{
		// Ansible roles
		"ansible": {
			Repositories: map[string]string{
				"roles": "github.com/create-go-app/ansible-roles",
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

	// CMD commands collection
	cmds = map[string]*Command{
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
				"cwd":          "--cwd",
				"name":         "--name",
				"skip-git":     "--git",
				"skip-install": "--install",
			},
		},
	}
)
