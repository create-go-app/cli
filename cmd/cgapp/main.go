package main

import "github.com/create-go-app/cli/internal/cgapp"

var (
	// cgapp CLI version
	version string = "0.6.0b3"

	// Templates registry
	registry = map[string]string{
		// Backend templates
		"net/http": "create-go-app/net_http-go-template",
		"echo":     "create-go-app/echo-go-template",
		"fiber":    "create-go-app/fiber-go-template",

		// Frontend templates
		"react":  "create-go-app/react-js-template",
		"preact": "create-go-app/preact-js-template",

		// Docker containers
		"nginx":    "create-go-app/nginx-certbot-docker",
		"postgres": "create-go-app/postgres-docker",
	}
)

func main() {
	// Start new CLI app
	cgapp.New(version, registry)
}
