package dyn4go

var Epsilon = 0.5

func init() {
	for 1.0+Epsilon > 1.0 {
		Epsilon *= 0.5
	}
}
