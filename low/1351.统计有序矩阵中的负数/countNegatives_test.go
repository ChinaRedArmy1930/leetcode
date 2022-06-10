package main

import "testing"

func Test_countOneLine(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "4,3,2,-1",
			args: args{[]int{4, 3, 2, -1}},
			want: 1,
		},
		{
			name: "3,2,1,-1",
			args: args{[]int{3, 2, 1, -1}},
			want: 1,
		},
		{
			name: "1,1,-1,-2",
			args: args{[]int{1, 1, -1, -2}},
			want: 2,
		},
		{
			name: "-1,-1,-2,-3",
			args: args{[]int{-1, -1, -2, -3}},
			want: 4,
		},
		{
			name: "1, 1, 2, 3",
			args: args{[]int{1, 1, 2, 3}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countOneLine(tt.args.nums); got != tt.want {
				t.Errorf("countOneLine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countNegatives(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[[4,3,2,-1],[3,2,1,-1],[1,1,-1,-2],[-1,-1,-2,-3]]",
			args: args{[][]int{{4, 3, 2, -1}, {3, 2, 1, -1}, {1, 1, -1, -2}, {-1, -1, -2, -3}}},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countNegatives(tt.args.grid); got != tt.want {
				t.Errorf("countNegatives() = %v, want %v", got, tt.want)
			}
		})
	}
}
