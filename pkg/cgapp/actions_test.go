package cgapp

import (
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
			"failed",
			args{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateCLIAction(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("CreateCLIAction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
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
			"failed",
			args{},
			true,
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
