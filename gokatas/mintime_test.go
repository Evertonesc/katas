package gokatas

import "testing"

func Test_minTimeToVisitAllPoints(t *testing.T) {
	type args struct {
		points [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				points: [][]int{
					{1, 1}, {2, 2},
				},
			},
			want: 1,
		},
		{
			args: args{
				points: [][]int{
					{3, 2}, {-2, 2},
				},
			},
			want: 5,
		},
		{
			args: args{
				points: [][]int{
					{1, 1}, {3, 4}, {-1, 0},
				},
			},
			want: 7,
		},
		{
			args: args{
				points: [][]int{
					{1, 1},
					{2, 2},
					{5, 5},
					{0, 0},
					{-2, -3},
				},
			},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minTimeToVisitAllPoints(tt.args.points); got != tt.want {
				t.Errorf("minTimeToVisitAllPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
