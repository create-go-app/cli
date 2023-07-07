package config

import (
	"github.com/koddr/gosl"

	"github.com/create-go-app/cli/v4/internal/config/containers"
	"github.com/create-go-app/cli/v4/internal/config/tools"
)

// Config ...
type Config struct {
	Project    *Project    `koanf:"project"`
	Backend    *Template   `koanf:"backend"`
	Frontend   *Template   `koanf:"frontend"`
	Proxy      *Template   `koanf:"proxy"`
	Containers *Containers `koanf:"containers"`
	Deploy     *Deploy     `koanf:"deploy"`
}

// Project ...
type Project struct {
	Name string `koanf:"name"`
}

// Template ...
type Template struct {
	Name       string      `koanf:"template"`
	Repository *Repository `koanf:"repository"`
}

// Repository ...
type Repository struct {
	URL     string   `koanf:"url"`
	Private *Private `koanf:"private"`
}

// Private ...
type Private struct {
	Key string `koanf:"key"`
}

// Containers ...
type Containers struct {
	Golang   *containers.Golang   `koanf:"golang"`
	Postgres *containers.Postgres `koanf:"postgres"`
	Redis    *containers.Redis    `koanf:"redis"`
	NodeJS   *containers.NodeJS   `koanf:"nodejs"`
	Nginx    *containers.Nginx    `koanf:"nginx"`
	Traefik  *containers.Traefik  `koanf:"traefik"`
}

// Deploy ...
type Deploy struct {
	Ansible     *tools.Ansible     `koanf:"ansible"`
	Terraform   *tools.Terraform   `koanf:"terraform"`
	Docker      *tools.Docker      `koanf:"docker"`
	LetsEncrypt *tools.LetsEncrypt `koanf:"lets_encrypt"`
}

// New ...
func New(path string) (*Config, error) {
	// Parse the given config to a struct.
	cfg, err := gosl.ParseFileWithEnvToStruct(path, "CGAPP", &Config{})
	if err != nil {
		return nil, err
	}

	// Validate config.
	if err = cfg.Validate(); err != nil {
		return nil, err
	}

	return cfg, nil
}
