package gokatas

import (
	"reflect"
	"testing"
)

func Test_groupThePeople(t *testing.T) {
	type args struct {
		groupSizes []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			args: args{
				groupSizes: []int{3, 3, 3, 3, 3, 1, 3},
			},
			want: [][]int{
				{5}, {0, 1, 2}, {3, 4, 6},
			},
		},
		{
			args: args{
				groupSizes: []int{2, 2, 1, 1, 1, 1, 1, 1},
			},
			want: [][]int{
				{2}, {3}, {4}, {5}, {6}, {7}, {0, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := groupThePeople(tt.args.groupSizes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("groupThePeople() = %v, want %v", got, tt.want)
			}
		})
	}
}
