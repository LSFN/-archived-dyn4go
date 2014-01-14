package dyn4go

import (
	"math"
)

func RadToDeg(x float64) float64 {
	return x * 180.0 / math.Pi
}

func DegToRad(x float64) float64 {
	return x * math.Pi / 180.0
}
