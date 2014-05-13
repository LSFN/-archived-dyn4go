package narrowphase

import (
	"github.com/LSFN/dyn4go/geometry"
)

type Separation struct {
	normal, point1, point2 *geometry.Vector2
	distance               float64
}

func NewSeparation() *Separation {
	s := new(Separation)
	s.normal = new(geometry.Vector2)
	s.point1 = new(geometry.Vector2)
	s.point2 = new(geometry.Vector2)
	return s
}

func NewSeparationVector2Float64Vector2Vector2(normal *geometry.Vector2, distance float64, point1, point2 *geometry.Vector2) *Separation {
	s := new(Separation)
	s.normal = normal
	s.distance = distance
	s.point1 = point1
	s.point2 = point2
	return s
}

func (s *Separation) Clear() {
	s.normal = nil
	s.distance = 0
	s.point1 = nil
	s.point2 = nil
}

func (s *Separation) GetNormal() *geometry.Vector2 {
	return s.normal
}

func (s *Separation) GetDistance() float64 {
	return s.distance
}

func (s *Separation) GetPoint1() *geometry.Vector2 {
	return s.point1
}

func (s *Separation) GetPoint2() *geometry.Vector2 {
	return s.point2
}

func (s *Separation) SetNormal(normal *geometry.Vector2) {
	s.normal = normal
}

func (s *Separation) SetDistance(distance float64) {
	s.distance = distance
}

func (s *Separation) SetPoint1(point1 *geometry.Vector2) {
	s.point1 = point1
}

func (s *Separation) SetPoint2(point2 *geometry.Vector2) {
	s.point2 = point2
}
