package geometry

import (
	"math"

	"code.google.com/p/uuid"
)

type Ellipse struct {
	AbstractShape
	width, height, a, b float64
	localXAxis          *Vector2
}

func NewEllipse(width, height float64) *Ellipse {
	if width <= 0 || height <= 0 {
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
	e.id = uuid.New()
	return e
}

func (e *Ellipse) GetAxes(foci []*Vector2, transform *Transform) []*Vector2 {
	panic("This operation is not supported on Ellipses")
}

func (e *Ellipse) GetFoci(transform *Transform) []*Vector2 {
	panic("This operation is not supported on Ellipses")
}

func (e *Ellipse) GetFarthestPoint(n *Vector2, transform *Transform) *Vector2 {
	localAxis := transform.GetInverseTransformedR(n)
	r := e.GetRotation()
	localAxis.RotateAboutOrigin(-r)
	localAxis.X *= e.a
	localAxis.Y *= e.b
	localAxis.Normalize()
	p := NewVector2FromXY(localAxis.X*e.a, localAxis.Y*e.b)
	p.RotateAboutOrigin(r)
	p.AddVector2(e.center)
	transform.Transform(p)
	return p
}

func (e *Ellipse) GetFarthestFeature(n *Vector2, transform *Transform) Feature {
	farthest := e.GetFarthestPoint(n, transform)
	return NewVertexFromVector2(farthest)
}

func (e *Ellipse) ProjectTransform(n *Vector2, transform *Transform) *Interval {
	p1 := e.GetFarthestPoint(n, transform)
	center := transform.GetTransformedVector2(e.center)
	c := center.DotVector2(n)
	d := p1.DotVector2(n)
	return NewIntervalFromMinMax(2*c-d, d)
}

func (e *Ellipse) CreateAABBTransform(transform *Transform) *AABB {
	x := e.ProjectTransform(NewVector2FromVector2(&X_AXIS), transform)
	y := e.ProjectTransform(NewVector2FromVector2(&Y_AXIS), transform)
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

func (e *Ellipse) ContainsTransform(point *Vector2, transform *Transform) bool {
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

func (e *Ellipse) RotateAboutXY(theta, x, y float64) {
	if !(e.center.X == x && e.center.Y == y) {
		e.center.RotateAboutXY(theta, x, y)
	}
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

func (e *Ellipse) GetID() string {
	return e.id
}

func (e *Ellipse) GetCenter() *Vector2 {
	return e.center
}

func (e *Ellipse) GetUserData() interface{} {
	return e.userData
}

func (e *Ellipse) SetUserData(data interface{}) {
	e.userData = data
}

func (e *Ellipse) RotateAboutOrigin(theta float64) {
	e.RotateAboutXY(theta, 0, 0)
}

func (e *Ellipse) RotateAboutCenter(theta float64) {
	e.RotateAboutXY(theta, e.center.X, e.center.Y)
}

func (e *Ellipse) RotateAboutVector2(theta float64, v *Vector2) {
	e.RotateAboutXY(theta, v.X, v.Y)
}

func (e *Ellipse) TranslateXY(x, y float64) {
	e.center.AddXY(x, y)
}

func (e *Ellipse) TranslateVector2(v *Vector2) {
	e.TranslateXY(v.X, v.Y)
}

func (e *Ellipse) Project(v *Vector2) *Interval {
	return e.ProjectTransform(v, NewTransform())
}

func (e *Ellipse) Contains(v *Vector2) bool {
	return e.ContainsTransform(v, NewTransform())
}

func (e *Ellipse) CreateAABB() *AABB {
	return e.CreateAABBTransform(NewTransform())
}
