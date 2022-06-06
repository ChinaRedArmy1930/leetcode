package main

import (
	"reflect"
	"testing"
)

func Test_searchRange(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "5,7,7,8,8,10",
			args: args{[]int{5, 7, 7, 8, 8, 10}, 7},
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchRange(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchRange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_searchLeft(t *testing.T) {
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
			name: "找左边",
			args: args{[]int{1, 2, 2, 2, 2, 3, 3, 3, 3}, 2},
			want: 1,
		},
		{
			name: "找左边",
			args: args{[]int{1, 2, 2, 2, 2, 3, 3, 3, 3}, 1},
			want: 0,
		},
		{
			name: "找左边",
			args: args{[]int{1, 2, 2, 2, 2, 3, 3, 3, 3}, 3},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchLeft(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("searchLeft() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_searchRight(t *testing.T) {
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
			name: "找右边",
			args: args{[]int{1, 2, 2, 2, 2, 3, 3, 3, 3}, 2},
			want: 4,
		},
		{
			name: "找右边",
			args: args{[]int{1, 2, 2, 2, 2, 3, 3, 3, 3}, 1},
			want: 0,
		},
		{
			name: "找右边",
			args: args{[]int{1, 2, 2, 2, 2, 3, 3, 3, 3}, 3},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchRight(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("searchRight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_searchRange1(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "[]int{5, 7, 7, 8, 8, 10}, 7",
			args: args{[]int{5, 7, 7, 8, 8, 10}, 7},
			want: []int{1, 2},
		},
		{
			name: "[]int{5, 7, 7, 8, 8, 10}, 6",
			args: args{[]int{5, 7, 7, 8, 8, 10}, 6},
			want: []int{-1, -1},
		},
		{
			name: "[]int{}, 6",
			args: args{[]int{}, 6},
			want: []int{-1, -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchRange(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
