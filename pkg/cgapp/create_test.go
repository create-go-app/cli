package cgapp

import (
	"os"
	"testing"
)

func TestCreate(t *testing.T) {
	type args struct {
		c        *Config
		registry map[string]string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"success create backend",
			args{
				&Config{
					Name:   "echo",
					Match:  "^(echo)$",
					View:   "backend",
					Folder: "./tmp",
				},
				map[string]string{
					"echo": "create-go-app/echo-go-template",
				},
			},
			false,
		},
		{
			"success create frontend",
			args{
				&Config{
					Name:   "github.com/create-go-app/react-js-template",
					Match:  "^(react)$",
					View:   "frontend",
					Folder: "./tmp",
				},
				map[string]string{
					"react": "create-go-app/react-js-template",
				},
			},
			false,
		},
		{
			"success create frontend from user template",
			args{
				&Config{
					Name:   "github.com/create-go-app/echo-go-template",
					Match:  "^(echo)$",
					View:   "backend",
					Folder: "./tmp",
				},
				map[string]string{
					"echo": "create-go-app/echo-go-template",
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Create(tt.args.c, tt.args.registry); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}

			// Clean
			os.RemoveAll(tt.args.c.Folder)
		})
	}
}
