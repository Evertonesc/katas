package katas

import "testing"

func TestIsValidWalk(t *testing.T) {
	type args struct {
		walk []rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "should return an invalid walk",
			args: args{
				walk: []rune{'n', 's'},
			},
			want: false,
		},
		{
			name: "should return a valid walk",
			args: args{
				walk: []rune{'n', 's', 'n', 's', 'n', 's', 'n', 's', 'n', 's'},
			},
			want: true,
		},
		{
			name: "should return an invalid walk",
			args: args{
				walk: []rune{'n', 'n', 'n', 's', 'n', 'w', 'e', 'e', 'e', 's'},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidWalk(tt.args.walk); got != tt.want {
				t.Errorf("IsValidWalk() = %v, want %v", got, tt.want)
			}
		})
	}
}
