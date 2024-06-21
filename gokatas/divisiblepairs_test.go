package gokatas

import "testing"

func Test_countDivisiblePairs(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				nums: []int{3, 1, 2, 2, 2, 1, 3},
				k:    2,
			},
			want: 4,
		},
		{
			args: args{
				nums: []int{1, 2, 3, 4},
				k:    1,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countDivisiblePairs(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("countDivisiblePairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
