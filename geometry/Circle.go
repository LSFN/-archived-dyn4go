package geometry2

import (
	"math"

	"code.google.com/p/uuid"
)

type Circle struct {
	AbstractShape
}

func NewCircle(radius float64) *Circle {
	if radius <= 0 {
		panic("Circle radius must be positive")
	}
	c := new(Circle)
	c.center = new(Vector2)
	c.radius = radius
	c.id = uuid.New()
	return c
}

func (c *Circle) GetRadiusVector2(v *Vector2) float64 {
	return c.radius + v.DistanceFromVector2(c.center)
}

func (c *Circle) ContainsVector2Transform(v *Vector2, t *Transform) bool {
	v2 := t.GetTransformedVector2(c.center)
	radiusSquared := c.radius * c.radius
	v2.SubtractVector2(v)
	return v2.GetMagnitudeSquared() <= radiusSquared
}

func (c *Circle) ProjectVector2Transform(v *Vector2, t *Transform) *Interval {
	center := t.GetTransformedVector2(c.center)
	c2 := center.DotVector2(v)
	return NewIntervalFromMinMax(c2-c.radius, c2+c.radius)
}

func (c *Circle) GetFarthestFeature(v *Vector2, t *Transform) Featurer {
	return NewVertexVector2(c.GetFarthestPoint(v, t))
}

func (c *Circle) GetFarthestPoint(v *Vector2, t *Transform) *Vector2 {
	nAxis := v.GetNormalized()
	center := t.GetTransformedVector2(c.center)
	center.AddXY(c.radius*nAxis.X, c.radius*nAxis.Y)
	return center
}

func (c *Circle) GetAxes(foci []*Vector2, t *Transform) []*Vector2 {
	return nil
}

func (c *Circle) GetFoci(t *Transform) []*Vector2 {
	foci := make([]*Vector2, 1)
	foci[0] = t.GetTransformedVector2(c.center)
	return foci
}

func (c *Circle) CreateMass(density float64) *Mass {
	mass := density * math.Pi * c.radius * c.radius
	inertia := mass * c.radius * c.radius * 0.5
	return NewMassFromCenterMassInertia(c.center, mass, inertia)
}

func (c *Circle) CreateAABBTransform(t *Transform) *AABB {
	center := t.GetTransformedVector2(c.center)
	return NewAABBFromCenterRadius(center, c.radius)
}

func (c *Circle) ContainsVector2(v *Vector2) bool {
	return c.ContainsVector2Transform(v, NewTransform())
}

func (c *Circle) ProjectVector2(v *Vector2) *Interval {
	return c.ProjectVector2Transform(v, NewTransform())
}

func (c *Circle) CreateAABB() *AABB {
	return c.CreateAABBTransform(NewTransform())
}

func (c *Circle) RotateAboutOrigin(theta float64) {
	c.RotateAboutXY(theta, 0, 0)
}

func (c *Circle) RotateAboutCenter(theta float64) {
	c.RotateAboutXY(theta, c.center.X, c.center.Y)
}

func (c *Circle) RotateAboutVector2(theta float64, v *Vector2) {
	c.RotateAboutXY(theta, v.X, v.Y)
}

func (c *Circle) TranslateVector2(v *Vector2) {
	c.TranslateXY(v.X, v.Y)
}
