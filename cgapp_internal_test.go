package main

import (
	"os"
	"testing"
)

func Test_createConfig(t *testing.T) {
	type args struct {
		e *embedConfig
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"success create config files",
			args{
				&embedConfig{
					embedFolder: "./configs",
					appFolder:   "./tmp",
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := createConfig(tt.args.e); (err != nil) != tt.wantErr {
				t.Errorf("createConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
			// Clean
			os.RemoveAll(tt.args.e.appFolder)
		})
	}
}

func Test_createApp(t *testing.T) {
	type args struct {
		c *appConfig
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"success create backend from deafault template",
			args{
				&appConfig{
					name:   "echo",
					match:  "^(echo)$",
					view:   "backend",
					folder: "./tmp",
				},
			},
			false,
		},
		{
			"success create frontend from user template",
			args{
				&appConfig{
					name:   "https://github.com/create-go-app/preact-js-template",
					match:  "^(vue)$",
					view:   "frontend",
					folder: "./tmp",
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := createApp(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("createApp() error = %v, wantErr %v", err, tt.wantErr)
			}
			// Clean
			os.RemoveAll(tt.args.c.folder)
		})
	}
}

func Test_errChecker(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"no error", args{err: nil}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errChecker(tt.args.err)
		})
	}
}
