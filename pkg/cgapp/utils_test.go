package cgapp

import (
	"testing"
	"time"
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

func TestShowMessage(t *testing.T) {
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
			ShowMessage(tt.args.level, tt.args.text, tt.args.startWithNewLine, tt.args.endWithNewLine)
		})
	}
}

func TestCalculateDurationTime(t *testing.T) {
	type args struct {
		startTimer time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"successfully",
			args{
				startTimer: time.Now(),
			},
			"0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateDurationTime(tt.args.startTimer); got != tt.want {
				t.Errorf("CalculateDurationTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
