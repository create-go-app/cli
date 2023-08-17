package config

const (
	FiberGoTemplateURL   = "https://github.com/create-go-app/fiber-go-template.git"
	NetHttpGoTemplateURL = "https://github.com/create-go-app/net_http-go-template.git"
	ChiGoTemplateURL     = "https://github.com/create-go-app/chi-go-template.git"
)

var (
	// BackendTemplates ...
	BackendTemplates = []string{
		"fiber",
		"go-chi",
		"net/http",
	}

	// FrontendTemplates ...
	FrontendTemplates = []string{
		"vanilla", "vanilla-ts",
		"vue", "vue-ts",
		"react", "react-ts",
		"react-swc", "react-swc-ts",
		"preact", "preact-ts",
		"lit", "lit-ts",
		"svelte", "svelte-ts",
		"solid", "solid-ts",
		"qwik", "qwik-ts",
		"next", "nuxt",
	}

	// ProxyTemplates ...
	ProxyTemplates = []string{
		"nginx",
		"traefik",
	}
)
