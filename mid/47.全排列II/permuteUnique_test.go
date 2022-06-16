package main

import (
	"reflect"
	"testing"
)

func Test_permuteUnique(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "nums = [1,1,2]",
			args: args{[]int{1, 1, 2}},
			want: [][]int{[]int{1, 1, 2}, []int{1, 2, 1}, []int{2, 1, 1}},
		},
		{
			name: "nums = [1,2,3]",
			args: args{[]int{1, 2, 3}},
			want: [][]int{[]int{1, 2, 3}, []int{1, 3, 2}, []int{2, 1, 3}, []int{2, 3, 1}, []int{3, 1, 2}, []int{3, 2, 1}},
		},
		{
			name: "nums = [[1,1,0,0,1,-1,-1,1]",
			args: args{[]int{1, 1, 0, 0, 1, -1, -1, 1}},
			want: [][]int{[]int{1, 2, 3}, []int{1, 3, 2}, []int{2, 1, 3}, []int{2, 3, 1}, []int{3, 1, 2}, []int{3, 2, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := permuteUnique(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("permuteUnique() = %v, want %v", got, tt.want)
			}
		})
	}
}
