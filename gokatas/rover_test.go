package gokatas

import "testing"

func Test_roverCMD(t *testing.T) {
	type args struct {
		cmd string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "should move rover to the disired position",
			args: args{
				cmd: "R",
			},
			want: "0:0:E",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := roverCMD(tt.args.cmd); got != tt.want {
				t.Errorf("roverCMD() = %v, want %v", got, tt.want)
			}
		})
	}
}
