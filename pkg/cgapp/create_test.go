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
			},
			true,
		},
		{
			"failed to create (project registry is nil)",
			args{
				p: nil,
				r: registry.Repositories,
			},
			true,
		},
		{
			"failed to create (project registry & repositories are nil)",
			args{
				p: nil,
				r: nil,
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateProjectFromRegistry(tt.args.p, tt.args.r); (err != nil) != tt.wantErr {
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
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"successfully create svelte",
			args{
				p: &registry.Project{
					Type:       "frontend",
					Name:       "svelte",
					RootFolder: "../../tmp",
				},
				c: registry.Commands,
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
			},
			false,
		},
		{
			"failed to create (project registry & commands are nil)",
			args{
				p: nil,
				c: nil,
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateProjectFromCmd(tt.args.p, tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("CreateProjectFromCmd() error = %v, wantErr %v", err, tt.wantErr)
			}
		})

		// Clean
		os.RemoveAll("../../tmp")
	}
}
