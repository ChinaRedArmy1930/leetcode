package main

import (
	"reflect"
	"testing"
)

func Test_grayCode(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{7},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := grayCode(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("grayCode() = %v, want %v", got, tt.want)
			}
		})
	}
}
