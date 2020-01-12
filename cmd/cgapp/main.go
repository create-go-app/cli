package main

import (
	"github.com/create-go-app/cli/internal/cgapp"
)

var (
	// CLI version
	version string = "0.2.1"

	// Templates registry
	registry = map[string]string{
		// Backend templates
		"net/http": "create-go-app/net_http-go-template",
		"echo":     "create-go-app/echo-go-template",
		// "gin":      "create-go-app/gin-go-template",
		// "iris":     "create-go-app/iris-go-template",

		// Frontend templates
		"react":  "create-go-app/react-js-template",
		"preact": "create-go-app/preact-js-template",
		// "vue":      "create-go-app/vue-js-template",
		// "svelte":   "create-go-app/svelte-js-template",

		// Docker containers
		"nginx": "create-go-app/nginx-certbot-docker",
	}
)

func main() {
	// Start new CLI app
	cgapp.New(version, registry)
}
