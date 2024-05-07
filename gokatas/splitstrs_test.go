package gokatas

import (
	"reflect"
	"testing"
)

func Test_solution(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			args: args{
				str: "ab",
			},
			want: []string{"ab"},
		},
		{
			args: args{
				str: "abc",
			},
			want: []string{"ab", "c_"},
		},
		{
			args: args{
				str: "dawsd",
			},
			want: []string{"da", "ws", "d_"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
