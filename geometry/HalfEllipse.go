package geometry2

import (
	"math"

	"code.google.com/p/uuid"
)

const (
	INERTIA_CONSTANT = math.Pi/8.0 - 8.0/(9.0*math.Pi)
)

type HalfEllipse struct {
	AbstractShape
	width, height, a          float64
	localXAxis, ellipseCenter *Vector2
	vertices                  []*Vector2
}

func NewHalfEllipse(width, height float64) *HalfEllipse {
	if width <= 0 || height <= 0 {
		panic("Width and height of new half ellipse must be strictly positive")
	}
	h := new(HalfEllipse)
	h.width = width
	h.height = height
	h.a = width * 0.5
	h.ellipseCenter = new(Vector2)
	h.center = NewVector2FromXY(0, (4*height)/(3*math.Pi))
	h.localXAxis = NewVector2FromXY(1, 0)
	h.vertices = []*Vector2{
		NewVector2FromXY(-h.a, 0),
		NewVector2FromXY(h.a, 0),
	}
	h.radius = h.center.DistanceFromVector2(h.vertices[1])
	h.id = uuid.New()
	return h
}

func (h *HalfEllipse) GetAxes(foci []*Vector2, transform *Transform) []*Vector2 {
	panic("This operation is not supported on half ellipses")
}

func (h *HalfEllipse) GetFoci(transform *Transform) []*Vector2 {
	panic("This operation is not supported on half ellipses")
}

func (h *HalfEllipse) GetFarthestPoint(n *Vector2, transform *Transform) *Vector2 {
	localAxis := transform.GetInverseTransformedR(n)
	r := h.GetRotation()
	localAxis.RotateAboutOrigin(-r)
	localAxis.X *= h.a
	localAxis.Y *= h.height
	localAxis.Normalize()
	if localAxis.Y <= 0 && localAxis.X >= 0 {
		return transform.GetTransformedVector2(h.vertices[1])
	} else if localAxis.Y <= 0 && localAxis.X <= 0 {
		return transform.GetTransformedVector2(h.vertices[0])
	}
	p := NewVector2FromXY(localAxis.X*h.a, localAxis.Y*h.height)
	p.RotateAboutOrigin(r)
	p.AddVector2(h.ellipseCenter)
	transform.Transform(p)
	return p
}

func (h *HalfEllipse) GetFarthestFeature(n *Vector2, transform *Transform) Featurer {
	localAxis := transform.GetInverseTransformedR(n)
	if localAxis.GetAngleBetween(h.localXAxis) < 0 {
		return NewVertexVector2(h.GetFarthestPoint(n, transform))
	} else {
		return GetFarthestFeature(h.vertices[0], h.vertices[1], n, transform)
	}
}

func (h *HalfEllipse) ProjectVector2Transform(n *Vector2, transform *Transform) *Interval {
	p1 := h.GetFarthestPoint(n, transform)
	p2 := h.GetFarthestPoint(n.GetNegative(), transform)
	d1 := p1.DotVector2(n)
	d2 := p2.DotVector2(n)
	return NewIntervalFromMinMax(d2, d1)
}

func (h *HalfEllipse) CreateAABBTransform(transform *Transform) *AABB {
	x := h.ProjectVector2Transform(NewVector2FromVector2(&X_AXIS), transform)
	y := h.ProjectVector2Transform(NewVector2FromVector2(&Y_AXIS), transform)
	return NewAABBFromFloats(x.GetMin(), y.GetMin(), x.GetMax(), y.GetMax())
}

func (h *HalfEllipse) CreateMass(density float64) *Mass {
	area := math.Pi * h.a * h.height
	m := area * density * 0.5
	I := m * (h.a*h.a + h.height*h.height) * INERTIA_CONSTANT
	return NewMassFromCenterMassInertia(h.center, m, I)
}

func (h *HalfEllipse) GetRadiusVector2(center *Vector2) float64 {
	return h.radius + center.DistanceFromVector2(h.center)
}

func (h *HalfEllipse) ContainsVector2Transform(point *Vector2, transform *Transform) bool {
	localPoint := transform.GetInverseTransformedVector2(point)
	r := h.GetRotation()
	localPoint.RotateAboutXY(-r, h.ellipseCenter.X, h.ellipseCenter.Y)
	x := localPoint.X - h.ellipseCenter.X
	y := localPoint.Y - h.ellipseCenter.Y
	if y < 0 {
		return false
	}
	x2 := x * x
	y2 := y * y
	a2 := h.a * h.a
	b2 := h.height * h.height
	value := x2/a2 + y2/b2
	return value <= 1.0
}

func (h *HalfEllipse) RotateAboutXY(theta, x, y float64) {
	if !(h.center.X == x && h.center.Y == y) {
		h.center.RotateAboutXY(theta, x, y)
	}
	h.localXAxis.RotateAboutOrigin(theta)
	for _, v := range h.vertices {
		v.RotateAboutXY(theta, x, y)
	}
	h.ellipseCenter.RotateAboutXY(theta, x, y)
}

func (h *HalfEllipse) TranslateXY(x, y float64) {
	h.center.AddXY(x, y)
	for _, v := range h.vertices {
		v.AddXY(x, y)
	}
	h.ellipseCenter.AddXY(x, y)
}

func (h *HalfEllipse) GetRotation() float64 {
	return X_AXIS.GetAngleBetween(h.localXAxis)
}

func (h *HalfEllipse) GetWidth() float64 {
	return h.width
}

func (h *HalfEllipse) GetHeight() float64 {
	return h.height
}

func (h *HalfEllipse) GetHalfWidth() float64 {
	return h.a
}

func (h *HalfEllipse) GetEllipseCenter() *Vector2 {
	return h.ellipseCenter
}

func (h *HalfEllipse) ContainsVector2(v *Vector2) bool {
	return h.ContainsVector2Transform(v, NewTransform())
}

func (h *HalfEllipse) ProjectVector2(v *Vector2) *Interval {
	return h.ProjectVector2Transform(v, NewTransform())
}

func (h *HalfEllipse) CreateAABB() *AABB {
	return h.CreateAABBTransform(NewTransform())
}

func (h *HalfEllipse) RotateAboutOrigin(theta float64) {
	h.RotateAboutXY(theta, 0, 0)
}

func (h *HalfEllipse) RotateAboutCenter(theta float64) {
	h.RotateAboutXY(theta, h.center.X, h.center.Y)
}

func (h *HalfEllipse) RotateAboutVector2(theta float64, v *Vector2) {
	h.RotateAboutXY(theta, v.X, v.Y)
}

func (h *HalfEllipse) TranslateVector2(v *Vector2) {
	h.TranslateXY(v.X, v.Y)
}
