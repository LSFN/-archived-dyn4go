package geometry

import (
	"math"
)

type Circle AbstractShape

func NewCircle(radius float64) *Circle {
	if radius <= 0 {
		panic("Circle radius must be positive")
	}
	c := new(Circle)
	c.center = new(Vector2)
	c.radius = radius
	return c
}

func (c *Circle) GetRadius(v *Vector2) float64 {
	return c.radius + v.DistanceFromVector2(c.center)
}

func (c *Circle) ContainsTransform(v *Vector2, t *Transform) bool {
	v2 := t.GetTransformedVector2(c.center)
	radiusSquared := c.radius * c.radius
	v2.SubtractVector2(v)
	return v2.GetMagnitudeSquared() <= radiusSquared
}

func (c *Circle) ProjectTransform(v *Vector2, t *Transform) *Interval {
	center := t.GetTransformedVector2(c.center)
	c2 := center.DotVector2(v)
	return NewIntervalFromMinMax(c2-c.radius, c2+c.radius)
}

func (c *Circle) GetFarthestFeature(v *Vector2, t *Transform) *Vertex {
	return NewVertexFromVector2(c.GetFarthestPoint(v, t))
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

func (c *Circle) GetID() string {
	return c.id
}

func (c *Circle) GetCenter() *Vector2 {
	return c.center
}

func (c *Circle) GetUserData() interface{} {
	return c.userData
}

func (c *Circle) SetUserData(data interface{}) {
	c.userData = data
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

func (c *Circle) RotateAboutXY(theta, x, y float64) {
	if !(c.center.X == x && c.center.Y == y) {
		c.center.RotateAboutXY(theta, x, y)
	}
}

func (c *Circle) TranslateXY(x, y float64) {
	c.center.AddXY(x, y)
}

func (c *Circle) TranslateVector2(v *Vector2) {
	c.TranslateXY(v.X, v.Y)
}

func (c *Circle) Project(v *Vector2) *Interval {
	return c.ProjectTransform(v, NewTransform())
}

func (c *Circle) Contains(v *Vector2) bool {
	return c.ContainsTransform(v, NewTransform())
}

func (c *Circle) CreateAABB() *AABB {
	return c.CreateAABBTransform(NewTransform())
}
