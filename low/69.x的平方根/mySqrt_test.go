package main

import "testing"

func Test_mySqrt(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{1},
			want: 1,
		},
		{
			name: "2",
			args: args{2},
			want: 1,
		},
		{
			name: "3",
			args: args{3},
			want: 1,
		},
		{
			name: "4",
			args: args{4},
			want: 2,
		},
		{
			name: "5",
			args: args{5},
			want: 2,
		},
		{
			name: "6",
			args: args{6},
			want: 2,
		},
		{
			name: "7",
			args: args{7},
			want: 2,
		},
		{
			name: "8",
			args: args{8},
			want: 2,
		},
		{
			name: "9",
			args: args{9},
			want: 3,
		},
		{
			name: "10",
			args: args{10},
			want: 3,
		},
		{
			name: "11",
			args: args{11},
			want: 3,
		}, {
			name: "12",
			args: args{12},
			want: 3,
		}, {
			name: "13",
			args: args{13},
			want: 3,
		}, {
			name: "14",
			args: args{14},
			want: 3,
		}, {
			name: "15",
			args: args{15},
			want: 3,
		}, {
			name: "16",
			args: args{16},
			want: 4,
		}, {
			name: "17",
			args: args{17},
			want: 4,
		}, {
			name: "18",
			args: args{18},
			want: 4,
		}, {
			name: "19",
			args: args{19},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mySqrt(tt.args.x); got != tt.want {
				t.Errorf("mySqrt() = %v, want %v", got, tt.want)
			}
		})
	}
}
