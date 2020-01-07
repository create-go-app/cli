package cgapp

import (
	"os"
	"testing"
)

func Test_CreateConfig(t *testing.T) {
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
					embedFolder: "/configs",
					appFolder:   "./tmp",
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateConfig(tt.args.e); (err != nil) != tt.wantErr {
				t.Errorf("createConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
			// Clean
			os.RemoveAll(tt.args.e.appFolder)
		})
	}
}
