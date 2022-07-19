package main

import "testing"

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "nums = [4,5,6,7,0,1,2], target = 0",
			args: args{[]int{4, 5, 6, 7, 0, 1, 2}, 0},
			want: 4,
		},
		{
			name: "nums = [4,5,6,7,0,1,2], target = 3",
			args: args{[]int{4, 5, 6, 7, 0, 1, 2}, 3},
			want: -1,
		},
		{
			name: "nums = [4,5,6,7,8,9,10,0,1,2], target = 0",
			args: args{[]int{4, 5, 6, 7, 8, 9, 10, 0, 1, 2}, 0},
			want: 7,
		},
		{
			name: "nums = [1,3], target = 3",
			args: args{[]int{1, 3}, 3},
			want: 1,
		},
		{
			name: "nums = [1,3], target = 3",
			args: args{[]int{1, 3}, 1},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
