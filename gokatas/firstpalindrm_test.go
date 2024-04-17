package gokatas

import "testing"

func Test_firstPalindrome(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				words: []string{
					"abc",
				},
			},
			want: "",
		},
		{
			args: args{
				words: []string{
					"aba",
				},
			},
			want: "aba",
		},
		{
			args: args{
				words: []string{
					"abc",
					"car",
					"ada",
					"racecar",
					"cool",
				},
			},
			want: "ada",
		},
		{
			args: args{
				words: []string{
					"notapalindrome",
					"racecar",
				},
			},
			want: "racecar",
		},
		{
			args: args{
				words: []string{
					"def",
					"ghi",
				},
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstPalindrome(tt.args.words); got != tt.want {
				t.Errorf("firstPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
