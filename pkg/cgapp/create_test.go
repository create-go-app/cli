package cgapp

import (
	"os"
	"testing"

	"github.com/create-go-app/cli/pkg/registry"
)

func TestCreateProjectFromRegistry(t *testing.T) {
	type args struct {
		p *registry.Project
		r map[string]*registry.Repository
		m string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"successfully created backend",
			args{
				p: &registry.Project{
					Type:       "backend",
					Name:       "echo",
					RootFolder: "../../tmp",
				},
				r: registry.Repositories,
				m: registry.RegexpBackendPattern,
			},
			false,
		},
		{
			"successfully created webserver",
			args{
				p: &registry.Project{
					Type:       "webserver",
					Name:       "nginx",
					RootFolder: "../../tmp",
				},
				r: registry.Repositories,
				m: registry.RegexpWebServerPattern,
			},
			false,
		},
		{
			"successfully created database",
			args{
				p: &registry.Project{
					Type:       "database",
					Name:       "postgres",
					RootFolder: "../../tmp",
				},
				r: registry.Repositories,
				m: registry.RegexpDatabasePattern,
			},
			false,
		},
		{
			"successfully created Ansible deploy roles",
			args{
				p: &registry.Project{
					Type:       "roles",
					Name:       "deploy",
					RootFolder: "../../tmp",
				},
				r: registry.Repositories,
				m: registry.RegexpAnsiblePattern,
			},
			false,
		},
		{
			"failed to create (not valid repository)",
			args{
				p: &registry.Project{
					Type:       "roles",
					Name:       "deploy",
					RootFolder: "../../tmp",
				},
				r: map[string]*registry.Repository{
					"roles": {
						List: map[string]string{},
					},
				},
				m: registry.RegexpAnsiblePattern,
			},
			true,
		},
		{
			"failed to create (not valid repository with wrong name)",
			args{
				p: &registry.Project{
					Type:       "roles",
					Name:       "wrong-name",
					RootFolder: "../../tmp",
				},
				r: map[string]*registry.Repository{
					"roles": {
						List: map[string]string{},
					},
				},
				m: registry.RegexpAnsiblePattern,
			},
			true,
		},
		{
			"failed to create (repositories is nil)",
			args{
				p: &registry.Project{
					Type:       "roles",
					Name:       "deploy",
					RootFolder: "../../tmp",
				},
				r: nil,
				m: registry.RegexpAnsiblePattern,
			},
			true,
		},
		{
			"failed to create (project registry is nil)",
			args{
				p: nil,
				r: registry.Repositories,
				m: registry.RegexpAnsiblePattern,
			},
			true,
		},
		{
			"failed to create (pattern is nil)",
			args{
				p: &registry.Project{
					Type:       "roles",
					Name:       "deploy",
					RootFolder: "../../tmp",
				},
				r: registry.Repositories,
				m: "",
			},
			true,
		},
		{
			"failed to create (repositories, project registry & pattern are nil)",
			args{
				p: nil,
				r: nil,
				m: "",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateProjectFromRegistry(tt.args.p, tt.args.r, tt.args.m); (err != nil) != tt.wantErr {
				t.Errorf("CreateProjectFromRegistry() error = %v, wantErr %v", err, tt.wantErr)
			}
		})

		// Clean
		os.RemoveAll("../../tmp")
	}
}

func TestCreateProjectFromCmd(t *testing.T) {
	type args struct {
		p *registry.Project
		c map[string]*registry.Command
		m string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"successfully create react:redux",
			args{
				p: &registry.Project{
					Type:       "frontend",
					Name:       "react:redux",
					RootFolder: "../../tmp",
				},
				c: map[string]*registry.Command{
					"react": {
						Runner: "echo",
						Create: "react",
						Args:   map[string]string{},
					},
				},
				m: registry.RegexpFrontendPattern,
			},
			false,
		},
		{
			"successfully create preact:simple",
			args{
				p: &registry.Project{
					Type:       "frontend",
					Name:       "preact:simple",
					RootFolder: "../../tmp",
				},
				c: map[string]*registry.Command{
					"preact": {
						Runner: "echo",
						Create: "preact",
						Args:   map[string]string{},
					},
				},
				m: registry.RegexpFrontendPattern,
			},
			false,
		},
		{
			"successfully create vue (with mock for GitHub)",
			args{
				p: &registry.Project{
					Type:       "frontend",
					Name:       "vue:mock",
					RootFolder: "../../tmp",
				},
				c: map[string]*registry.Command{
					"vue": {
						Runner: "echo",
						Create: "vue",
						Args:   map[string]string{},
					},
				},
				m: registry.RegexpFrontendPattern,
			},
			false,
		},
		{
			"successfully create vue (with mock for others)",
			args{
				p: &registry.Project{
					Type:       "frontend",
					Name:       "vue:mock:mock",
					RootFolder: "../../tmp",
				},
				c: map[string]*registry.Command{
					"vue": {
						Runner: "echo",
						Create: "vue",
						Args:   map[string]string{},
					},
				},
				m: registry.RegexpFrontendPattern,
			},
			false,
		},
		{
			"successfully create angular",
			args{
				p: &registry.Project{
					Type:       "frontend",
					Name:       "angular",
					RootFolder: "../../tmp",
				},
				c: map[string]*registry.Command{
					"angular": {
						Runner: "echo",
						Create: "angular",
						Args:   map[string]string{},
					},
				},
				m: registry.RegexpFrontendPattern,
			},
			false,
		},
		{
			"successfully create svelte",
			args{
				p: &registry.Project{
					Type:       "frontend",
					Name:       "svelte",
					RootFolder: "../../tmp",
				},
				c: map[string]*registry.Command{
					"svelte": {
						Runner: "echo",
						Create: "svelte",
						Args:   map[string]string{},
					},
				},
				m: registry.RegexpFrontendPattern,
			},
			false,
		},
		{
			"successfully create from user repository",
			args{
				p: &registry.Project{
					Type:       "frontend",
					Name:       "github.com/create-go-app/preact-js-template",
					RootFolder: "../../tmp",
				},
				c: registry.Commands,
				m: registry.RegexpFrontendPattern,
			},
			false,
		},
		{
			"failed to create from user repository",
			args{
				p: &registry.Project{
					Type:       "frontend",
					Name:       "unknown.com/user/repo",
					RootFolder: "../../tmp",
				},
				c: registry.Commands,
				m: registry.RegexpFrontendPattern,
			},
			true,
		},
		{
			"failed to create (project registry & commands are nil)",
			args{
				p: nil,
				c: nil,
				m: "",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateProjectFromCmd(tt.args.p, tt.args.c, tt.args.m); (err != nil) != tt.wantErr {
				t.Errorf("CreateProjectFromCmd() error = %v, wantErr %v", err, tt.wantErr)
			}
		})

		// Clean
		os.RemoveAll("../../tmp")
	}
}
