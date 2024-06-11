package gokatas

import (
	"reflect"
	"testing"
)

func Test_findDuplicates(t *testing.T) {
	type args struct {
		paths []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			args: args{
				paths: []string{
					"root/a 1.txt(abcd) 2.txt(efgh)",
					"root/c 3.txt(abcd)",
				},
			},
			want: [][]string{
				{
					"root/a/1.txt",
					"root/c/3.txt",
				},
			},
		},
		{
			args: args{
				paths: []string{
					"root/a 1.txt(abcd) 2.txt(efgh)",
					"root/c 3.txt(abcd)",
					"root/c/d 4.txt(efgh)",
					"root 4.txt(efgh)",
				},
			},
			want: [][]string{
				{
					"root/a/1.txt", "root/c/3.txt",
				},
				{
					"root/a/2.txt", "root/c/d/4.txt", "root/4.txt",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDuplicates(tt.args.paths); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
