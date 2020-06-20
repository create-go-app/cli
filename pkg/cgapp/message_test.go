package cgapp

import "testing"

func TestSendMessage(t *testing.T) {
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
			SendMessage(tt.args.text, tt.args.color)
		})
	}
}
