package geometry

import (
	"math"
)

type Ellipse struct {
	AbstractShape
	width, height, a, b float64
	localXAxis          *Vector2
}

func NewEllipse(width, height float64) *Ellipse {
	if width < 0 || height < 0 {
		panic("Ellipse may not have negative width or height")
	}
	e := new(Ellipse)
	e.center = new(Vector2)
	e.width = width
	e.height = height
	e.a = width * 0.5
	e.b = height * 0.5
	e.radius = math.Max(e.a, e.b)
	e.localXAxis = NewVector2FromXY(1, 0)
}

func (e *Ellipse) GetAxes(foci []*Vector2, transform *Transform) []*Vector2 {
	panic("This operation is not supported on Ellipses")
}

func (e *Ellipse) GetFoci(transform *Transform) []*Vector2 {
	panic("This operation is not supported on Ellipses")
}

func (e *Ellipse) GetFarthestPoint(n *Vector2, transform *Transform) *Vector2 {
	localAxis := transform.GetInverseTransformedVector2(n)
	r := e.GetRotation()
	localAxis.RotateAboutOrigin(-r)
	localAxis.X *= e.a
	localAxis.Y *= e.b
	localAxis.Normalize()
	p := NewVector2FromXY(localAxis.X*e.a, localAxis.Y*e.b)
	p.RotateAboutOrigin(r)
	p.AddVector2(e.center)
	transform.TranslateVector2(p)
	return p
}

func (e *Ellipse) GetFarthestFeature(n *Vector2, transform *Transform) *Feature {
	farthest := e.GetFarthestPoint(n, transform)
	return NewVertexFromPoint(farthest)
}

func (e *Ellipse) Project(n *Vector2, transform *Transform) *Interval {
	p1 := e.GetFarthestPoint(n, transform)
	center := transform.GetInverseTransformedVector2(e.center)
	c := center.DotVector2(n)
	d := p1.DotVector2(n)
	return NewIntervalFromMinMax(2*c-d, d)
}

func (e *Ellipse) CreateAABB(transform *Transform) *AABB {
	x := e.Project(NewVector2FromVector2(X_AXIS), transform)
	y := e.Project(NewVector2FromVector2(Y_AXIS), transform)
	return NewAABBFromFloats(x.GetMin(), y.GetMin(), x.GetMax(), y.GetMax())
}

func (e *Ellipse) CreateMass(density float64) *Mass {
	area := math.Pi * e.a * e.b
	m := area * density
	I := m * (e.a*e.a + e.b*e.b) / 4
	return NewMassFromCenterMassInertia(e.center, m, I)
}

func (e *Ellipse) GetRadius(center *Vector2) float64 {
	return e.radius + center.DistanceFromVector2(e.center)
}

func (e *Ellipse) Contains(point *Vector2, transform *Transform) bool {
	localPoint := transform.GetInverseTransformedVector2(point)
	r := e.GetRotation()
	localPoint.RotateAboutXY(-r, e.center.X, e.center.Y)
	x := localPoint.X - e.center.X
	y := localPoint.Y - e.center.Y
	x2 := x * x
	y2 := y * y
	a2 := e.a * e.a
	b2 := e.b * e.b
	value := x2/a2 + y2/b2
	return value <= 1
}

func (e *Ellipse) Rotate(theta, x, y float64) {
	e.RotateAboutXY(theta, x, y)
	e.localXAxis.RotateAboutOrigin(theta)
}

func (e *Ellipse) GetRotation() float64 {
	return X_AXIS.GetAngleBetween(e.localXAxis)
}

func (e *Ellipse) GetWidth() float64 {
	return e.width
}

func (e *Ellipse) GetHeight() float64 {
	return e.height
}

func (e *Ellipse) GetHalfWidth() float64 {
	return e.a
}

func (e *Ellipse) GetHalfHeight() float64 {
	return e.b
}
