package cgapp

import (
	"os"
	"testing"

	"github.com/urfave/cli/v2"
)

func TestCreateCLIAction(t *testing.T) {
	type args struct {
		c *cli.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"successfully executing create action",
			args{},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateCLIAction(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("CreateCLIAction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

	files := []string{".editorconfig", ".gitignore", "deploy-playbook.yml", "Taskfile.yml", "roles"}
	for _, name := range files {
		os.RemoveAll(name)
	}
}

func TestDeployCLIAction(t *testing.T) {
	type args struct {
		c *cli.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"successfully executing deploy action",
			args{},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DeployCLIAction(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("DeployCLIAction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
