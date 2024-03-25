package katas

import "testing"

func Test_singleNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "should return the single number that do not repeat",
			args: args{
				nums: []int{1},
			},
			want: 1,
		},
		{
			name: "should return the single number that do not repeat",
			args: args{
				nums: []int{2, 2, 1},
			},
			want: 1,
		},
		{
			name: "should return the single number that do not repeat",
			args: args{
				nums: []int{1, 0, 1},
			},
			want: 0,
		},
		{
			name: "should return the single number that do not repeat",
			args: args{
				nums: []int{4, 1, 2, 1, 2},
			},
			want: 4,
		},
		{
			name: "should return the single number that do not repeat",
			args: args{
				nums: []int{4, 3, 5, 6, 6, 5, 3},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNumber(tt.args.nums); got != tt.want {
				t.Errorf("singleNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
