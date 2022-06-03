package main

import (
	"reflect"
	"testing"
)

func Test_rotate(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1,2,3,4,5,6,7",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6, 7},
				k:    1,
			},
			want: []int{7, 1, 2, 3, 4, 5, 6},
		},
		{
			name: "1,2,3,4,5,6,7",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6, 7},
				k:    2,
			},
			want: []int{6, 7, 1, 2, 3, 4, 5},
		},
		{
			name: "1,2,3,4,5,6,7",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6, 7},
				k:    3,
			},
			want: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			name: "-1,-100,3,99",
			args: args{
				nums: []int{-1, -100, 3, 99},
				k:    2,
			},
			want: []int{3, 99, -1, -100},
		},
		{
			name: "1,2",
			args: args{
				nums: []int{1, 2},
				k:    0,
			},
			want: []int{1, 2},
		},
		{
			name: "1,2",
			args: args{
				nums: []int{1, 2},
				k:    3,
			},
			want: []int{2, 1},
		},
		{
			name: "1,2",
			args: args{
				nums: []int{1, 2},
				k:    2,
			},
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate(tt.args.nums, tt.args.k)
			t.Logf("nums = %v, want %v", tt.args.nums, tt.want)
			if !reflect.DeepEqual(tt.args.nums, tt.want) {
				t.Errorf("nums = %v, want %v", tt.args.nums, tt.want)
			}
		})
	}
}
