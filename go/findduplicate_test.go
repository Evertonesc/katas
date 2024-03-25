package katas

import "testing"

func Test_findDuplicate(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "should return the repeated number in nums",
			args: args{
				nums: []int{1, 1, 1, 1, 1},
			},
			want: 1,
		},
		{
			name: "should return the repeated number in nums",
			args: args{
				nums: []int{1, 3, 4, 2, 2},
			},
			want: 2,
		},
		{
			name: "should return the repeated number in nums",
			args: args{
				nums: []int{3, 1, 3, 4, 2},
			},
			want: 3,
		},
		{
			name: "should return the repeated number in nums",
			args: args{
				nums: []int{3, 3, 3, 3, 3},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDuplicate(tt.args.nums); got != tt.want {
				t.Errorf("findDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
