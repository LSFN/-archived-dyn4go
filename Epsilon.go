package dyn4go

type Epsilon struct {
	epsilon float64
}

func (e *Epsilon) getEpsilon() float64 {
	if e.epsilon == 0 {
		e.computeEpsilon()
	}
	return e.epsilon
}

func (e *Epsilon) computeEpsilon() {
	e.epsilon = 0.5
	for 1.0+e.epsilon > 1.0 {
		e.epsilon *= 0.5
	}
}
