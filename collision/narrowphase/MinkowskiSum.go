package narrowphase

import (
	"github.com/LSFN/dyn4go/geometry"
)

type MinkowskiSumPoint struct {
	p1, p2, p *geometry.Vector2
}

func NewMinkowskiSumPoint() *MinkowskiSumPoint {
	m := new(MinkowskiSumPoint)
	m.p = new(geometry.Vector2)
	m.p1 = new(geometry.Vector2)
	m.p2 = new(geometry.Vector2)
	return m
}

func NewMinkowskiSumPointVector2Vector2(p1, p2 *geometry.Vector2) *MinkowskiSumPoint {
	m := new(MinkowskiSumPoint)
	m.p = new(geometry.Vector2)
	m.Set(p1, p2)
	return m
}

func (m *MinkowskiSumPoint) Set(p1, p2 *geometry.Vector2) {
	m.p1 = p1
	m.p2 = p2
	m.p = p1.DifferenceVector2(p2)
}

func (m *MinkowskiSumPoint) Set(p *MinkowskiSumPoint) {
	m.p1 = p.p1
	m.p2 = p.p2
	m.p = p.p
}

type MinkowskiSum struct {
	convex1, convex2       geometry.Convexer
	transform1, transform2 *geometry.Transform
}

func NewMinkowskiSum(convex1 geometry.Convexer, transform1 *geometry.Transform, convex2 geometry.Convexer, transform2 *geometry.Transform) *MinkowskiSum {
	m := new(MinkowskiSum)
	m.convex1 = convex1
	m.convex2 = convex2
	m.transform1 = transform1
	m.transform2 = transform2
}

func (m *MinkowskiSum) Support(direction *geometry.Vector2) *geometry.Vector2 {
	point1 := m.convex1.GetFarthestPoint(direction, m.transform1)
	direction.Negate()
	point2 := m.convex2.GetFarthestPoint(direction, m.transform2)
	direction.Negate()
	return point1.SubtractVector2(point2)
}

func (m *MinkowskiSum) SupportMinkowskiSumPoint(direction *geometry.Vector2, p *MinkowskiSumPoint) {
	point1 := m.convex1.GetFarthestPoint(direction, m.transform1)
	direction.Negate()
	point2 := m.convex2.GetFarthestPoint(direction, m.transform2)
	direction.Negate()
	p.Set(point1, point2)
}

func (m *MinkowskiSum) GetConvex1() geometry.Convexer {
	return m.convex1
}

func (m *MinkowskiSum) SetConvex1(convex1 geometry.Convexer) {
	m.convex1 = convex1
}

func (m *MinkowskiSum) GetConvex2() geometry.Convexer {
	return m.convex2
}

func (m *MinkowskiSum) SetConvex2(convex2 geometry.Convexer) {
	m.convex2 = convex2
}

func (m *MinkowskiSum) GetTransform1() *geometry.Transform {
	return m.transform1
}

func (m *MinkowskiSum) SetTransform1(transform1 *geometry.Transform) {
	m.transform1 = transform1
}

func (m *MinkowskiSum) GetTransform2() *geometry.Transform {
	return m.transform2
}

func (m *MinkowskiSum) SetTransform2(transform2 *geometry.Transform) {
	m.transform2 = transform2
}
