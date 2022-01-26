package cgapp

import (
	"os"
	"testing"
)

func TestGitClone(t *testing.T) {
	type args struct {
		rootFolder   string
		templateName string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"successfully cloned project",
			args{
				rootFolder:   "../../tmp/test",
				templateName: "github.com/koddr/koddr",
			},
			false,
		},
		{
			"failed clone project (empty template)",
			args{
				rootFolder:   "../../tmp/test",
				templateName: "",
			},
			true,
		},
		{
			"failed clone project",
			args{
				rootFolder:   "../../tmp/test",
				templateName: "404.404/404/404",
			},
			true,
		},
		{
			"failed clone project (empty args)",
			args{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GitClone(tt.args.rootFolder, tt.args.templateName); (err != nil) != tt.wantErr {
				t.Errorf("GitClone() error = %v, wantErr %v", err, tt.wantErr)
			}
		})

		// Clean
		err := os.RemoveAll("../../tmp")
		if err != nil {
			return
		}
	}
}

func Test_getAbsoluteURL(t *testing.T) {
	type args struct {
		templateURL   string
	}
	tests := []struct {
		name    string
		args    args
		want string
	}{
		{
			"successfully get absolute url from url with scheme",
			args{
				templateURL: "https://github.com/create-go-app/net_http-go-template",
			},
			"https://github.com/create-go-app/net_http-go-template",
		},
		{
			"successfully get absolute url from url without scheme",
			args{
				templateURL: "github.com/create-go-app/net_http-go-template",
			},
			"https://github.com/create-go-app/net_http-go-template",
		},
		{
			"successfully get absolute url from url starting space",
			args{
				templateURL: " github.com/create-go-app/net_http-go-template",
			},
			"https://github.com/create-go-app/net_http-go-template",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getAbsoluteURL(tt.args.templateURL); got != tt.want {
				t.Errorf("getAbsoluteURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
