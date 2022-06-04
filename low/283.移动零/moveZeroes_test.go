package main

import (
	"reflect"
	"testing"
)

func Test_moveZeroes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "0,1,0,3,12",
			args: args{[]int{0, 1, 0, 3, 12}},
			want: []int{1, 3, 12, 0, 0},
		},
		{
			name: "0",
			args: args{[]int{0}},
			want: []int{0},
		},
		{
			name: "1, 0",
			args: args{[]int{1, 0}},
			want: []int{1, 0},
		},
		{
			name: "1, 0, 1",
			args: args{[]int{1, 0, 1}},
			want: []int{1, 1, 0},
		},
		{
			name: "1,0,1,0,1,2,3,4,5,9,0,1",
			args: args{[]int{1, 0, 1, 0, 1, 2, 3, 4, 5, 9, 0, 1}},
			want: []int{1, 1, 1, 2, 3, 4, 5, 9, 1, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moveZeroes(tt.args.nums)
			if !reflect.DeepEqual(tt.args.nums, tt.want) {
				t.Errorf("nums = %v, want %v", tt.args.nums, tt.want)
			}
		})
	}
}

func Test_moveZeroes2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "0,1,0,3,12",
			args: args{[]int{0, 1, 0, 3, 12}},
			want: []int{1, 3, 12, 0, 0},
		},
		{
			name: "0",
			args: args{[]int{0}},
			want: []int{0},
		},
		{
			name: "1, 0",
			args: args{[]int{1, 0}},
			want: []int{1, 0},
		},
		{
			name: "1, 0, 1",
			args: args{[]int{1, 0, 1}},
			want: []int{1, 1, 0},
		},
		{
			name: "1,0,1,0,1,2,3,4,5,9,0,1",
			args: args{[]int{1, 0, 1, 0, 1, 2, 3, 4, 5, 9, 0, 1}},
			want: []int{1, 1, 1, 2, 3, 4, 5, 9, 1, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moveZeroes2(tt.args.nums)
			if !reflect.DeepEqual(tt.args.nums, tt.want) {
				t.Errorf("nums = %v, want %v", tt.args.nums, tt.want)
			}
		})
	}
}

func Test_moveZeroes3(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "0,1,0,3,12",
			args: args{[]int{0, 1, 0, 3, 12}},
			want: []int{1, 3, 12, 0, 0},
		},
		{
			name: "0",
			args: args{[]int{0}},
			want: []int{0},
		},
		{
			name: "1, 0",
			args: args{[]int{1, 0}},
			want: []int{1, 0},
		},
		{
			name: "1, 0, 1",
			args: args{[]int{1, 0, 1}},
			want: []int{1, 1, 0},
		},
		{
			name: "1,0,1,0,1,2,3,4,5,9,0,1",
			args: args{[]int{1, 0, 1, 0, 1, 2, 3, 4, 5, 9, 0, 1}},
			want: []int{1, 1, 1, 2, 3, 4, 5, 9, 1, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moveZeroes3(tt.args.nums)
			if !reflect.DeepEqual(tt.args.nums, tt.want) {
				t.Errorf("nums = %v, want %v", tt.args.nums, tt.want)
			}
		})
	}
}
