package geometry

import (
	"math"
)

type Interval struct {
	min, max float64
}

func NewIntervalFromMinMax(min, max float64) *Interval {
	i := new(Interval)
	i.min = min
	i.max = max
	return i
}

func NewIntervalFromIntervat(orig *Interval) *Interval {
	i := *orig
	return &i
}

func (i *Interval) GetMin() float64 {
	return i.min
}

func (i *Interval) GetMax() float64 {
	return i.max
}

func (i *Interval) SetMin(min float64) {
	i.min = min
}

func (i *Interval) SetMax(max float64) {
	i.max = max
}

func (i *Interval) IncludesInclusive(x float64) bool {
	return x <= i.max && x >= i.min
}

func (i *Interval) IncludesExclusive(x float64) bool {
	return x < i.max && x > i.min
}

func (i *Interval) IncludesInclusiveMin(x float64) bool {
	return x < i.max && x >= i.min
}

func (i *Interval) IncludesInclusiveMax(x float64) bool {
	return x <= i.max && x > i.min
}

func (i *Interval) Overlaps(i2 *Interval) bool {
	return !(i.min > i2.max || i2.min > i.max)
}

func (i *Interval) GetOverlap(i2 *Interval) float64 {
	if i.Overlaps(i2) {
		return math.Min(i.max, i2.max) - math.Max(i.min, i2.min)
	}
	return 0
}

func (i *Interval) Clamp(x float64) float64 {
	return Clamp(x, i.min, i.max)
}

func Clamp(x, min, max float64) float64 {
	if x <= max && x >= min {
		return x
	} else if max < x {
		return max
	} else {
		return min
	}
}

func (i *Interval) IsDegenerate() bool {
	return i.min == i.max
}

func (i *Interval) IsDegenerateWithError(e float64) bool {
	return math.Abs(i.max-i.min) <= e
}

func (i *Interval) Contains(i2 *Interval) bool {
	return i2.min > i.min && i2.max < i2.max
}

func (i *Interval) Union(i2 *Interval) {
	i.min = math.Min(i.min, i2.min)
	i.max = math.Max(i.max, i2.max)
}

func (i *Interval) GetUnion(i2 *Interval) *Interval {
	i3 := new(Interval)
	i3.min = math.Min(i.min, i2.min)
	i3.max = math.Max(i.max, i2.max)
	return i3
}

func (i *Interval) Intersection(i2 *Interval) {
	if i.Overlaps(i2) {
		i.min = math.Max(i.min, i2.min)
		i.max = math.Min(i.max, i2.max)
	} else {
		i.min = 0
		i.max = 0
	}
}

func (i *Interval) GetIntersection(i2 *Interval) *Interval {
	if i.Overlaps(i2) {
		return NewIntervalFromMinMax(math.Max(i2.min, i.min), math.Min(i2.max, i.max))
	} else {
		return NewIntervalFromMinMax(0, 0)
	}
}

func (i *Interval) Distance(i2 *Interval) float64 {
	if !i.Overlaps(i2) {
		if i.max < i2.min {
			return i2.min - i.max
		} else {
			return i.min - i2.max
		}
	}
	return 0
}

func (i *Interval) Expand(x float64) {
	e := x * 0.5
	i.min -= e
	i.max += e
	if x < 0 && i.min > i.max {
		p := (i.min + i.max) * 0.5
		i.min = p
		i.max = p
	}
}

func (i *Interval) GetExpanded(x float64) *Interval {
	e := x * 0.5
	min := i.min - e
	max := i.max + e
	if x < 0 && min > max {
		p := (min + max) * 0.5
		min = p
		max = p
	}
	return NewIntervalFromMinMax(min, max)
}

func (i *Interval) GetLength() float64 {
	return i.max - i.min
}
