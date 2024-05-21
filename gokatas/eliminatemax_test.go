package gokatas

import "testing"

func Test_eliminateMax(t *testing.T) {
	type args struct {
		dist  []int
		speed []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				dist:  []int{1, 3, 4},
				speed: []int{1, 1, 1},
			},
			want: 3,
		},
		{
			args: args{
				dist:  []int{1, 1, 2, 3},
				speed: []int{1, 1, 1, 1},
			},
			want: 1,
		},
		{
			args: args{
				dist:  []int{3, 2, 4},
				speed: []int{5, 3, 2},
			},
			want: 1,
		},
		{
			args: args{
				dist:  []int{4, 2, 3},
				speed: []int{2, 1, 1},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := eliminateMax(tt.args.dist, tt.args.speed); got != tt.want {
				t.Errorf("eliminateMax() = %v, want %v", got, tt.want)
			}
		})
	}
}
