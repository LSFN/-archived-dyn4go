package manifold

import (
	"github.com/LSFN/dyn4go/geometry"
)

type Manifold struct {
	points []*ManifoldPoint
	normal *geometry.Vector2
}

func NewManifold() *Manifold {
	m := new(Manifold)
	m.points = make([]*ManifoldPoint, 2)
	return m
}

func NewManifoldManifoldPointsVector2(points []*ManifoldPoint, normal *geometry.Vector2) *Manifold {
	m := new(Manifold)
	m.points = points
	m.normal = normal
	return m
}

func (m *Manifold) Clear() {
	m.points = []*ManifoldPoint{}
	m.normal = nil
}

func (m *Manifold) GetPoints() []*ManifoldPoint {
	return m.points
}

func (m *Manifold) GetNormal() *geometry.Vector2 {
	return m.normal
}

func (m *Manifold) SetPoints(points []*ManifoldPoint) {
	m.points = points
}

func (m *Manifold) SetNormal(normal *geometry.Vector2) {
	m.normal = normal
}
