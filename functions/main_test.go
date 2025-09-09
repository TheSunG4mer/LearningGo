package main

import "testing"

func Test_add(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"2+2=4", args{x: 2, y: 2}, 4},
		{"0+0=0", args{x: 0, y: 0}, 0},
		{"-2+2=0", args{x: -2, y: 2}, 0},
		{"-3-2=-5", args{x: -3, y: -2}, -5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := add(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("add() = %v, want %v", got, tt.want)
			}
		})
	}
}
