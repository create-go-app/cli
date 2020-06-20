package cgapp

import (
	"errors"
	"testing"
)

func Test_ErrChecker(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"no error",
			args{
				err: nil,
			},
		},
		{
			"error",
			args{
				err: errors.New("This is error"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ErrChecker(tt.args.err)
		})
	}
}
