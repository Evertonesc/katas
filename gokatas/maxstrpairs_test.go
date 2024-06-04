package gokatas

import "testing"

func Test_maxNumStrPairs(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				words: []string{"ab", "ba"},
			},
			want: 1,
		},
		{
			args: args{
				words: []string{"ab", "ba", "cd", "ef"},
			},
			want: 1,
		},
		{
			args: args{
				words: []string{"cd", "ac", "dc", "ca", "zz"},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxNumStrPairs(tt.args.words); got != tt.want {
				t.Errorf("maxNumStrPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
