package cgapp

import (
	"os"
	"reflect"
	"testing"
)

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
			_ = BeautifyText(tt.args.text, tt.args.color)
		})
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

func Test_stringSplit(t *testing.T) {
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
			got, err := stringSplit(tt.args.pattern, tt.args.match)
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
