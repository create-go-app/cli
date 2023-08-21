package executor

import "testing"

func TestExecCommand(t *testing.T) {
	type args struct {
		command string
		options []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"successfully executing command",
			args{
				command: "echo",
				options: []string{"ping"},
			},
			false,
		},
		{
			"successfully executing command with silent mode",
			args{
				command: "echo",
				options: []string{"ping"},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Execute(tt.args.command, tt.args.options...); (err != nil) != tt.wantErr {
				t.Errorf("Execute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
