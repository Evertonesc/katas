package katas

import (
	"reflect"
	"testing"
)

func Test_arrayDiff(t *testing.T) {
	type args struct {
		a []int
		b []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "should return an empty array for empty inputs",
			args: args{
				a: []int{},
				b: []int{},
			},
			want: []int{},
		},
		{
			name: "should return an empty array when the array a is empty",
			args: args{
				a: []int{},
				b: []int{1},
			},
			want: []int{},
		},
		{
			name: "should return an empty array when the array b is empty",
			args: args{
				a: []int{1, 2, 2},
				b: []int{},
			},
			want: []int{1, 2, 2},
		},
		{
			name: "should remove all occurrences of b from a",
			args: args{
				a: []int{1, 2},
				b: []int{1},
			},
			want: []int{2},
		},
		{
			name: "should remove all occurrences of b from a",
			args: args{
				a: []int{1, 2, 2},
				b: []int{1},
			},
			want: []int{2, 2},
		},
		{
			name: "should remove all occurrences of b from a",
			args: args{
				a: []int{1, 2, 2},
				b: []int{2},
			},
			want: []int{1},
		},
		{
			name: "should remove all occurrences of b from a",
			args: args{
				a: []int{1, 2, 3},
				b: []int{1, 2},
			},
			want: []int{3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := arrayDiff(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("arrayDiff() = %v, want %v", got, tt.want)
			}
		})
	}
}
