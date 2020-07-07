package utils

import (
	"os"
	"reflect"
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
			"successfully executing command",
			args{
				command: "echo",
				options: []string{"ping"},
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
			"successfully created folder",
			args{
				folderName: "../../tmp",
				chmod:      0750,
			},
			false,
		},
		{
			"failed, folder is exists",
			args{
				folderName: "",
				chmod:      0750,
			},
			true,
		},
		{
			"failed, folder is exists",
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
			"successfully created files",
			args{
				rootFolder: "../../tmp",
				filesToMake: map[string][]byte{
					"test.txt": []byte("test"),
				},
			},
			false,
		},
		{
			"failed created files",
			args{
				rootFolder: "./does/not-exists",
				filesToMake: map[string][]byte{
					"test.txt": []byte("test"),
				},
			},
			true,
		},
	}

	_ = os.Mkdir("../../tmp", 0750)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := MakeFiles(tt.args.rootFolder, tt.args.filesToMake); (err != nil) != tt.wantErr {
				t.Errorf("MakeFiles() error = %v, wantErr %v", err, tt.wantErr)
			}
		})

		// Clean
		os.RemoveAll("../../tmp")
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
			"successfully send message without args",
			args{},
		},
		{
			"successfully send message with args",
			args{true, "!", "Test", "", true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SendMsg(tt.args.startWithNewLine, tt.args.caption, tt.args.text, tt.args.color, tt.args.endWithNewLine)
		})
	}
}

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
				rootFolder:   "../../tmp",
				templateName: "github.com/create-go-app/postgres-docker",
			},
			false,
		},
		{
			"failed clone project (empty template)",
			args{
				rootFolder:   "../../tmp",
				templateName: "",
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

func TestRemoveFolders(t *testing.T) {
	type args struct {
		rootFolder      string
		foldersToRemove []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"successfully removed",
			args{
				rootFolder:      "../../tmp",
				foldersToRemove: []string{"folder-1"},
			},
			false,
		},
	}

	_ = os.MkdirAll("../../tmp/folder-1", 0750)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := RemoveFolders(tt.args.rootFolder, tt.args.foldersToRemove); (err != nil) != tt.wantErr {
				t.Errorf("RemoveFolders() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStringSplit(t *testing.T) {
	type args struct {
		pattern string
		match   string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			"successfully matched",
			args{
				pattern: ":",
				match:   "react:redux",
			},
			[]string{"react", "redux"},
			false,
		},
		{
			"successfully not matched",
			args{
				pattern: "=",
				match:   "react:redux",
			},
			[]string{"react:redux"},
			false,
		},
		{
			"failed wrong pattern and match",
			args{},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StringSplit(tt.args.pattern, tt.args.match)
			if (err != nil) != tt.wantErr {
				t.Errorf("StringSplit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringSplit() = %v, want %v", got, tt.want)
			}
		})
	}
}
