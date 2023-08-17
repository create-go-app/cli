package app

import (
	"errors"

	"github.com/create-go-app/cli/v5/internal/config"
	"github.com/create-go-app/cli/v5/internal/executor"
)

// Create ...
func (app *App) Create() error {
	// Check, if a backend part is setting up in the config.
	if app.Config.Backend != nil {
		// Create a backend part.
		switch app.Config.Backend.Name {
		case "fiber":
			// Create a backend part with Fiber framework by the given template.
			// See: https://github.com/create-go-app/fiber-go-template
			if err := executor.Execute(
				"git", "clone", config.FiberGoTemplateURL, "backend",
			); err != nil {
				return err
			}
		case "net/http":
			// Create a backend part with built-in net/http package by the given template.
			// See: https://github.com/create-go-app/net_http-go-template
			if err := executor.Execute(
				"git", "clone", config.NetHttpGoTemplateURL, "backend",
			); err != nil {
				return err
			}
		case "go-chi":
			// Create a backend part with Chi router by the given template.
			// See: https://github.com/create-go-app/chi-go-template
			if err := executor.Execute(
				"git", "clone", config.ChiGoTemplateURL, "backend",
			); err != nil {
				return err
			}
		default:
			// Check, if config has a repository part.
			if app.Config.Backend.Repository != nil {
				// Create custom backend by the given repository link.
				if err := executor.Execute(
					"git", "clone", app.Config.Backend.Repository.URL, "backend",
				); err != nil {
					return err
				}
			} else {
				return errors.New("not set 'backend' config, but is present in the configuration file")
			}
		}
	}

	// Check, if a frontend part is setting up in the config.
	if app.Config.Frontend != nil {
		// Create a frontend part.
		switch app.Config.Frontend.Name {
		case "vanilla", "vanilla-ts",
			"vue", "vue-ts",
			"react", "react-ts",
			"react-swc", "react-swc-ts",
			"preact", "preact-ts",
			"lit", "lit-ts",
			"svelte", "svelte-ts",
			"solid", "solid-ts",
			"qwik", "qwik-ts":
			// Create a frontend part with Vite.js by the given template.
			// See: https://vitejs.dev/guide/#scaffolding-your-first-vite-project
			return executor.Execute(
				"npm", "create", "vite@latest", "frontend", "--", "--template",
				app.Config.Frontend.Name,
			)
		case "next":
			// Create a frontend part with Next.js.
			// See: https://nextjs.org/docs/getting-started/installation
			return executor.Execute(
				"npx", "create-next-app@latest", "frontend",
				"--javascript", "--eslint", "--app",
				"--tailwind", "false", "--src-dir", "false", "--import-alias", "false",
			)
		case "next-tailwind":
			// Create a frontend part with Next.js and Tailwind CSS.
			// See: https://nextjs.org/docs/getting-started/installation
			return executor.Execute(
				"npx", "create-next-app@latest", "frontend",
				"--javascript", "--tailwind", "--eslint", "--app",
				"--src-dir", "false", "--import-alias", "false",
			)
		case "next-ts":
			// Create a frontend part with Next.js with Typescript.
			// See: https://nextjs.org/docs/getting-started/installation
			return executor.Execute(
				"npx", "create-next-app@latest", "frontend",
				"--typescript", "--eslint", "--app",
				"--tailwind", "false", "--src-dir", "false", "--import-alias", "false",
			)
		case "next-tailwind-ts":
			// Create a frontend part with Next.js with Typescript and Tailwind CSS.
			// See: https://nextjs.org/docs/getting-started/installation
			return executor.Execute(
				"npx", "create-next-app@latest", "frontend",
				"--typescript", "--tailwind", "--eslint", "--app",
				"--src-dir", "false", "--import-alias", "false",
			)
		case "nuxt":
			// Create a frontend part with Nuxt.js.
			// See: https://nuxt.com/docs/getting-started/installation
			return executor.Execute("npx", "nuxi@latest", "init", "frontend")
		default:
			// Check, if config has a repository part.
			if app.Config.Backend.Repository != nil {
				// Create custom frontend by the given repository link.
				if err := executor.Execute(
					"git", "clone", app.Config.Frontend.Repository.URL, "frontend",
				); err != nil {
					return err
				}
			} else {
				return errors.New("not set 'frontend' config, but is present in the configuration file")
			}
		}
	}

	return nil
}
