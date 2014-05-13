package manifold

import (
	"github.com/LSFN/dyn4go/geometry"
)

type ManifoldPoint struct {
	id    interface{}
	point *geometry.Vector2
	depth float64
}

type ManifoldPointIDDistance struct{}

var DISTANCE ManifoldPointIDDistance

func NewManifoldPoint() *ManifoldPoint {
	m := new(ManifoldPoint)
	m.point = new(geometry.Vector2)
	return m
}

func NewManifoldPointInterfaceVector2Float64(id interface{}, point *geometry.Vector2, depth float64) *ManifoldPoint {
	m := new(ManifoldPoint)
	m.id = id
	m.point = point
	m.depth = depth
	return m
}

func (m *ManifoldPoint) GetID() interface{} {
	return m.id
}

func (m *ManifoldPoint) SetID(id interface{}) {
	m.id = id
}

func (m *ManifoldPoint) GetPoint() *geometry.Vector2 {
	return m.point
}

func (m *ManifoldPoint) SetPoint(point *geometry.Vector2) {
	m.point = point
}

func (m *ManifoldPoint) GetDepth() float64 {
	return m.depth
}

func (m *ManifoldPoint) SetDepth(depth float64) {
	m.depth = depth
}
