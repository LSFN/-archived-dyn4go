package geometry

import (
	"math"
)

type Circle struct {
	AbstractShape
}

func (c *Circle) NewCircle(radius float64) *Circle {
	if radius <= 0 {
		panic("Circle radius must be positive")
	}
	c.center = new(Vector2)
	c.radius = radius
}

func (c *Circle) GetRadius(v *Vector2) float64 {
	return c.radius + v.DistanceFromVector2(c.center)
}

func (c *Circle) Contains(v *Vector2, t *Transform) bool {
	v2 := t.GetTransformedVector2(c.center)
	radiusSquared := c.radius * c.radius
	v2.SubtractVector2(v)
	return v2.GetMagnitudeSquared() <= radiusSquared
}

func (c *Circle) Project(v *Vector2, t *Transform) {
	center := t.GetTransformedVector2(c.center)
	c2 := center.DotVector2(v)
	return NewIntervalFromMinMax(c2-c.radius, c2+c.radius)
}

func (c *Circle) GetFarthestFeature(v *Vector2, t *Transform) *Vertex {
	return NewVertexFromVector2(c.GetFarthestPoint(v, t))
}

func (c *Circle) GetFarthestPoint(v *Vector2, t *Transform) *Vector2 {
	nAxis := v.GetNormalized()
	center = t.GetTransformedVector2(c.center)
	center.add(c.radius*nAxis.X, c.radius*nAxis.Y)
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
	return NewMassFromCenterMassIntertia(c.center, mass, inertia)
}

func (c *Circle) CreateAABB(t *Transform) *AABB {
	center := t.GetTransformedVector2(c.center)
	return NewAABBFromCenterRadius(center, c.radius)
}
