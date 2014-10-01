package narrowphase

import (
	"github.com/LSFN/dyn4go/geometry"
	"math"
)

type SAT struct{}

func (s *SAT) DetectPenetration(convex1 geometry.Convexer, transform1 *geometry.Transform, convex2 geometry.Convexer, transform2 *geometry.Transform, penetration *Penetration) bool {
	circle1, ok1 := convex1.(*geometry.Circle)
	circle2, ok2 := convex2.(*geometry.Circle)
	if ok1 && ok2 {
		return DetectCirclePenetration(circle1, transform1, circle2, transform2, penetration)
	}
	n := new(geometry.Vector2)
	overlap := math.MaxFloat64
	foci1 := convex1.GetFoci(transform1)
	foci2 := convex2.GetFoci(transform2)
	axes1 := convex1.GetAxes(foci2, transform1)
	axes2 := convex2.GetAxes(foci1, transform2)
	if axes1 != nil {
		for _, axis := range axes1 {
			if !axis.IsZero() {
				intervalA := convex1.ProjectVector2Transform(axis, transform1)
				intervalB := convex2.ProjectVector2Transform(axis, transform2)
				if !intervalA.Overlaps(intervalB) {
					return false
				} else {
					o := intervalA.GetOverlap(intervalB)
					if intervalA.Contains(intervalB) || intervalB.Contains(intervalA) {
						max := math.Abs(intervalA.GetMax() - intervalB.GetMax())
						min := math.Abs(intervalA.GetMin() - intervalB.GetMin())
						if max > min {
							axis.Negate()
							o += min
						} else {
							o += max
						}
					}
					if o < overlap {
						overlap = o
						n = axis
					}
				}
			}
		}
	}
	if axes2 != nil {
		for _, axis := range axes2 {
			if !axis.IsZero() {
				intervalA := convex1.ProjectVector2Transform(axis, transform1)
				intervalB := convex2.ProjectVector2Transform(axis, transform2)
				if !intervalA.Overlaps(intervalB) {
					return false
				} else {
					o := intervalA.GetOverlap(intervalB)
					if intervalA.Contains(intervalB) || intervalB.Contains(intervalA) {
						max := math.Abs(intervalA.GetMax() - intervalB.GetMax())
						min := math.Abs(intervalA.GetMin() - intervalB.GetMin())
						if max > min {
							axis.Negate()
							o += min
						} else {
							o += max
						}
					}
					if o < overlap {
						overlap = o
						n = axis
					}
				}
			}
		}
	}
	c1 := transform1.GetTransformedVector2(convex1.GetCenter())
	c2 := transform2.GetTransformedVector2(convex2.GetCenter())
	cToC := c1.HereToVector2(c2)
	if cToC.DotVector2(n) < 0 {
		n.Negate()
	}
	penetration.normal = n
	penetration.depth = overlap
	return true
}

func (s *SAT) Detect(convex1 geometry.Convexer, transform1 *geometry.Transform, convex2 geometry.Convexer, transform2 *geometry.Transform) bool {
	circle1, ok1 := convex1.(*geometry.Circle)
	circle2, ok2 := convex2.(*geometry.Circle)
	if ok1 && ok2 {
		return DetectCircle(circle1, transform1, circle2, transform2)
	}
	foci1 := convex1.GetFoci(transform1)
	foci2 := convex2.GetFoci(transform2)
	axes1 := convex1.GetAxes(foci2, transform1)
	axes2 := convex2.GetAxes(foci1, transform2)
	if axes1 != nil {
		for _, axis := range axes1 {
			if !axis.IsZero() {
				intervalA := convex1.ProjectVector2Transform(axis, transform1)
				intervalB := convex2.ProjectVector2Transform(axis, transform2)
				if !intervalA.Overlaps(intervalB) {
					return false
				}
			}
		}
	}
	if axes2 != nil {
		for _, axis := range axes2 {
			if !axis.IsZero() {
				intervalA := convex1.ProjectVector2Transform(axis, transform1)
				intervalB := convex2.ProjectVector2Transform(axis, transform2)
				if !intervalA.Overlaps(intervalB) {
					return false
				}
			}
		}
	}
	return true
}
