package main

import "testing"

func Test_isPerfectSquare(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{1},
			want: true,
		},
		{
			name: "2",
			args: args{2},
			want: false,
		},
		{
			name: "3",
			args: args{3},
			want: false,
		},
		{
			name: "4",
			args: args{4},
			want: true,
		},
		{
			name: "5",
			args: args{5},
			want: false,
		},
		{
			name: "6",
			args: args{6},
			want: false,
		},
		{
			name: "7",
			args: args{7},
			want: false,
		},
		{
			name: "8",
			args: args{8},
			want: false,
		},
		{
			name: "9",
			args: args{9},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPerfectSquare(tt.args.num); got != tt.want {
				t.Errorf("isPerfectSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}
