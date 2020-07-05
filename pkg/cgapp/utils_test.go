package cgapp

import (
	"os"
	"testing"
)

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
			"success",
			args{
				command: "echo",
				options: []string{"ping"},
			},
			false,
		},
		{
			"fail",
			args{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ExecCommand(tt.args.command, tt.args.options); (err != nil) != tt.wantErr {
				t.Errorf("ExecCommand() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBeautifyText(t *testing.T) {
	type args struct {
		text  string
		color string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"successfully send message",
			args{
				text:  "Hello World!",
				color: "",
			},
		},
		{
			"successfully send colored message",
			args{
				text:  "Hello World!",
				color: "green",
			},
		},
		{
			"successfully send colored message",
			args{
				text:  "Hello World!",
				color: "yellow",
			},
		},
		{
			"successfully send colored message",
			args{
				text:  "Hello World!",
				color: "cyan",
			},
		},
		{
			"successfully send colored message",
			args{
				text:  "Hello World!",
				color: "red",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			BeautifyText(tt.args.text, tt.args.color)
		})
	}
}

func TestMakeFolder(t *testing.T) {
	type args struct {
		folderName string
		chmod      os.FileMode
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"success",
			args{
				folderName: "../../tmp",
				chmod:      0750,
			},
			false,
		},
		{
			"fail, folder is exists",
			args{
				folderName: "",
				chmod:      0750,
			},
			true,
		},
		{
			"fail, folder is exists",
			args{
				folderName: "cgapp-project",
				chmod:      0750,
			},
			true,
		},
	}

	_ = os.Mkdir("cgapp-project", 0750)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := MakeFolder(tt.args.folderName, tt.args.chmod); (err != nil) != tt.wantErr {
				t.Errorf("MakeFolder() error = %v, wantErr %v", err, tt.wantErr)
			}
		})

		// Clean
		os.RemoveAll(tt.args.folderName)
	}
}

func TestMakeFiles(t *testing.T) {
	type args struct {
		rootFolder  string
		filesToMake map[string][]byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"success",
			args{
				rootFolder: "./",
				filesToMake: map[string][]byte{
					"test.txt": []byte("test"),
				},
			},
			false,
		},
		{
			"fail",
			args{
				rootFolder: "./does/not-exists",
				filesToMake: map[string][]byte{
					"test.txt": []byte("test"),
				},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := MakeFiles(tt.args.rootFolder, tt.args.filesToMake); (err != nil) != tt.wantErr {
				t.Errorf("MakeFiles() error = %v, wantErr %v", err, tt.wantErr)
			}
		})

		// Clean
		os.RemoveAll("test.txt")
	}
}

func TestSendMsg(t *testing.T) {
	type args struct {
		startWithNewLine bool
		caption          string
		text             string
		color            string
		endWithNewLine   bool
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"success without args",
			args{},
		},
		{
			"success with args",
			args{true, "!", "Test", "", true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SendMsg(tt.args.startWithNewLine, tt.args.caption, tt.args.text, tt.args.color, tt.args.endWithNewLine)
		})
	}
}
