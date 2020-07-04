package cgapp

var (
	// cgapp CLI version
	version = "1.2.1"

	// Templates registry
	registry = map[string]string{
		// Ansible roles
		"roles": "create-go-app/ansible-roles",

		// Backend templates
		"net/http": "create-go-app/net_http-go-template",
		"fiber":    "create-go-app/fiber-go-template",
		"echo":     "create-go-app/echo-go-template",

		// Frontend templates
		"react-js": "create-go-app/react-js-template",
		"preact":   "create-go-app/preact-js-template",
		"react-ts": "create-go-app/react-ts-template",

		// Docker containers
		"nginx":    "create-go-app/nginx-docker",
		"postgres": "create-go-app/postgres-docker",
	}
)
