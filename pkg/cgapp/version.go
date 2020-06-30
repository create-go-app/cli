package cgapp

var (
	// cgapp CLI version
	version = "1.1.1"

	// Templates registry
	registry = map[string]string{
		// Backend templates
		"net/http": "create-go-app/net_http-go-template",
		"fiber":    "create-go-app/fiber-go-template",

		// Frontend templates
		"react-js": "create-go-app/react-js-template",
		"react-ts": "create-go-app/react-ts-template",
		"preact":   "create-go-app/preact-js-template",

		// Docker containers
		"nginx":    "create-go-app/nginx-docker",
		"postgres": "create-go-app/postgres-docker",
	}
)
