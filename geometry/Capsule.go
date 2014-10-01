package geometry

import (
	"math"

	"code.google.com/p/uuid"
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
	if width <= 0 || height <= 0 {
		panic("Capsule cannot be created from non-positive width or height.")
	}
	major := width
	minor := height
	vertical := false
	if width < height {
		major, minor = minor, major
		vertical = true
	}

	c := new(Capsule)
	c.length = major
	c.capRadius = minor * 0.5
	c.radius = major * 0.5
	c.center = NewVector2FromXY(0, 0)

	f := (major - minor) * 0.5
	if vertical {
		c.foci = []*Vector2{NewVector2FromXY(0, -f), NewVector2FromXY(0, f)}
		c.localXAxis = NewVector2FromXY(0, 1)
	} else {
		c.foci = []*Vector2{NewVector2FromXY(-f, 0), NewVector2FromXY(f, 0)}
		c.localXAxis = NewVector2FromXY(1, 0)
	}
	c.id = uuid.New()
	return c
}

func (c *Capsule) GetAxes(foci []*Vector2, t *Transform) []*Vector2 {
	if foci != nil {
		axes := make([]*Vector2, 2+len(foci))
		axes[0] = t.GetTransformedR(c.localXAxis)
		axes[1] = t.GetTransformedR(c.localXAxis.GetRightHandOrthogonalVector())
		f1 := t.GetTransformedVector2(c.foci[0])
		f2 := t.GetTransformedVector2(c.foci[1])
		for i, f := range foci {
			d1 := f1.DistanceSquaredFromVector2(f)
			d2 := f2.DistanceSquaredFromVector2(f)
			var v *Vector2
			if d1 < d2 {
				v = f1.HereToVector2(f)
			} else {
				v = f2.HereToVector2(f)
			}
			v.Normalize()
			axes[2+i] = v
		}
		return axes
	} else {
		return []*Vector2{
			t.GetTransformedR(c.localXAxis),
			t.GetTransformedR(c.localXAxis.GetRightHandOrthogonalVector()),
		}
	}
}

func (c *Capsule) GetFoci(t *Transform) []*Vector2 {
	return []*Vector2{
		t.GetTransformedVector2(c.foci[0]),
		t.GetTransformedVector2(c.foci[1]),
	}
}

func (c *Capsule) GetFarthestPoint(v *Vector2, t *Transform) *Vector2 {
	v.Normalize()
	p := GetFarthestPoint(c.foci[0], c.foci[1], v, t)
	return p.AddVector2(v.Product(c.capRadius))
}

func (c *Capsule) GetFarthestFeature(n *Vector2, transform *Transform) Featurer {
	localAxis := transform.GetInverseTransformedR(n)
	n1 := c.localXAxis.GetLeftHandOrthogonalVector()
	d := localAxis.DotVector2(localAxis) * EDGE_FEATURE_SELECTION_CRITERIA
	d1 := localAxis.DotVector2(n1)
	if math.Abs(d1) < d {
		point := c.GetFarthestPoint(n, transform)
		return NewVertexVector2(point)
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

func (c *Capsule) ProjectVector2Transform(n *Vector2, transform *Transform) *Interval {
	p1 := c.GetFarthestPoint(n, transform)
	center := transform.GetTransformedVector2(c.center)
	cDot := center.DotVector2(n)
	d := p1.DotVector2(n)
	return NewIntervalFromMinMax(2*cDot-d, d)
}

func (c *Capsule) CreateAABBTransform(transform *Transform) *AABB {
	x := c.ProjectVector2Transform(NewVector2FromVector2(&X_AXIS), transform)
	y := c.ProjectVector2Transform(NewVector2FromVector2(&Y_AXIS), transform)
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

func (c *Capsule) GetRadiusVector2(center *Vector2) float64 {
	return c.radius + c.center.DistanceFromVector2(center)
}

func (c *Capsule) ContainsVector2Transform(point *Vector2, transform *Transform) bool {
	p := GetPointOnSegmentClosestToPoint(point, transform.GetTransformedVector2(c.foci[0]), transform.GetTransformedVector2(c.foci[1]))
	r2 := c.capRadius * c.capRadius
	d2 := p.DistanceSquaredFromVector2(point)
	return d2 <= r2
}

func (c *Capsule) RotateAboutXY(theta, x, y float64) {
	if !(c.center.X == x && c.center.Y == y) {
		c.center.RotateAboutXY(theta, x, y)
	}
	c.foci[0].RotateAboutXY(theta, x, y)
	c.foci[1].RotateAboutXY(theta, x, y)
	c.localXAxis.RotateAboutOrigin(theta)
}

func (c *Capsule) TranslateXY(x, y float64) {
	c.center.AddXY(x, y)
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

func (c *Capsule) ContainsVector2(v *Vector2) bool {
	return c.ContainsVector2Transform(v, NewTransform())
}

func (c *Capsule) ProjectVector2(v *Vector2) *Interval {
	return c.ProjectVector2Transform(v, NewTransform())
}

func (c *Capsule) CreateAABB() *AABB {
	return c.CreateAABBTransform(NewTransform())
}

func (c *Capsule) RotateAboutOrigin(theta float64) {
	c.RotateAboutXY(theta, 0, 0)
}

func (c *Capsule) RotateAboutCenter(theta float64) {
	c.RotateAboutXY(theta, c.center.X, c.center.Y)
}

func (c *Capsule) RotateAboutVector2(theta float64, v *Vector2) {
	c.RotateAboutXY(theta, v.X, v.Y)
}

func (c *Capsule) TranslateVector2(v *Vector2) {
	c.TranslateXY(v.X, v.Y)
}
