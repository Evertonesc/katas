package katas

import "testing"

func Test_thirdMax(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "should return the third max distinct for a 1 length array",
			args: args{
				nums: []int{1},
			},
			want: 1,
		},
		{
			name: "should return the third max distinct for a 2 length array",
			args: args{
				nums: []int{1, 2},
			},
			want: 2,
		},
		{
			name: "should return the third max distinct for a 3 length array",
			args: args{
				nums: []int{3, 2, 1},
			},
			want: 1,
		},
		{
			name: "should return the third max distinct for a 3 length array",
			args: args{
				nums: []int{2, 2, 3, 1},
			},
			want: 1,
		},
		{
			name: "should return the third max distinct for a 3 length array",
			args: args{
				nums: []int{5, 2, 2},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := thirdMax(tt.args.nums); got != tt.want {
				t.Errorf("thirdMax() = %v, want %v", got, tt.want)
			}
		})
	}
}
