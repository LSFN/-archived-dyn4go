package geometry

import (
	"math"
)

const (
	EDGE_FEATURE_SELECTION_CRITERIA = 0.98
	EDGE_FEATURE_EXPANSION_FACTOR   = 0.1
)

type Capsule struct {
	AbstractShape
	length, capRadius float64
	foci              []*Vector2
	localXAxis        *Vector2
}

func NewCapsule(width, height float64) *Capsule {
	major := width
	minor := height
	vertical := false
	if width < heigth {
		major, minor = minor, major
		vertical = true
	}

	c := new(Capsule)
	c.length = major
	c.capRadius = minor * 0.5
	c.radius = major * 0.5
	c.center = NewVector2FromXY(0, 0)

	f = (major - minor) * 0.5
	if vertical {
		c.foci = [2]*Vector2{NewVector2FromXY(0, -f), NewVector2FromXY(0, f)}
		c.loaclXAxis = NewVector2FromXY(0, 1)
	} else {
		c.foci = [2]*Vector2{NewVector2FromXY(-f, 0), NewVector2FromXY(f, 0)}
		c.loaclXAxis = NewVector2FromXY(1, 0)
	}

	return c
}

func (c *Capsule) GetAxes(foci []*Vector2, t *Transform) {
	if c.foci != nil {
		axes := make([]*Vector2, 2+len(c.foci))
		axes[0] = t.GetTransformedR(c.locallXAxis)
		axes[1] = t.GetTransformedR(c.locallXAxis.GetRightHandOrthogonalVector())
		f1 := t.GetTransformed(c.foci[0])
		f2 := t.GetTransformed(c.foci[1])
		for i, f := range c.foci {
			if f1.DistanceSquared(f) < f2.DistanceSquared(f) {
				axes[2+i] = f1.HereToVector2(f)
			} else {
				axes[2+i] = f2.HereToVector2(f)
			}
		}
		return axes
	} else {
		return []*Vector2{t.GetTransformedR(c.localXAxis), t.GetTransformedR(c.localXAxis).GetRightHandOrthogonalVector()}
	}
}

func (c *Capsule) GetFoci(t *Transform) {
	return []*Vector2{t.GetTransformed(c.foci[0]), t.GetTransformed(c.foci[1])}
}

func (c *Capsule) GetFarthestPoint(v *Vector2, t *Transform) {
	v.Normalize()
	p := GetFarthestPoint(c.foci[0], c.foci[1], v, t)
	return p.AddVector2(v.Product(c.capRadius))
}

func (c *Capsule) GetFarthestFeature(n *Vector2, transform *Transform) *Feature {
	localAxis := transform.GetInverseTransformedVector2(n)
	n1 := c.localXAxis.GetLeftHandOrthogonalVector()
	d := localAxis.DotVector2(localAxis)
	d1 := localAxis.DotVector2(n1)
	if math.Abs(d1) < d {
		point := c.GetFarthestPoint(n, transform)
		return NewVertexFromPoint(point)
	} else {
		v := n1.Multiply(c.capRadius)
		e := c.localXAxis.Product(c.length * 0.5 * EDGE_FEATURE_EXPANSION_FACTOR)
		if d1 > 0 {
			p1 := c.foci[0].SumVector2(v).SubtractVector2(e)
			p2 := c.foci[1].SumVector2(v).AddVector2(e)
			return GetFarthestFeature(p1, p2, n, transform)
		} else {
			p1 := c.foci[0].DifferenceVector2(v).SubtractVector2(e)
			p2 := c.foci[1].DifferenceVector2(v).AddVector2(e)
			return GetFarthestFeature(p1, p2, n, transform)
		}
	}
}

func (c *Capsule) Project(n *Vector2, transform *Transform) *Interval {
	p1 := c.GetFarthestPoint(n, transform)
	center := transform.GetTransformedVector2(c.center)
	cDot := center.DotVector2(n)
	d := p1.DotVector2(n)
	return NewIntervalFromMinMax(2*cDot-d, d)
}

func (c *Capsule) CreateAABB(transform *Transform) *AABB {
	x := c.Project(NewVector2FromVector2(X_AXIS), transform)
	y := c.Project(NewVector2FromVector2(Y_AXIS), transform)
	return NewAABBFromFloats(x.GetMin(), y.GetMin(), x.GetMax(), y.GetMax())
}

func (c *Capsule) CreateMass(density float64) *Mass {
	h := c.capRadius * 2
	w := c.length - h
	r2 := c.capRadius * c.capRadius
	ra := w * h
	ca := r2 * math.Pi
	rm := density * ra
	cm := density * ca
	m := rm + cm
	d := w * 0.5
	cI := 0.5*cm*r2 + cm*d*d
	rI := rm * (h*h + w*w) / 12
	I := rI + cI
	return NewMassFromCenterMassInertia(c.center, m, I)
}

func (c *Capsule) GetRadius(center *Vector2) float64 {
	return c.radius + c.center.DistanceFromVector2(center)
}

func (c *Capsule) Contains(point *Vector2, transform *Transform) bool {
	p := GetPointOnSegmentClosestToPoint(point, transform.getTransformed(c.foci[0]), transform.getTransformed(c.foci[1]))
	r2 := c.capRadius * c.capRadius
	d2 := p.DistanceSquaredFromVector2(point)
	return d2 <= r2
}

func (c *Capsule) Rotate(theta, x, y float64) {
	c.RotateAboutXY(theta, x, y)
	c.foci[0].RotateAboutXY(theta, x, y)
	c.foci[1].RotateAboutXY(theta, x, y)
	c.localXAxis.RotateAboutOrigin(theta)
}

func (c *Capsule) Translate(x, y float64) {
	c.TranslateXY(x, y)
	c.foci[0].AddXY(x, y)
	c.foci[1].AddXY(x, y)
}

func (c *Capsule) GetRotation() float64 {
	return X_AXIS.GetAngleBetween(c.localXAxis)
}

func (c *Capsule) GetLength() float64 {
	return c.length
}

func (c *Capsule) GetCapRadius() float64 {
	return c.capRadius
}
