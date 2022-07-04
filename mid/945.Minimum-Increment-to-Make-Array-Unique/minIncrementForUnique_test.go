package main

import "testing"

func Test_minIncrementForUnique(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1,2,2",
			args: args{nums: []int{1, 2, 2}},
			want: 1,
		},
		{
			name: "3,2,1,2,1,7",
			args: args{nums: []int{3, 2, 1, 2, 1, 7}},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minIncrementForUnique(tt.args.nums); got != tt.want {
				t.Errorf("minIncrementForUnique() = %v, want %v", got, tt.want)
			}
		})
	}
}
