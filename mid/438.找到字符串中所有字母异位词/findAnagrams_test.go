package _38_找到字符串中所有字母异位词

import (
	"reflect"
	"testing"
)

func Test_findAnagrams(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "s = \"cbaebabacd\", p = \"abc\"",
			args: args{s: "cbaebabacd", p: "abc"},
			want: []int{0, 6},
		},
		{
			name: "s = \"abab\", p = \"ab\"",
			args: args{s: "abab", p: "ab"},
			want: []int{0, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findAnagrams(tt.args.s, tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}
