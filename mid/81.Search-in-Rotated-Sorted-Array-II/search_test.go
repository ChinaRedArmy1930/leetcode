package main

import "testing"

func Test_search(t *testing.T) {
	type args struct {
		q      []int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "nums = [2,5,6,0,0,1,2], target = 0",
			args: args{[]int{2, 5, 6, 0, 0, 1, 2}, 0},
			want: true,
		},
		{
			name: "nums = [2,5,6,0,0,1,2], target = 3",
			args: args{[]int{2, 5, 6, 0, 0, 1, 2}, 3},
			want: false,
		},
		{
			name: "nums = [2,5], target = 3",
			args: args{[]int{2, 5}, 3},
			want: false,
		},
		{
			name: "nums = [2,5], target = 2",
			args: args{[]int{2, 5}, 2},
			want: true,
		},
		{
			name: "nums = [2,5], target = 5",
			args: args{[]int{2, 5}, 5},
			want: true,
		},
		{
			name: "nums = [5,2], target = 2",
			args: args{[]int{5, 2}, 5},
			want: true,
		},
		{
			name: "nums = [5,2], target = 2",
			args: args{[]int{5, 2}, 2},
			want: true,
		},
		{
			name: "nums = [1,0,1,1,1], target = 0",
			args: args{[]int{1, 0, 1, 1, 1}, 0},
			want: true,
		},
		{
			name: "nums = [1,1,1,1,1,1,1,1,1,1,1,1,1,2,1,1,1,1,1], target = 2",
			args: args{[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1}, 2},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.q, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
