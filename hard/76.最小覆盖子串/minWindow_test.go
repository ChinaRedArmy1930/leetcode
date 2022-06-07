package main

import "testing"

func Test_minWindow(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "ADOBECODEBANC-ABC",
			args: args{s: "ADOBECODEBANC", t: "ABC"},
			want: "BANC",
		},
		{
			name: "AA-AA",
			args: args{s: "AA", t: "AA"},
			want: "AA",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minWindow(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("minWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}
