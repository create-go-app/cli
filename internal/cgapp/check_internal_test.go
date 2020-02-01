package cgapp

import "testing"

func Test_ErrChecker(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			"no error",
			args{
				err: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ErrChecker(tt.args.err)
		})
	}
}
