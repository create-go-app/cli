package cmd

import (
	"testing"
)

func Test_initConfig(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			"successfully init",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			initConfig()
		})
	}
}

func TestExecute(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			"successfully",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Execute()
		})
	}
}
