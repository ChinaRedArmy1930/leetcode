package main

import (
	"reflect"
	"testing"
)

func TestInsertSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "插入排序",
			args: args{[]int{6, 5, 4, 3, 2, 1}},
			want: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name: "插入排序",
			args: args{[]int{1, 3, 4, 1, 2, 3, 7, 6, 5, 8, 9, 5, 3, 56, 76, 8, 9}},
			want: []int{1, 1, 2, 3, 3, 3, 4, 5, 5, 6, 7, 8, 8, 9, 9, 56, 76},
		},
		{
			name: "插入排序",
			args: args{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name: "插入排序",
			args: args{[]int{1}},
			want: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InsertSort(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
