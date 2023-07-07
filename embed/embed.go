package embed

import (
	"embed"
)

var (
	//go:embed configs/*
	ConfigsFiles embed.FS

	//go:embed misc/*
	MiscFiles embed.FS

	//go:embed roles/*
	RolesFiles embed.FS

	//go:embed templates/*
	TemplatesFiles embed.FS
)

// Files ...
type Files struct {
	Configs, Misc, Roles, Templates embed.FS
}

func New() *Files {
	return &Files{
		Configs:   ConfigsFiles,
		Misc:      MiscFiles,
		Roles:     RolesFiles,
		Templates: TemplatesFiles,
	}
}
