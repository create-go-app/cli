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
		os.RemoveAll("../../tmp")
	}
}
