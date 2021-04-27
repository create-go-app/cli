package cgapp

import (
	"os"
	"testing"
)

func TestCreateProjectFromCmd(t *testing.T) {
	type args struct {
		options []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"successfully create React template",
			args{
				options: []string{
					"init", "@vitejs/app", "../../tmp/frontend", "--", "--template", "react",
				},
			},
			false,
		},
		{
			"failed to create",
			args{
				options: []string{},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateProjectFromCmd(tt.args.options); (err != nil) != tt.wantErr {
				t.Errorf("CreateProjectFromCmd() error = %v, wantErr %v", err, tt.wantErr)
			}
		})

		// Clean.
		os.RemoveAll("../../tmp")
	}
}

func TestCreateProjectFromGit(t *testing.T) {
	type args struct {
		projectType       string
		projectRepository string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"successfully create template",
			args{
				projectType:       "../../tmp/test",
				projectRepository: "github.com/koddr/koddr",
			},
			false,
		},
		{
			"failed to create",
			args{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateProjectFromGit(tt.args.projectType, tt.args.projectRepository); (err != nil) != tt.wantErr {
				t.Errorf("CreateProjectFromGit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})

		// Clean.
		os.RemoveAll("../../tmp")
	}
}
