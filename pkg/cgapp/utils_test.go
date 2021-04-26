package cgapp

import (
	"reflect"
	"testing"
)

func Test_colorizeLevel(t *testing.T) {
	type args struct {
		level string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"successfully send message",
			args{level: ""},
		},
		{
			"successfully send success message",
			args{level: "success"},
		},
		{
			"successfully send warning message",
			args{level: "warning"},
		},
		{
			"successfully send error message",
			args{level: "error"},
		},
		{
			"successfully send info message",
			args{level: "info"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = colorizeLevel(tt.args.level)
		})
	}
}

func TestSendMsg(t *testing.T) {
	type args struct {
		level            string
		text             string
		startWithNewLine bool
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
			args{"success", "Test", true, true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = ShowMessage(tt.args.level, tt.args.text, tt.args.startWithNewLine, tt.args.endWithNewLine)
		})
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
