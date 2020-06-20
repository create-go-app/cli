package cgapp

import "testing"

func TestNew(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			"successfully created new CLI",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			New()
		})
	}
}
