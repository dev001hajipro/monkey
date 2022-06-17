package parser

import "testing"

func Test_untrace(t *testing.T) {
	type args struct {
		msg string
	}
	tests := []struct {
		name string
		args args
	}{
		{"xuntrace", args{"message"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			untrace(trace(tt.args.msg))
		})
	}
}
