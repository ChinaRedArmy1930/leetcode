package main

import (
	"reflect"
	"testing"
)

func Test_insert(t *testing.T) {
	type args struct {
		intervals   [][]int
		newInterval []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "intervals = [[1,3],[6,9]], newInterval = [2,5]",
			args: args{intervals: [][]int{[]int{1, 3}, []int{6, 9}}, newInterval: []int{2, 5}},
			want: [][]int{{1, 5}, {6, 9}},
		},
		{
			name: "intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]",
			args: args{intervals: [][]int{[]int{1, 2}, []int{3, 5}, {6, 7}, {8, 10}, {12, 16}}, newInterval: []int{4, 8}},
			want: [][]int{{1, 2}, {3, 10}, {12, 16}},
		},
		{
			name: "intervals = [], newInterval = [4,8]",
			args: args{intervals: [][]int{}, newInterval: []int{4, 8}},
			want: [][]int{{4, 8}},
		},
		{
			name: "intervals = [1,5], newInterval = [6,8]",
			args: args{intervals: [][]int{{1, 5}}, newInterval: []int{6, 8}},
			want: [][]int{{1, 5}, {6, 8}},
		},
		{
			name: "intervals = [1,5], newInterval = [5,7]",
			args: args{intervals: [][]int{{1, 5}}, newInterval: []int{5, 7}},
			want: [][]int{{1, 7}},
		},
		{
			name: "intervals = [1,5], newInterval = [0.6]",
			args: args{intervals: [][]int{{1, 5}}, newInterval: []int{0, 6}},
			want: [][]int{{0, 6}},
		},
		{
			name: "intervals = [3,5] [12,15], newInterval = [6.6]",
			args: args{intervals: [][]int{{3, 5}, {12, 15}}, newInterval: []int{6, 6}},
			want: [][]int{{3, 5}, {6, 6}, {12, 15}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insert(tt.args.intervals, tt.args.newInterval); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_insert2(t *testing.T) {
	type args struct {
		intervals   [][]int
		newInterval []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "intervals = [[1,3],[6,9]], newInterval = [2,5]",
			args: args{intervals: [][]int{[]int{1, 3}, []int{6, 9}}, newInterval: []int{2, 5}},
			want: [][]int{{1, 5}, {6, 9}},
		},
		{
			name: "intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]",
			args: args{intervals: [][]int{[]int{1, 2}, []int{3, 5}, {6, 7}, {8, 10}, {12, 16}}, newInterval: []int{4, 8}},
			want: [][]int{{1, 2}, {3, 10}, {12, 16}},
		},
		{
			name: "intervals = [], newInterval = [4,8]",
			args: args{intervals: [][]int{}, newInterval: []int{4, 8}},
			want: [][]int{{4, 8}},
		},
		{
			name: "intervals = [1,5], newInterval = [6,8]",
			args: args{intervals: [][]int{{1, 5}}, newInterval: []int{6, 8}},
			want: [][]int{{1, 5}, {6, 8}},
		},
		{
			name: "intervals = [1,5], newInterval = [5,7]",
			args: args{intervals: [][]int{{1, 5}}, newInterval: []int{5, 7}},
			want: [][]int{{1, 7}},
		},
		{
			name: "intervals = [1,5], newInterval = [0.6]",
			args: args{intervals: [][]int{{1, 5}}, newInterval: []int{0, 6}},
			want: [][]int{{0, 6}},
		},
		{
			name: "intervals = [3,5] [12,15], newInterval = [6.6]",
			args: args{intervals: [][]int{{3, 5}, {12, 15}}, newInterval: []int{6, 6}},
			want: [][]int{{3, 5}, {6, 6}, {12, 15}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insert2(tt.args.intervals, tt.args.newInterval); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insert() = %v, want %v", got, tt.want)
			}
		})
	}
}
