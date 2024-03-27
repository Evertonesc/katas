package gokatas

import (
	"testing"
)

func Test_RoverCMD(t *testing.T) {
	type args struct {
		cmd string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "should turn the rover direction to east",
			args: args{
				cmd: "R",
			},
			want: "0:0:E",
		},
		{
			name: "should turn the rover direction to weast",
			args: args{
				cmd: "L",
			},
			want: "0:0:W",
		},
		{
			name: "should turn the rover direction to south using two R cmds",
			args: args{
				cmd: "RR",
			},
			want: "0:0:S",
		},
		{
			name: "should turn the rover direction to south using two L cmds",
			args: args{
				cmd: "LL",
			},
			want: "0:0:S",
		},
		{
			name: "should turn rover to the east and move",
			args: args{
				cmd: "RM",
			},
			want: "1:0:E",
		},
		{
			name: "should turn rover to the weast and move",
			args: args{
				cmd: "LM",
			},
			want: "-1:0:W",
		},
		{
			name: "should move rover to north without turn direction",
			args: args{
				cmd: "M",
			},
			want: "0:1:N",
		},
		{
			name: "should move rover to south with two L commands",
			args: args{
				cmd: "LLM",
			},
			want: "0:-1:S",
		},
		{
			name: "should calculate complext commands",
			args: args{
				cmd: "RMMLML",
			},
			want: "2:1:W",
		},
		{
			name: "should calculate complext commands",
			args: args{
				cmd: "LLMMRMMMRMMMMLML",
			},
			want: "-4:2:S",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RoverCMD(tt.args.cmd); got != tt.want {
				t.Errorf("roverCMD() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComputedCMD_CalcDirection(t *testing.T) {
	type fields struct {
		Direction string
		XAxis     int
		YAxis     int
	}
	type args struct {
		directionCMD rune
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		wantXAxis     int
		wantYAxis     int
		wantDirection string
	}{
		{
			name: "should move from north to east",
			fields: fields{
				Direction: NorthDirection,
				XAxis:     0,
				YAxis:     0,
			},
			args: args{
				RightCMD,
			},
			wantXAxis:     0,
			wantYAxis:     0,
			wantDirection: EastDirection,
		},
		{
			name: "should move from east to south",
			fields: fields{
				Direction: EastDirection,
				XAxis:     0,
				YAxis:     0,
			},
			args: args{
				RightCMD,
			},
			wantXAxis:     0,
			wantYAxis:     0,
			wantDirection: SouthDirection,
		},
		{
			name: "should move from south to weast",
			fields: fields{
				Direction: SouthDirection,
				XAxis:     0,
				YAxis:     0,
			},
			args: args{
				RightCMD,
			},
			wantXAxis:     0,
			wantYAxis:     0,
			wantDirection: WeastDirection,
		},
		{
			name: "should move from weast to north",
			fields: fields{
				Direction: WeastDirection,
				XAxis:     0,
				YAxis:     0,
			},
			args: args{
				RightCMD,
			},
			wantXAxis:     0,
			wantYAxis:     0,
			wantDirection: NorthDirection,
		},
		{
			name: "should move from north to weast",
			fields: fields{
				Direction: NorthDirection,
				XAxis:     0,
				YAxis:     0,
			},
			args: args{
				LeftCMD,
			},
			wantXAxis:     0,
			wantYAxis:     0,
			wantDirection: WeastDirection,
		},
		{
			name: "should move from weast to south",
			fields: fields{
				Direction: WeastDirection,
				XAxis:     0,
				YAxis:     0,
			},
			args: args{
				LeftCMD,
			},
			wantXAxis:     0,
			wantYAxis:     0,
			wantDirection: SouthDirection,
		},
		{
			name: "should move from south to east",
			fields: fields{
				Direction: SouthDirection,
				XAxis:     0,
				YAxis:     0,
			},
			args: args{
				LeftCMD,
			},
			wantXAxis:     0,
			wantYAxis:     0,
			wantDirection: EastDirection,
		},
		{
			name: "should move from east to north",
			fields: fields{
				Direction: EastDirection,
				XAxis:     0,
				YAxis:     0,
			},
			args: args{
				LeftCMD,
			},
			wantXAxis:     0,
			wantYAxis:     0,
			wantDirection: NorthDirection,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cc := &ComputedCMD{
				Direction: tt.fields.Direction,
				XAxis:     tt.fields.XAxis,
				YAxis:     tt.fields.YAxis,
			}
			cc.CalcDirection(tt.args.directionCMD)

			if cc.XAxis != tt.wantXAxis {
				t.Errorf("x axis = %v, want %v", cc.XAxis, tt.wantXAxis)
			}

			if cc.YAxis != tt.wantYAxis {
				t.Errorf("y axis = %v, want %v", cc.YAxis, tt.wantYAxis)
			}

			if cc.Direction != tt.wantDirection {
				t.Errorf("direction = %v, want %v", cc.Direction, tt.wantDirection)
			}
		})
	}
}
