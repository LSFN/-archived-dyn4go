package geometry

import (
	"math"
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

func NewHalfEllipse(width, height *float64) *HalfEllipse {
	if width < 0 || height < 0 {
		panci("Cannot define half ellipse with negative width or height")
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
	return e
}

func (h *HalfEllipse) GetAxes(foci []*Vector2, transform *Transform) []*Vector2 {
	panic("This operation is not supported on half ellipses")
}

func (h *HalfEllipse) GetFoci(transform *Transform) []*Vector2 {
	panic("This operation is not supported on half ellipses")
}

func (h *HalfEllipse) GetFarthestPoint(n *Vector2, transform *Transform) *Vector2 {
	localAxis := transform.GetInverseTransformedVector2(n)
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
	p.AddVector2(h.center)
	transform.TranslateVector2(p)
	return p
}

func (h *HalfEllipse) GetFarthestFeature(n *Vector2, transform *Transform) *Feature {
	localAxis := transform.GetInverseTransformedR(n)
	if localAxis.GetAngleBetween(h.localXAxis) < 0 {
		return NewVertexFromPoint(h.GetFarthestPoint(n, transform))
	} else {
		return GetFarthestFeature(h.vertices[0], h.vertices[1], n, t)
	}
}

func (h *Ellipse) Project(n *Vector2, transform *Transform) *Interval {
	p1 := h.GetFarthestPoint(n, transform)
	p2 := h.GetFarthestPoint(n.GetNegative(), transform)
	d1 := p1.DotVector2(n)
	d2 := p2.DotVector2(n)
	return NewIntervalFromMinMax(d2, d1)
}

func (h *Ellipse) CreateAABB(transform *Transform) *AABB {
	x := h.Project(NewVector2FromVector2(X_AXIS), transform)
	y := h.Project(NewVector2FromVector2(Y_AXIS), transform)
	return NewAABBFromFloats(x.GetMin(), y.GetMin(), x.GetMax(), y.GetMax())
}

func (h *HalfEllipse) CreateMass(density float64) *Mass {
	area := math.PI * h.a * h.height
	m := area * density * 0.5
	I := m * (h.a*h.a + h.height*h.height) * INERTIA_CONSTANT
	return NewMassFromCenterMassInertia(h.center, m, I)
}

func (h *Ellipse) GetRadius(center *Vector2) float64 {
	return h.radius + center.DistanceFromVector2(h.center)
}

func (h *Ellipse) Contains(point *Vector2, transform *Transform) bool {
	localPoint := transform.GetInverseTransformedVector2(point)
	r := h.GetRotation()
	localPoint.RotateAboutXY(-r, h.ellipseCenter.X, h.ellipseCenter.Y)
	x := localPoint.X - h.ellipseCenter.X
	y := localPoint.Y - h.ellipseCenter.Y
	x2 := x * x
	y2 := y * y
	a2 := h.a * h.a
	b2 := h.b * h.b
	value := x2/a2 + y2/b2
	return value <= 1
}

func (h *HalfEllipse) Rotate(theta, x, y float64) {
	h.Rotate(theta, x, y)
	h.localXAxis.RotateAboutOrigin(theta)
	for _, v := range h.vertices {
		v.RotateAboutXY(theta, x, y)
	}
	h.ellipseCenter.RotateAboutXY(theta, x, y)
}

func (h *HalfEllipse) Translate(x, y float64) {
	h.TranslateXY(x, y)
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

func (h *HalfEllipse) GetEllipseCenter() float64 {
	return h.ellipseCenter
}
