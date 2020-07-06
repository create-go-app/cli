package cgapp

import (
	"os"
	"testing"
)

func TestCreateProjectFromRegistry(t *testing.T) {
	type args struct {
		project  *Project
		registry map[string]*Registry
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"successfully create default Ansible roles",
			args{
				project: &Project{
					Type:       "ansible",
					Name:       "roles",
					RootFolder: "../../tmp",
				},
				registry: registry,
			},
			false,
		},
		{
			"successfully create default backend",
			args{
				project: &Project{
					Type:       "backend",
					Name:       "echo",
					RootFolder: "../../tmp",
				},
				registry: registry,
			},
			false,
		},
		{
			"successfully create default webserver",
			args{
				project: &Project{
					Type:       "webserver",
					Name:       "nginx",
					RootFolder: "../../tmp",
				},
				registry: registry,
			},
			false,
		},
		{
			"successfully create default database",
			args{
				project: &Project{
					Type:       "database",
					Name:       "postgres",
					RootFolder: "../../tmp",
				},
				registry: registry,
			},
			false,
		},
		{
			"successfully create backend from user template",
			args{
				project: &Project{
					Type:       "backend",
					Name:       "github.com/create-go-app/echo-go-template",
					RootFolder: "../../tmp",
				},
				registry: registry,
			},
			false,
		},
		{
			"failed create (empty Project struct)",
			args{
				project:  nil,
				registry: registry,
			},
			true,
		},
		{
			"failed create (empty Registry struct)",
			args{
				project: &Project{
					Type:       "backend",
					Name:       "github.com/create-go-app/echo-go-template",
					RootFolder: "../../tmp",
				},
				registry: nil,
			},
			true,
		},
		{
			"failed create (empty Project and Registry struct)",
			args{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateProjectFromRegistry(tt.args.project, tt.args.registry); (err != nil) != tt.wantErr {
				t.Errorf("CreateProjectFromRegistry() error = %v, wantErr %v", err, tt.wantErr)
			}
		})

		// Clean
		if tt.args.project != nil {
			os.RemoveAll(tt.args.project.RootFolder)
		}
	}
}
