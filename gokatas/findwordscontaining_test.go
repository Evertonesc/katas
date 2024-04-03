package gokatas

import (
	"reflect"
	"testing"
)

func Test_findWordsContaining(t *testing.T) {
	type args struct {
		words []string
		x     byte
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "should return the index of words that have x",
			args: args{
				words: []string{"eggs, moon"},
				x:     'a',
			},
			want: []int{},
		},
		{
			name: "should return the index of words that have x",
			args: args{
				words: []string{"eggs"},
				x:     'e',
			},
			want: []int{0},
		},
		{
			name: "should return the index of words that have x",
			args: args{
				words: []string{"leet", "code"},
				x:     'e',
			},
			want: []int{0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findWordsContaining(tt.args.words, tt.args.x); !reflect.DeepEqual(
				got,
				tt.want,
			) {
				t.Errorf("findWordsContaining() = %v, want %v", got, tt.want)
			}
		})
	}
}
