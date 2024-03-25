package katas

import "testing"

func Test_removeElement(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "should return the k not equal to val as 1",
			args: args{
				nums: []int{4, 3},
				val:  4,
			},
			want: 1,
		},
		{
			name: "should return the k not equal to val as 2",
			args: args{
				nums: []int{4, 3, 4},
				val:  4,
			},
			want: 1,
		},
		{
			name: "should return the k not equal to val as 4",
			args: args{
				nums: []int{4, 3, 4, 2, 6, 7},
				val:  4,
			},
			want: 4,
		},
		{
			name: "should return the k not equal to val as 2",
			args: args{
				nums: []int{3, 2, 2, 3},
				val:  3,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElement(tt.args.nums, tt.args.val); got != tt.want {
				t.Errorf("removeElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
