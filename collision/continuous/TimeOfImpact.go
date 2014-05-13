package continuous

import (
	"github.com/LSFN/dyn4go/collision/narrowphase"
)

type TimeOfImpact struct {
	time       float64
	separation *narrowphase.Separation
}

func NewTimeOfImpact(time float64, separation *narrowphase.Separation) *TimeOfImpact {
	t := new(TimeOfImpact)
	t.time = time
	t.separation = separation
	return t
}

func (t *TimeOfImpact) GetTime() float64 {
	return t.time
}

func (t *TimeOfImpact) SetTime(time float64) {
	t.time = time
}

func (t *TimeOfImpact) GetSeparation() *narrowphase.Separation {
	return t.separation
}

func (t *TimeOfImpact) SetSeparation(separation *narrowphase.Separation) {
	t.separation = separation
}
