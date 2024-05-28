package gokatas

import "testing"

func Test_maxWealth(t *testing.T) {
	type args struct {
		accounts [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				accounts: [][]int{
					{1, 5},
					{7, 3},
					{3, 5},
				},
			},
			want: 10,
		},
		{
			args: args{
				accounts: [][]int{
					{1, 2, 3},
					{3, 2, 1},
				},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxWealth(tt.args.accounts); got != tt.want {
				t.Errorf("maxWealth() = %v, want %v", got, tt.want)
			}
		})
	}
}
