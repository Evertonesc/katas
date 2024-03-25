package katas

import "testing"

func TestToCamelCase(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "should return a camel case from a dashed word",
			args: args{
				s: "the-stealth-warrior",
			},
			want: "theStealthWarrior",
		},
		{
			name: "should return a camel case from a dashed word",
			args: args{
				s: "the_Stealth_Warrior",
			},
			want: "theStealthWarrior",
		},
		{
			name: "should return a camel case from a dashed word",
			args: args{
				s: "The_Stealth-Warrior",
			},
			want: "TheStealthWarrior",
		},
		{
			name: "should return a camel case from a dashed word",
			args: args{
				s: "The-Stealth_Warrior",
			},
			want: "TheStealthWarrior",
		},
		{
			name: "should return a camel case from a dashed word",
			args: args{
				s: "The--Stealth__Warrior",
			},
			want: "TheStealthWarrior",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToCamelCase(tt.args.s); got != tt.want {
				t.Errorf("ToCamelCase() = %v, want %v", got, tt.want)
			}
		})
	}
}
