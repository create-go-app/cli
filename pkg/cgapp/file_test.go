package cgapp

import (
	"os"
	"testing"
)

func TestFile(t *testing.T) {
	type args struct {
		name string
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"success",
			args{
				name: "test.txt",
				data: []byte("test"),
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := File(tt.args.name, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("File() error = %v, wantErr %v", err, tt.wantErr)
			}
		})

		os.Remove(tt.args.name)
	}
}
