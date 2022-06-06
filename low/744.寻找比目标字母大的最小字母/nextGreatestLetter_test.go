package main

import "testing"

func Test_nextGreatestLetter(t *testing.T) {
	type args struct {
		letters []byte
		target  byte
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{
			name: "[\"c\", \"f\", \"j\"]，target = \"a\"",
			args: args{
				letters: []byte{'c', 'f', 'j'},
				target:  'a',
			},
			want: 'c',
		},
		{
			name: "[\"c\", \"f\", \"j\"]，target = \"c\"",
			args: args{
				letters: []byte{'c', 'f', 'j'},
				target:  'c',
			},
			want: 'f',
		},
		{
			name: "[\"c\", \"f\", \"j\"]，target = \"d\"",
			args: args{
				letters: []byte{'c', 'f', 'j'},
				target:  'd',
			},
			want: 'f',
		},
		{
			name: "[\"c\", \"f\", \"j\"]，target = \"j\"",
			args: args{
				letters: []byte{'c', 'f', 'j'},
				target:  'j',
			},
			want: 'c',
		},
		{
			name: "[\"c\", \"f\", \"j\"]，target = \"g\"",
			args: args{
				letters: []byte{'c', 'f', 'j'},
				target:  'g',
			},
			want: 'j',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextGreatestLetter(tt.args.letters, tt.args.target); got != tt.want {
				t.Errorf("nextGreatestLetter() = %c, want %c", got, tt.want)
			}
		})
	}
}
