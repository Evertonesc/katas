package gokatas

import "testing"

func Test_countPair(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				nums:   []int{-1, 1, 2, 3, 1},
				target: 2,
			},
			want: 3,
		},
		{
			args: args{
				nums:   []int{-6, 2, 5, -2, -7, -1, 3},
				target: -2,
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPair(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("countPair() = %v, want %v", got, tt.want)
			}
		})
	}
}
