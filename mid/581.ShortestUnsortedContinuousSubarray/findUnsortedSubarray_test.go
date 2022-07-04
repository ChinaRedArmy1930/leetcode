package main

import (
	"testing"
)

func Test_findUnsortedSubarray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "2, 6, 4, 8, 10, 9, 15",
			args: args{[]int{2, 6, 4, 8, 10, 9, 15}},
			want: 5,
		},
		{
			name: "1, 2, 3, 4",
			args: args{[]int{1, 2, 3, 4}},
			want: 0,
		},
		{
			name: "1",
			args: args{[]int{1}},
			want: 0,
		},
		{
			name: "1, 3, 2, 2, 2",
			args: args{[]int{1, 3, 2, 2, 2}},
			want: 4,
		},
		{
			name: "1, 2, 4, 5, 3",
			args: args{[]int{1, 2, 4, 5, 3}},
			want: 3,
		},
		{
			name: "2, 1",
			args: args{[]int{2, 1}},
			want: 2,
		},
		{
			name: "5, 4, 3, 2, 1",
			args: args{[]int{5, 4, 3, 2, 1}},
			want: 5,
		},
		{
			name: "1, 2, 5, 3, 4",
			args: args{[]int{1, 2, 5, 3, 4}},
			want: 3,
		},
		{
			name: "1, 3, 5, 4, 2",
			args: args{[]int{1, 3, 5, 4, 2}},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findUnsortedSubarray(tt.args.nums); got != tt.want {
				t.Errorf("findUnsortedSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
