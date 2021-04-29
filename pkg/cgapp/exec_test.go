package cgapp

import "testing"

func TestExecCommand(t *testing.T) {
	type args struct {
		command    string
		options    []string
		silentMode bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"successfully executing command",
			args{
				command:    "echo",
				options:    []string{"ping"},
				silentMode: false,
			},
			false,
		},
		{
			"successfully executing command with silent mode",
			args{
				command:    "echo",
				options:    []string{"ping"},
				silentMode: true,
			},
			false,
		},
		{
			"failed executing command",
			args{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ExecCommand(tt.args.command, tt.args.options, tt.args.silentMode); (err != nil) != tt.wantErr {
				t.Errorf("ExecCommand() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
