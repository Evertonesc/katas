package gokatas

import "testing"

func Test_findTheArrayConcVal(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			args: args{
				nums: []int{15, 49},
			},
			want: 1549,
		},
		{
			args: args{
				nums: []int{15, 49, 10},
			},
			want: 1559,
		},
		{
			args: args{
				nums: []int{7, 52, 2, 4},
			},
			want: 596,
		},
		{
			args: args{
				nums: []int{5, 14, 13, 8, 12},
			},
			want: 673,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTheArrayConcVal(tt.args.nums); got != tt.want {
				t.Errorf("findTheArrayConcVal() = %v, want %v", got, tt.want)
			}
		})
	}
}
