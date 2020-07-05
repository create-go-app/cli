package cgapp

import (
	"os"
	"testing"
)

func TestCreateProjectFromRegistry(t *testing.T) {
	type args struct {
		p        *Project
		registry map[string]string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"success create backend",
			args{
				p: &Project{
					Name:       "echo",
					Type:       "backend",
					RootFolder: "../../tmp",
				},
				registry: map[string]string{
					"echo": "create-go-app/echo-go-template",
				},
			},
			false,
		},
		{
			"success create webserver",
			args{
				p: &Project{
					Name:       "nginx",
					Type:       "webserver",
					RootFolder: "../../tmp",
				},
				registry: map[string]string{
					"nginx": "create-go-app/nginx-docker",
				},
			},
			false,
		},
		{
			"success create backend from user template",
			args{
				p: &Project{
					Name:       "github.com/create-go-app/echo-go-template",
					Type:       "backend",
					RootFolder: "../../tmp",
				},
				registry: map[string]string{},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateProjectFromRegistry(tt.args.p, tt.args.registry); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}

			// Clean
			os.RemoveAll(tt.args.p.RootFolder)
		})
	}
}
