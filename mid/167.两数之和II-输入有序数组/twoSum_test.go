package main

import (
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		numbers []int
		target  int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "2,7,11,15",
			args: args{[]int{2, 7, 11, 15}, 9},
			want: []int{1, 2},
		},
		{
			name: "2,3,4",
			args: args{[]int{2, 3, 4}, 6},
			want: []int{1, 3},
		},
		{
			name: "-1, 0",
			args: args{[]int{-1, 0}, -1},
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.numbers, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_twoSum2(t *testing.T) {
	type args struct {
		numbers []int
		target  int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "2,7,11,15",
			args: args{[]int{2, 7, 11, 15}, 9},
			want: []int{1, 2},
		},
		{
			name: "2,3,4",
			args: args{[]int{2, 3, 4}, 6},
			want: []int{1, 3},
		},
		{
			name: "-1, 0",
			args: args{[]int{-1, 0}, -1},
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum2(tt.args.numbers, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
