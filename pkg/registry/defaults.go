// Copyright 2022 Vic Shóstak and Create Go App Contributors. All rights reserved.
// Use of this source code is governed by Apache 2.0 license
// that can be found in the LICENSE file.

package registry

import (
	"embed"

	"github.com/AlecAivazis/survey/v2"
)

// CLIVersion version of Create Go App CLI.
const CLIVersion string = "3.6.2"

// Variables struct for Ansible variables (inventory, hosts).
type Variables struct {
	List map[string]interface{}
}

// CreateAnswers struct for a survey's answers for `create` command.
type CreateAnswers struct {
	Backend       string
	Frontend      string
	Proxy         string
	AgreeCreation bool `survey:"agree"`
}

var (
	// EmbedMiscFiles misc files and configs.
	//go:embed misc/*
	EmbedMiscFiles embed.FS

	// EmbedRoles Ansible roles.
	//go:embed roles/*
	EmbedRoles embed.FS

	// EmbedTemplates template files.
	//go:embed templates/*
	EmbedTemplates embed.FS

	// CreateQuestions survey's questions for `create` command.
	CreateQuestions = []*survey.Question{
		{
			Name: "backend",
			Prompt: &survey.Select{
				Message: "Choose a backend framework:",
				Options: []string{
					"net/http",
					"fiber",
					"chi",
				},
				Default:  "fiber",
				PageSize: 3,
			},
			Validate: survey.Required,
		},
		{
			Name: "frontend",
			Prompt: &survey.Select{
				Message: "Choose a frontend framework/library:",
				Help:    "Option with a `*-ts` tail will create a TypeScript template.",
				Options: []string{
					"none",
					"vanilla",
					"vanilla-ts",
					"react",
					"react-ts",
					"preact",
					"preact-ts",
					"next",
					"next-ts",
					"nuxt3",
					"vue",
					"vue-ts",
					"svelte",
					"svelte-ts",
					"lit-element",
					"lit-element-ts",
				},
				Default:  "none",
				PageSize: 16,
			},
		},
		{
			Name: "proxy",
			Prompt: &survey.Select{
				Message: "Choose a web/proxy server:",
				Options: []string{
					"none",
					"traefik",
					"traefik-acme-dns",
					"nginx",
				},
				Default:  "none",
				PageSize: 4,
			},
		},
		{
			Name: "agree",
			Prompt: &survey.Confirm{
				Message: "If everything is okay, can I create this project for you? ;)",
				Default: true,
			},
		},
	}

	// CustomCreateQuestions survey's questions for `create -c` command.
	CustomCreateQuestions = []*survey.Question{
		{
			Name: "backend",
			Prompt: &survey.Input{
				Message: "Enter URL to the custom backend repository:",
			},
			Validate: survey.Required,
		},
		{
			Name: "frontend",
			Prompt: &survey.Input{
				Message: "Enter URL to the custom frontend repository:",
				Default: "none",
			},
		},
		{
			Name: "proxy",
			Prompt: &survey.Select{
				Message: "Choose a web/proxy server:",
				Options: []string{
					"none",
					"traefik",
					"traefik-acme-dns",
					"nginx",
				},
				Default:  "none",
				PageSize: 4,
			},
		},
		{
			Name: "agree",
			Prompt: &survey.Confirm{
				Message: "If everything is okay, can I create this project for you? ;)",
				Default: true,
			},
		},
	}

	// AnsibleInventoryVariables list of variables for inventory.
	AnsibleInventoryVariables = map[string]*Variables{
		"none": {
			List: map[string]interface{}{
				"Proxy": "none",
			},
		},
		"traefik": {
			List: map[string]interface{}{
				"Proxy":    "traefik",
				"Wildcard": false,
			},
		},
		"traefik-acme-dns": {
			List: map[string]interface{}{
				"Proxy":    "traefik",
				"Wildcard": true,
			},
		},
		"nginx": {
			List: map[string]interface{}{
				"Proxy": "nginx",
			},
		},
	}

	// AnsiblePlaybookVariables list of variables for playbook.
	AnsiblePlaybookVariables = map[string]*Variables{
		"none": {
			List: map[string]interface{}{
				"Proxy": "none",
			},
		},
		"traefik": {
			List: map[string]interface{}{
				"Proxy": "traefik",
			},
		},
		"traefik-acme-dns": {
			List: map[string]interface{}{
				"Proxy": "traefik",
			},
		},
		"nginx": {
			List: map[string]interface{}{
				"Proxy": "nginx",
			},
		},
	}
)
