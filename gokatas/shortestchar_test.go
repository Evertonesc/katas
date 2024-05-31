package gokatas

import (
	"reflect"
	"testing"
)

func Test_shortestsToChar(t *testing.T) {
	type args struct {
		s string
		c byte
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				s: "ab",
				c: 'b',
			},
			want: []int{1, 0},
		},
		{
			args: args{
				s: "aaab",
				c: 'b',
			},
			want: []int{3, 2, 1, 0},
		},
		{
			args: args{
				s: "aaba",
				c: 'b',
			},
			want: []int{2, 1, 0, 1},
		},
		{
			args: args{
				s: "loveleetcode",
				c: 'e',
			},
			want: []int{3, 2, 1, 0, 1, 0, 0, 1, 2, 2, 1, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shortestsToChar(tt.args.s, tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("shortestsToChar() = %v, want %v", got, tt.want)
			}
		})
	}
}
