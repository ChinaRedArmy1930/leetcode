package main

import (
	"leetcode/common"
	"reflect"
	"testing"
)

func Test_middleNode(t *testing.T) {
	head1 := common.BuildList([]int{1, 2, 3, 4, 5})
	head2 := common.BuildList([]int{1, 2, 3, 3, 4, 5})
	type args struct {
		head *common.ListNode
	}
	tests := []struct {
		name string
		args args
		want *common.ListNode
	}{
		{
			name: "1,2,3,4,5",
			args: args{head1},
			want: head1.Next.Next,
		},
		{
			name: "1, 2, 3, 3, 4, 5",
			args: args{head2},
			want: head2.Next.Next.Next,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := middleNode(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("middleNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
