package gokatas

import "fmt"

const (
	MoveCMD  = 'M'
	RightCMD = 'R'
	LeftCMD  = 'L'

	NorthDirection = "N"
	SouthDirection = "S"
	EastDirection  = "E"
	WeastDirection = "W"

	XAxis = iota
	YAxis
)

var AxisOfDirectionRef = map[string]int{
	NorthDirection: YAxis,
}

type ComputedCMD struct {
	Direction    string
	XAxis, YAxis int
}

// roverCMD refers to assets/rover_kata.txt
func RoverCMD(cmds string) string {
	// initial Rover state
	computedCMD := &ComputedCMD{
		Direction: NorthDirection,
	}

	for _, cmd := range []rune(cmds) {
		if cmd != MoveCMD {
			computedCMD.CalcDirection(cmd)
			continue
		}

		computedCMD.Move()
	}

	return fmt.Sprintf("%d:%d:%s", computedCMD.XAxis, computedCMD.YAxis, computedCMD.Direction)
}

// CalcDirection calculates which direction the Rover will be facing after a direction command.
// A given direction will always yield the same result for a right or left command
// on a cartesian plan.
// Example: A RightCMD for a NorthDirection will always be EastDirection.
func (cc *ComputedCMD) CalcDirection(directionCMD rune) {
	if directionCMD == RightCMD {
		cc.moveAxisToRight()
		return
	}

	cc.moveAxisToLeft()
}

func (cc *ComputedCMD) moveAxisToRight() {
	switch cc.Direction {
	case NorthDirection:
		cc.Direction = EastDirection
	case EastDirection:
		cc.Direction = SouthDirection
	case SouthDirection:
		cc.Direction = WeastDirection
	case WeastDirection:
		cc.Direction = NorthDirection
	}
}

func (cc *ComputedCMD) moveAxisToLeft() {
	switch cc.Direction {
	case NorthDirection:
		cc.Direction = WeastDirection
	case WeastDirection:
		cc.Direction = SouthDirection
	case SouthDirection:
		cc.Direction = EastDirection
	case EastDirection:
		cc.Direction = NorthDirection
	}
}

func (cc *ComputedCMD) Move() {
	switch cc.Direction {
	case NorthDirection:
		cc.YAxis++
	case SouthDirection:
		cc.YAxis--
	case EastDirection:
		cc.XAxis++
	case WeastDirection:
		cc.XAxis--
	}
}
