package cgapp

const (
	// Create Go App CLI version
	version = "1.3.0"

	// Regexp patterns
	regexpAnsiblePattern   = "^(roles)$"
	regexpBackendPattern   = "^(net/http|fiber|echo)$"
	regexpWebServerPattern = "^(nginx)$"
	regexpDatabasePattern  = "^(postgres)$"
)

var (
	// Templates registry
	registry = map[string]string{
		// Ansible roles
		"roles": "github.com/create-go-app/ansible-roles",

		// Backend templates
		"net/http": "github.com/create-go-app/net_http-go-template",
		"fiber":    "github.com/create-go-app/fiber-go-template",
		"echo":     "github.com/create-go-app/echo-go-template",

		// Docker containers
		"nginx":    "github.com/create-go-app/nginx-docker",
		"postgres": "github.com/create-go-app/postgres-docker",
	}

	// CMD commands collection
	commands = map[string]map[string]string{
		"react": {
			"runner":   "npx",
			"create":   "create-react-app",
			"template": "--template",
		},
	}
)
