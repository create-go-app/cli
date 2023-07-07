package config

var (
	// BackendTemplates ...
	BackendTemplates = []string{
		"fiber",
		"go-chi",
		"net/http",
	}

	// FrontendTemplates ...
	FrontendTemplates = []string{
		"vanilla",
		"vanilla-ts",
		"react",
		"react-ts",
		"preact",
		"preact-ts",
		"next",
		"next-ts",
		"nuxt",
		"vue",
		"vue-ts",
		"svelte",
		"svelte-ts",
		"lit",
		"lit-ts",
	}

	// ProxyTemplates ...
	ProxyTemplates = []string{
		"nginx",
		"traefik",
	}
)
