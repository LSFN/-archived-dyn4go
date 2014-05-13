package continuous

import (
	"github.com/LSFN/dyn4go"
	"github.com/LSFN/dyn4go/collision/narrowphase"
	"github.com/LSFN/dyn4go/geometry"
	"math"
)

type ConservativeAdvancement struct {
	distanceDetector *narrowphase.DistanceDetector
	distanceEpsilon  float64
	maxIterations    int
}

func NewConservativeAdvancement(distanceDetector *narrowphase.DistanceDetector) *narrowphase.DistanceDetector {
	if distanceDetector == nil {
		panic("Distance detector cannot be nil")
	}
	c := new(ConservativeAdvancement)
	c.distanceDetector = distanceDetector
	return c
}

func (c *ConservativeAdvancement) GetTimeOfImpact(convex1 geometry.Convexer, transform1 *geometry.Transform, dp1 *geometry.Vector2, da1 float64, convex2 geometry.Convexer, transform2 *geometry.Transform, dp2 *geometry.Vector2, da2 float64, toi *TimeOfImpact) bool {
	return c.GetTimeOfImpactBounded(convex1, transform1, dp1, da1, convex2, transform2, dp2, da2, 0, 1, toi)
}

func (c *ConservativeAdvancement) GetTimeOfImpactBounded(convex1 geometry.Convexer, transform1 *geometry.Transform, dp1 *geometry.Vector2, da1 float64, convex2 geometry.Convexer, transform2 *geometry.Transform, dp2 *geometry.Vector2, da2, t1, t2 float64, toi *TimeOfImpact) bool {
	iterations := 0
	lerpTx1 := geometry.NewTransform()
	lerpTx2 := geometry.NewTransform()
	separation := narrowphase.NewSeparation()
	separated := c.distanceDetector.Distance(convex1, transform1, convex2, transform2, separation)
	if !separated {
		return false
	}
	d := separation.GetDistance()
	if d < c.distanceEpsilon {
		toi.time = 0
		toi.separation = separation
		return true
	}
	n := separation.GetNormal()
	rmax1 := convex1.GetRadius()
	rmax2 := convex2.GetRadius()
	rv := dp1.DifferenceVector2(dp2)
	rvl := rv.GetMagnitude()
	amax := rmax1*math.Abs(da1) + rmax2*math.Abs(da2)
	if rvl+amax == 0 {
		return false
	}
	l := t1
	l0 := l
	for d > c.distanceEpsilon && iterations < c.maxIterations {
		rvDotN := rv.DotVector2(n)
		drel := rvDotN + amax
		if drel < dyn4go.Epsilon {
			return false
		} else {
			dt := d / drel
			l += dt
			if l < t1 {
				return false
			}
			if l > t2 {
				return false
			}
			if l <= l0 {
				break
			}
			l0 = l
		}
		iterations++
		transform1.LerpDeltaInDestination(dp1, da1, l, lerpTx1)
		transform2.LerpDeltaInDestination(dp2, da2, l, lerpTx2)
		separated := c.distanceDetector.Distance(convex1, lerpTx1, convex2, lerpTx2, separation)
		d := separation.GetDistance()
		if !separated {
			l -= 0.5 * c.distanceEpsilon / drel
			transform1.LerpDeltaInDestination(dp1, da1, l, lerpTx1)
			transform2.LerpDeltaInDestination(dp2, da2, l, lerpTx2)
			separated = c.distanceDetector.Distance(convex1, lerpTx1, convex2, lerpTx2, separation)
			d = separation.GetDistance()
			break
		}
		n = separation.GetNormal()
		d = separation.GetDistance()
	}
	toi.time = l
	toi.separation = separation
	return true
}

func (c *ConservativeAdvancement) GetDistanceDetector() narrowphase.DistanceDetector {
	return c.distanceDetector
}

func (c *ConservativeAdvancement) SetDistanceDetector(distanceDetector narrowphase.DistanceDetector) {
	if distanceDetector == nil {
		panic("Distance detector cannot be nil")
	}
	c.distanceDetector = distanceDetector
}

func (c *ConservativeAdvancement) GetDistanceEpsilon() float64 {
	return c.distanceEpsilon
}

func (c *ConservativeAdvancement) SetDistanceEpsilon(distanceEpsilon float64) {
	if distanceEpsilon <= 0 {
		panic("Distance epsilong must be strictly positive")
	}
	c.distanceEpsilon = distanceEpsilon
}

func (c *ConservativeAdvancement) GetMaxIterations() int {
	return c.maxIterations
}

func (c *ConservativeAdvancement) SetMaxIterations(maxIterations int) {
	if maxIterations < 5 {
		panic("Max iterations cannot be less than 5")
	}
	c.maxIterations = maxIterations
}
